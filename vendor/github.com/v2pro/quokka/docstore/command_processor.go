package docstore

import (
	"time"
	"github.com/v2pro/quokka/docstore/runtime"
	"errors"
	"github.com/v2pro/quokka/kvstore"
	"github.com/v2pro/plz/countlog"
	"github.com/json-iterator/go"
	"encoding/json"
	"context"
	"fmt"
)

const ErrUnknown = 10000
const ErrEventLogConflict = 10001 // the master is no longer in charge, should find out who is the master
const ErrForwardedTooManyTimes = 10002 // the master can not be found
const ErrRequestSchemaViolated = 10003
const ErrDocumentSchemaViolated = 10004
const ErrResponseSchemaViolated = 10005
const ErrEntityMissing = 10006 // query can not find the entity

// error number within range [LoggedErrStart, LoggedErrEnd) is committed in event_log as fact
const LoggedErrStart = 20000
const ErrDuplicatedEntity = LoggedErrStart // create entity with same entity id but with different command id
const ErrMustCreateFirst = 20001           // entity not created yet, so can not exec other command
const ErrBusinessRuleViolated = 20002      // `throw` in javascript code means business rule violation
const LoggedErrEnd = 30000

type commandProcessor struct {
	partitionId   uint64
	entityType    string
	reqChan       chan *command
	lastEventId   uint64
	entityLookup  *entityLookup
	commandLookup *commandLookup
	lookupSyncer  *eventProcessor
}

type command struct {
	EntityType     string
	CommandType    string
	EntityId       string
	CommandId      string
	CommandRequest json.RawMessage
	IsPromoting    bool // update topo when command executed successfully
	ForwardedTimes int
	respChan       chan []byte
}

func init() {
	for partitionId := uint64(0); partitionId < kvstore.PartitionsCount; partitionId++ {
		commandProcessors[partitionId] = map[string]*commandProcessor{}
	}
}

func newCommandProcessor(partitionId uint64, entityType string) *commandProcessor {
	processor := &commandProcessor{
		partitionId:   partitionId,
		entityType:    entityType,
		reqChan:       make(chan *command),
		entityLookup:  newEntityLookup(),
		commandLookup: newCommandLookup(),
	}
	processor.lookupSyncer = newEventProcessor(partitionId, processor)
	return processor
}

func (processor *commandProcessor) init(ctx context.Context) error {
	offset, err := processor.entityLookup.kvstoreLookup.getOffset(ctx, processor.partitionId, processor.entityType)
	if err != nil {
		return err
	}
	iter, err := kvstore.Scan(ctx, processor.partitionId, processor.entityType, offset+1)
	if err != nil {
		return err
	}
	for {
		rows, err := iter()
		if err != nil {
			return err
		}
		if rows == nil {
			processor.lastEventId = offset
			return nil
		}
		for _, row := range rows {
			offset = row.Key
			var event Event
			err = runtime.Json.Unmarshal(row.Value, &event)
			if err != nil {
				countlog.Error("event!command_processor.failed to unmarshal event",
					"err", err,
					"encodedEvent", row.Value)
				return err
			}
			event.PartitionId = processor.partitionId
			event.EntityType = processor.entityType
			event.EventId = row.Key
			processor.commandLookup.cacheCommand(event.CommandId, event.CommandResponse, event.EventId)
			err := processor.replayEvent(ctx, &event)
			if err != nil {
				return err
			}
		}
	}
}

func (processor *commandProcessor) replayEvent(ctx context.Context, event *Event) error {
	var err error
	entity, _ := processor.entityLookup.memLookup.getCacheValue(event.EntityId).(*entity)
	if entity == nil {
		entity = newEntity()
		err = loadEntity(ctx, processor.partitionId, processor.entityType, event.EntityId, event.EventId, entity)
		if err != nil {
			return err
		}
		processor.entityLookup.cacheEntity(event.EntityId, entity, event.EventId)
		return nil
	}
	if entity.doc != nil {
		err = runtime.DeltaJson.Unmarshal(event.Delta, entity.doc)
		if err != nil {
			countlog.Error("event!command_processor.failed to apply delta",
				"err", err,
				"state", event.State)
			return err
		}
	}
	entity.eventId = event.EventId
	entity.version++
	processor.entityLookup.cacheEntity(event.EntityId, entity, event.EventId)
	return nil
}

func (processor *commandProcessor) executeInBackground(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}
		processor.executeOne(ctx)
	}
}

func (processor *commandProcessor) executeOne(ctx context.Context) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!command_processor.executeOne.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	select {
	case command := <-processor.reqChan:
		resp := processor.exec(ctx, command)
		command.respChan <- resp
	case <-ctx.Done():
		return
	}
}

func (processor *commandProcessor) delegatedExec(cmd *command, timeout time.Duration) []byte {
	cmd.respChan = make(chan []byte, 1)
	timer := time.NewTimer(timeout)
	select {
	case processor.reqChan <- cmd:
		resp := <-cmd.respChan
		return resp
	case <-timer.C:
		return replyError(errors.New("timeout"))
	}
}

func (processor *commandProcessor) exec(ctx context.Context, cmd *command) []byte {
	partition := processor.partitionId
	entityId := cmd.EntityId
	req := cmd.CommandRequest
	entityType := cmd.EntityType
	commandType := cmd.CommandType
	commandId := cmd.CommandId
	previousResp, err := processor.commandLookup.getCommand(ctx, partition, entityType, cmd.CommandId)
	if err != nil {
		return replyError(err)
	}
	if previousResp != nil {
		return previousResp
	}
	var reqObj interface{}
	if len(req) > 0 {
		err = runtime.Json.Unmarshal(req, &reqObj)
		if err != nil {
			return replyError(err)
		}
	} else {
		req = json.RawMessage("{}")
	}

	entityTypeDef := getEntityType(entityType)
	if entityTypeDef == nil {
		return replyError(errors.New("entityTypeDef not defined for entity type " + entityType))
	}
	commandDef := entityTypeDef.getCommandDef(commandType)
	if commandDef == nil {
		return replyError(errors.New("handler not defined for command type " + commandType))
	}
	err = runtime.Validate(reqObj, commandDef.requestSchema)
	if err != nil {
		err = fmt.Errorf("request violated the schema: %s", err.Error())
		return replyError(withErrorNumber(err, ErrRequestSchemaViolated))
	}
	entity, respObj := processor.handle(ctx, cmd, commandDef.handler, reqObj)
	var resp []byte
	if err, _ := respObj.(error); err != nil {
		resp = replyError(err)
	} else {
		err = runtime.Validate(respObj, commandDef.responseSchema)
		if err != nil {
			err = fmt.Errorf("response violated the schema: %s", err.Error())
			return replyError(withErrorNumber(err, ErrResponseSchemaViolated))
		}
		resp, err = runtime.Json.Marshal(respObj)
		if err != nil {
			return replyError(err)
		}
		resp = replySuccess(resp)
	}
	event := &Event{
		PartitionId:     partition,
		EntityType:      entityType,
		EventId:         processor.lastEventId + 1,
		BaseEventId:     entity.eventId,
		EntityId:        entityId,
		Version:         entity.version + 1,
		CommandId:       commandId,
		CommandType:     commandType,
		CommandRequest:  req,
		HandlerVersion:  commandDef.handlerVersion,
		CommandResponse: resp,
		CommittedAt:     time.Now().UnixNano(),
	}
	if entity.version%16 == 0 {
		event.State, err = runtime.Json.Marshal(entity.doc)
	} else {
		event.State = []byte("null")
	}
	event.Delta, err = runtime.DeltaJson.Marshal(entity.doc)
	if err != nil {
		return replyError(err)
	}
	encodedEvent, err := eventJson.Marshal(event)
	if err != nil {
		return replyError(err)
	}
	err = kvstore.Append(ctx, processor.partitionId, event.EntityType, event.EventId, encodedEvent)
	if err != nil {
		if err == kvstore.KeyConflictError {
			return replyError(withErrorNumber(err, ErrEventLogConflict))
		}
		return replyError(err)
	}
	countlog.Trace("event!docstore.stored event",
		"partitionId", processor.partitionId,
		"eventId", event.EventId,
		"entityId", event.EntityId,
		"encodedEvent", encodedEvent)
	processor.lastEventId = event.EventId
	// update in memory lookup
	entity.eventId = event.EventId
	entity.version = event.Version
	processor.entityLookup.cacheEntity(event.EntityId, entity, event.EventId)
	processor.commandLookup.cacheCommand(event.CommandId, resp, event.EventId)
	// up kv entityTypeDef lookup in separate goroutine
	processor.lookupSyncer.enqueue(event)
	if cmd.IsPromoting {
		thisNodeExecutor.Go(func(ctx context.Context) {
			topo.savePromotedMasterInBackground(ctx, partition)
		})
	}
	return resp
}

func (processor *commandProcessor) handle(ctx context.Context, cmd *command, handler CommandHandler,
	reqObj interface{}) (*entity, interface{}) {

	partition := processor.partitionId
	entityId := cmd.EntityId
	entityType := cmd.EntityType
	commandType := cmd.CommandType
	entity, err := processor.entityLookup.getEntity(ctx, partition, entityType, entityId)
	if entity == nil {
		entity = newEntity()
	}
	if commandType == "create" {
		if err != nil && err != entityNotFoundError {
			return entity, err
		}
		if entity.doc != nil {
			return entity, withErrorNumber(errors.New("duplicated entity"), ErrDuplicatedEntity)
		}
		entity.doc = runtime.NewObject()
	} else if err != nil {
		return entity, err
	}
	return entity, handler(entity.doc, reqObj)
}

func (processor *commandProcessor) LoadOffset(ctx context.Context, partitionId uint64, entityType string) (uint64, error) {
	offset, err := kvstore.GetMonotonic(ctx, partitionId, entityType, "offset")
	if err != nil {
		return 0, err
	}
	return offset, nil
}

func (processor *commandProcessor) CommitOffset(ctx context.Context, partitionId uint64, entityType string, offset uint64) error {
	return kvstore.SetMonotonic(ctx, partitionId, entityType, "offset", offset)
}

func (processor *commandProcessor) Sync(ctx context.Context, event *Event) error {
	thisNodeExecutor.Go(func(ctx context.Context) {
		forwardEventInBackground(ctx, event)
	})
	err := processor.entityLookup.setEventId(ctx, event.PartitionId, event.EntityType, event.EntityId, event.EventId)
	if err != nil {
		countlog.Error("event!core_event_handler.failed to update entity lookup",
			"err", err,
			"partitionId", event.PartitionId,
			"entityId", event.EntityId,
			"eventId", event.EventId)
		return err
	}
	err = processor.entityLookup.setEventId(ctx, event.PartitionId, event.EntityType, event.CommandId, event.EventId)
	if err != nil {
		countlog.Error("event!core_event_handler.failed to update command lookup",
			"err", err,
			"partitionId", event.PartitionId,
			"entityId", event.EntityId,
			"eventId", event.EventId)
		return err
	}
	return nil
}

func replySuccess(resp []byte) []byte {
	output := append([]byte(`{"errno":0, "handled_by":"`), thisNodeAddr...)
	output = append(output, `", "data":`...)
	output = append(output, resp...)
	output = append(output, '}')
	return output
}

type NumberedError interface {
	error
	ErrorNumber() int
}

type numberedError struct {
	errorNumber int
	err         error
}

func (err *numberedError) Error() string {
	return err.err.Error()
}

func (err *numberedError) ErrorNumber() int {
	return err.errorNumber
}

func withErrorNumber(err error, errorNumber int) *numberedError {
	return &numberedError{err: err, errorNumber: errorNumber}
}

func replyError(err error) []byte {
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteObjectStart()
	stream.WriteObjectField("errno")
	numberedErr, _ := err.(NumberedError)
	if numberedErr == nil {
		stream.WriteVal(ErrUnknown)
	} else {
		errorNumber := numberedErr.ErrorNumber()
		stream.WriteVal(errorNumber)
		isCommitted := errorNumber >= LoggedErrStart && errorNumber < LoggedErrEnd
		stream.WriteMore()
		stream.WriteObjectField("committed")
		stream.WriteVal(isCommitted)
	}
	stream.WriteMore()
	stream.WriteObjectField("errmsg")
	stream.WriteString(err.Error())
	stream.WriteMore()
	stream.WriteObjectField("handled_by")
	stream.WriteVal(thisNodeAddr)
	stream.WriteObjectEnd()
	return stream.Buffer()
}
