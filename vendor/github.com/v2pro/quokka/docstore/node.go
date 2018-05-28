package docstore

import (
	"net/http"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net"
	"time"
	"bytes"
	"github.com/v2pro/quokka/kvstore"
	"strings"
	"github.com/v2pro/plz/countlog"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/plz/concurrent"
	"context"
	"sync"
	"errors"
)

// node.go is the process running endpoint
// it will forward to other nodes, or itself (command_processor)
// node.go will delegate the state to topo.go/cluster.go

var commandProcessors = make([]map[string]*commandProcessor, kvstore.PartitionsCount)
var commandProcessorsMutexs = make([]*sync.Mutex, kvstore.PartitionsCount)

func init() {
	for i := 0; i < kvstore.PartitionsCount; i++ {
		commandProcessorsMutexs[i] = &sync.Mutex{}
	}
}

var thisNodeAddr string
var thisNodeStarted bool
var thisNodeHttpServer *http.Server
var thisNodeExecutor *concurrent.UnboundedExecutor

func StartNode(ctx context.Context, addr string) {
	thisNodeExecutor = concurrent.NewUnboundedExecutor()
	entityCache := newEntityCache()
	var mux = &http.ServeMux{}
	mux.HandleFunc("/docstore/ping", ping)
	mux.HandleFunc("/docstore/entities/", entityCache.queryEntity)
	mux.HandleFunc("/docstore/exec", exec)
	mux.HandleFunc("/docstore/", exec)
	// will promote servers to master if needed
	thisNodeHttpServer = &http.Server{Addr: addr, Handler: mux}
	thisNodeExecutor.Go(func(ctx context.Context) {
		thisNodeHttpServer.ListenAndServe()
	})
	if !isAlive(addr) {
		countlog.Error("event!node.failed to listen http", "addr", addr)
		return
	}
	err := publishNode(ctx, addr)
	if err != nil {
		countlog.Error("event!node.failed to publish", "addr", addr)
		return
	}
	countlog.Info("event!node.started", "addr", addr)
}

func StopNode(ctx context.Context) {
	countlog.Info("event!node.stopping")
	if err := thisNodeHttpServer.Shutdown(ctx); err != nil {
		countlog.Error("event!node.failed to shutdown http server", "err", err)
	}
	thisNodeExecutor.StopAndWait(ctx)
	commandProcessors = make([]map[string]*commandProcessor, kvstore.PartitionsCount)
	topo = make(topology, kvstore.PartitionsCount)
	thisNodeStarted = false
	thisNodeAddr = ""
	countlog.Info("event!node.stopped")
	thisNodeExecutor = nil
}

func isAlive(addr string) bool {
	for i := 0; i < 3; i++ {
		_, err := http.Get("http://" + addr + "/docstore/ping")
		if err != nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		return true
	}
	return false
}

func ping(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Write([]byte("pong"))
}

func exec(respWriter http.ResponseWriter, req *http.Request) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!exec.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	var cmd command
	err = jsoniter.Unmarshal(body, &cmd)
	if err != nil {
		respWriter.Write(replyError(err))
		return
	}
	if !strings.HasSuffix(req.URL.Path, "/exec") {
		path := strings.Trim(req.URL.Path, "/")
		segments := strings.Split(path, "/")
		if len(segments) >= 2 {
			cmd.CommandType = segments[len(segments)-1]
			cmd.EntityType = segments[len(segments)-2]
		}
	}
	partitionId := kvstore.HashToPartition(cmd.EntityId)
	var commandResp []byte
	master, isLocal := chooseCommandTarget(partitionId, &cmd, req)
	countlog.Trace("event!node.choose command target",
		"target", master,
		"isLocal", isLocal,
		"cmd", &cmd)
	if master == "" {
		respWriter.Write(replyError(errors.New("master is not available for this partition")))
		return
	}
	if isLocal {
		commandProcessor, err := getOrCreateCommandProcessor(req.Context(), partitionId, cmd.EntityType)
		if err != nil {
			respWriter.Write(replyError(err))
			return
		}
		commandResp = commandProcessor.delegatedExec(&cmd, time.Second)
	} else {
		commandResp = forwardCommand(master, &cmd)
	}
	var resp respWithErrno
	jsoniter.Unmarshal(commandResp, &resp)
	resp.VerifyMaster(partitionId, master)
	respWriter.Write(commandResp)
}

type respWithErrno struct {
	Errno     int    `json:"errno"`
	Committed bool   `json:"committed"`
	HandledBy string `json:"handled_by"`
}

func (resp *respWithErrno) VerifyMaster(partitionId uint64, master string) {
	if master == "" {
		return
	}
	if resp.Errno == ErrEventLogConflict {
		countlog.Warn("event!node.find master failed to commit event log",
			"oldMaster", master,
			"partitionId", partitionId)
		topo.clearMaster(partitionId)
		return
	}
	isCommitted := resp.Errno == 0 || resp.Committed
	if isCommitted && resp.HandledBy != master {
		countlog.Warn("event!node.find the actual master is different",
			"oldMaster", master,
			"actualMaster", resp.HandledBy,
			"partitionId", partitionId)
		topo.clearMaster(partitionId)
		return
	}
	return
}

func getOrCreateCommandProcessor(ctx context.Context, partitionId uint64, entityType string) (*commandProcessor, error) {
	commandProcessorsMutexs[partitionId].Lock()
	defer commandProcessorsMutexs[partitionId].Unlock()
	if commandProcessors[partitionId] == nil {
		commandProcessors[partitionId] = map[string]*commandProcessor{}
	}
	commandProcessor := commandProcessors[partitionId][entityType]
	if commandProcessor != nil {
		return commandProcessor, nil
	}
	processor := newCommandProcessor(partitionId, entityType)
	err := processor.init(ctx)
	if err != nil {
		countlog.Error("event!node.failed to init command processor",
			"partitionId", partitionId,
			"entityType", entityType,
			"err", err)
		return nil, err
	}
	thisNodeExecutor.Go(func(ctx context.Context) {
		processor.executeInBackground(ctx)
	})
	thisNodeExecutor.Go(processor.lookupSyncer.syncInBackground)
	countlog.Info("event!node.created new command processor",
		"partitionId", partitionId,
		"entityType", entityType)
	commandProcessors[partitionId][entityType] = processor
	return processor, nil
}

func chooseCommandTarget(partitionId uint64, cmd *command, req *http.Request) (string, bool) {
	if cmd.IsPromoting {
		return "", true
	}
	ctx := req.Context()
	localAddr, _ := ctx.Value(http.LocalAddrContextKey).(*net.TCPAddr)
	partition := topo.get(partitionId)
	cmd.IsPromoting = partition.isPromoting
	master := partition.master
	// TODO: when master is empty, forward to any other node
	// TODO: prevent forward cycle, fail after 3 hops
	if localAddr == nil {
		return master, false
	}
	if master != localAddr.String() {
		return master, false
	}
	return master, true
}

func forwardCommand(target string, cmd *command) []byte {
	if cmd.ForwardedTimes >= 3 {
		return replyError(withErrorNumber(errors.New("forwarded too many times"), ErrForwardedTooManyTimes))
	}
	cmd.ForwardedTimes += 1
	req, err := runtime.Json.Marshal(cmd)
	if err != nil {
		return replyError(err)
	}
	url := "http://" + target + "/docstore/exec"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return replyError(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return replyError(err)
	}
	return body
}
