package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"net/http"
	"time"
	"io/ioutil"
	"github.com/v2pro/quokka/docstore/runtime"
	"bytes"
	"github.com/v2pro/quokka/kvstore"
	"os"
	"context"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	if thisNodeExecutor != nil {
		StopNode(context.TODO())
	}
	os.Exit(exitCode)
}

func Test_first_event(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "", nil, nil)
	StartNode(context.TODO(), "127.0.0.1:9879")
	execAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/user/create", "EntityId", "123")
}

func Test_duplicated_entity(t *testing.T) {
	should := require.New(t)
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil)
	ctx := context.TODO()
	StartNode(ctx, "127.0.0.1:9879")
	execAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/user/create", "EntityId", "123")
	StopNode(ctx)
	StartNode(ctx, "127.0.0.1:9879")
	execAndExpectError(t, "http://127.0.0.1:9879/docstore/user/create", ErrDuplicatedEntity,
		"EntityId", "123")
	event2 := debugGet(kvstore.HashToPartition("123"), "user", 2)
	should.Equal(1, jsoniter.Get(event2, "b").ToInt())
}

func Test_append_more_event_log_after_restart(t *testing.T) {
	should := require.New(t)
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil).AddCommand("noop",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil)
	StartNode(context.TODO(), "127.0.0.1:9879")
	execAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/user/create", "EntityId", "123")
	StopNode(context.TODO())
	StartNode(context.TODO(), "127.0.0.1:9879")
	execAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/user/noop", "EntityId", "123")
	should.Equal(1, jsoniter.Get(debugGet(kvstore.HashToPartition("123"), "user", 2), "b").ToInt())
}

func Test_command_before_create(t *testing.T) {
	should := require.New(t)
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil).AddCommand("noop",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil)
	StartNode(context.TODO(), "127.0.0.1:9879")
	execAndExpectError(t, "http://127.0.0.1:9879/docstore/user/noop", ErrMustCreateFirst,
		"EntityId", "123")
	event1 := debugGet(kvstore.HashToPartition("123"), "user", 1)
	should.Equal(ErrMustCreateFirst, jsoniter.Get(event1, "p", "errno").ToInt())
}

func Test_key_conflict_lost_master(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil).AddCommand("noop",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil)
	StartNode(context.TODO(), "127.0.0.1:9879")
	execAndExpectError(t, "http://127.0.0.1:9879/docstore/user/noop", ErrMustCreateFirst,
		"EntityId", "123")
	event, _ := eventJson.Marshal(&Event{
		EntityId:        "123",
		BaseEventId:     1,
		Version:         2,
		CommandId:       "",
		CommandType:     "create",
		CommandRequest:  []byte("{}"),
		CommandResponse: []byte("{}"),
		State:           []byte("null"),
		Delta:           []byte("{}"),
		CommittedAt:     time.Now().UnixNano(),
	})
	kvstore.Append(context.TODO(), kvstore.HashToPartition("123"), "user", 2, event)
	execAndExpectError(t, "http://127.0.0.1:9879/docstore/user/noop", ErrEventLogConflict,
		"EntityId", "123")
	execAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/user/noop",
		"EntityId", "123")
}

func execAndExpectSuccess(t *testing.T, url string, kv ...interface{}) {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}

func execAndExpectError(t *testing.T, url string, errorNumber int, kv ...interface{}) {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal(errorNumber, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}
