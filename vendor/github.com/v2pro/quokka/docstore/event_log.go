package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/plz/countlog"
	"errors"
	"context"
	"encoding/json"
	"bytes"
)

var eventJson = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

type entity struct {
	// is baseEventId before command, will be eventId after command
	eventId uint64
	version uint64
	doc     interface{}
}

func newEntity() *entity {
	return &entity{}
}

func castToEntity(obj interface{}) *entity {
	e, _ := obj.(*entity)
	return e
}

type eventForLoadEntity struct {
	BaseEventId uint64          `json:"b"`
	Version     uint64          `json:"v"`
	State       json.RawMessage `json:"s"`
	Delta       json.RawMessage `json:"d"`
}

type eventForGetCommandResponse struct {
	CommandResponse json.RawMessage `json:"p"`
}

func loadEntity(ctx context.Context, partitionId uint64, entityType string, entityId string,
	eventId uint64, /* in and out */ entity *entity) error {

	encodedEvent, err := kvstore.Get(ctx, partitionId, entityType, eventId)
	if err != nil {
		countlog.Trace("event!event_log.event not found",
			"err", err,
			"entityCommandHandlers", entityType,
			"entityId", entityId,
			"partitionId", partitionId,
			"eventId", eventId)
		return err
	}
	var event eventForLoadEntity
	if err = eventJson.Unmarshal(encodedEvent, &event); err != nil {
		countlog.Trace("event!event_log.failed to unmarshal event",
			"err", err,
			"partitionId", partitionId,
			"eventId", eventId,
			"encodedEvent", encodedEvent)
		return err
	}
	version := event.Version
	baseEventId := eventId
	var deltas [][]byte
	for bytes.Equal(event.State, []byte("null")) {
		deltas = append(deltas, event.Delta)
		if event.BaseEventId == 0 {
			countlog.Error("event!event_log.can not find event with state",
				"partitionId", partitionId,
				"baseEventId", baseEventId,
				"deltas", deltas)
			entity.eventId = baseEventId
			entity.version = version
			entity.doc = nil
			return nil
		}
		if event.BaseEventId == entity.eventId {
			break
		}
		encodedEvent, err = kvstore.Get(ctx, partitionId, entityType, event.BaseEventId)
		if err != nil {
			return err
		}
		event.BaseEventId = 0
		event.State = nil
		event.Delta = nil
		if err = eventJson.Unmarshal(encodedEvent, &event); err != nil {
			return err
		}
	}
	doc := entity.doc
	if err = runtime.Json.Unmarshal(event.State, &doc); err != nil {
		return err
	}
	// apply delta backwards
	for i := len(deltas) - 1; i >= 0; i-- {
		delta := deltas[i]
		if err = runtime.DeltaJson.Unmarshal(delta, &doc); err != nil {
			return err
		}
	}
	entity.eventId = baseEventId
	entity.version = version
	entity.doc = doc
	return nil
}

// scanEvents return events in range [from, to)
func scanEvents(ctx context.Context, partitionId uint64, entityType string, from uint64, to uint64) ([]*Event, error) {
	if to <= from {
		return []*Event{}, nil
	}
	iter, err := kvstore.Scan(ctx, partitionId, entityType, from)
	if err != nil {
		return nil, err
	}
	events := make([]*Event, 0, 10)
	for {
		rows, err := iter()
		if err != nil {
			return nil, err
		}
		if rows == nil {
			break
		}
		for _, row := range rows {
			eventId := row.Key
			if eventId >= to {
				return events, nil
			}
			var event Event
			err := eventJson.Unmarshal(row.Value, &event)
			if err != nil {
				return nil, err
			}
			event.EventId = eventId
			event.PartitionId = partitionId
			events = append(events, &event)
		}
	}
	if len(events) == 0 {
		return nil, errors.New("incomplete scan")
	}
	if events[len(events)-1].EventId != to-1 {
		return nil, errors.New("incomplete scan")
	}
	return events, nil
}

func getCommandResponse(ctx context.Context, partitionId uint64, entityType string, eventId uint64) ([]byte, error) {
	encodedEvent, err := kvstore.Get(ctx, partitionId, entityType, eventId)
	if err != nil {
		return nil, err
	}
	var event eventForGetCommandResponse
	err = eventJson.Unmarshal(encodedEvent, &event)
	if err != nil {
		return nil, err
	}
	return event.CommandResponse, err
}
