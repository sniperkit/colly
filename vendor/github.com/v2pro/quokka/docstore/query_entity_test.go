package docstore

import (
	"testing"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"net/http"
	"io/ioutil"
)

func Test_entity_missing(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil)
	StartNode(context.TODO(), "127.0.0.1:9879")
	queryAndExpectError(t, "http://127.0.0.1:9879/docstore/entities/user/123", ErrEntityMissing)
}

func Test_query_entity_cache(t *testing.T) {
	should := require.New(t)
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, "",nil, nil).AddCommand("noop",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		},"", nil, nil)
	StartNode(context.TODO(), "127.0.0.1:9879")
	queryAndExpectError(t, "http://127.0.0.1:9879/docstore/entities/user/123", ErrEntityMissing)
	execAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/user/create", "EntityId", "123")
	resp := queryAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/entities/user/123")
	should.Equal(1, jsoniter.Get(resp, "entity_version").ToInt())
	execAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/user/noop", "EntityId", "123")
	resp = queryAndExpectSuccess(t, "http://127.0.0.1:9879/docstore/entities/user/123")
	should.Equal(2, jsoniter.Get(resp, "entity_version").ToInt())
}

func queryAndExpectError(t *testing.T, url string, errorNumber int) {
	should := require.New(t)
	resp, err := http.Get(url)
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal(errorNumber, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}

func queryAndExpectSuccess(t *testing.T, url string) []byte {
	should := require.New(t)
	resp, err := http.Get(url)
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
	return body
}
