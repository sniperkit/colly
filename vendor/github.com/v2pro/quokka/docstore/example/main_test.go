package example

import (
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/kvstore/memkv"
	"github.com/v2pro/quokka/docstore"
	"testing"
	"context"
	"github.com/v2pro/quokka/docstore/runtime"
	"net/http"
	"bytes"
	"io/ioutil"
)

func TestMain(m *testing.M) {
	memkv.ResetKVStore()
	docstore.StartNode(context.TODO(), "127.0.0.1:9865")
	m.Run()
	docstore.StopNode(context.TODO())
}

func execAndExpectSuccess(t *testing.T, url string, kv ...interface{}) []byte {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post("http://127.0.0.1:9865/docstore"+url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
	return body
}

func execAndExpectError(t *testing.T, url string, errorNumber int, kv ...interface{}) []byte {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post("http://127.0.0.1:9865/docstore"+url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal(errorNumber, jsoniter.Get(body, "errno").MustBeValid().ToInt())
	return body
}

func queryAndExpectSuccess(t *testing.T, url string) []byte {
	should := require.New(t)
	resp, err := http.Get("http://127.0.0.1:9865/docstore" + url)
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
	return body
}
