package runtime

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_thrift(t *testing.T) {
	should := require.New(t)
	schemas, err := ThriftSchemas(`
	struct Request {
		1: required string world;
	}
	struct Response {
		2: required string hello;
		3: i64 i64_field;
		4: float64 float64_field;
	}
	`)
	should.Nil(err)
	obj := NewObject()
	obj.Schema = schemas["Response"]
	obj.Set("i64_field", 100)
	obj.Set("float64_field", 20.0)
	should.Panics(func() {
		obj.Set("hello", 1)
	})
}

func Test_validate_direct_field(t *testing.T) {
	should := require.New(t)
	obj := NewObject()
	obj.Schema = &Schema{
		Fields: map[string]*Schema{
			"hello": mustBeString,
		},
	}
	should.Panics(func() {
		obj.Set("hello", 1)
	})
}

func Test_validate_indirect_field(t *testing.T) {
	should := require.New(t)
	obj := NewObject()
	obj.Schema = &Schema{
		Fields: map[string]*Schema{
			"hello": {
				Fields: map[string]*Schema{
					"world": mustBeString,
				},
			},
		},
	}
	should.Panics(func() {
		obj.Set("hello", NewObject("world", 1))
	})
}

func Test_validate_list(t *testing.T) {
	should := require.New(t)
	schemas, err := ThriftSchemas(`
	struct Request {
		1: list<string> words;
	}
	`)
	should.Nil(err)
	obj := NewObject()
	obj.Schema = schemas["Request"]
	obj.Set("words", NewList("hello", "world"))
	should.Panics(func() {
		obj.Set("words", 1)
	})
	should.Panics(func() {
		obj.Set("words", NewList(1))
	})
}


func Test_validate_list_of_object(t *testing.T) {
	should := require.New(t)
	schemas, err := ThriftSchemas(`
	struct taken {
		1: string taken_by
	}
	struct Entity {
		1: list<taken> takens
	}`)
	should.Nil(err)
	obj := NewObject()
	obj.Schema = schemas["Entity"]
	obj.Set("takens", NewList(NewObject("taken_by", "taowen")))
	should.Panics(func() {
		obj.Set("takens", 1)
	})
	should.Panics(func() {
		obj.Set("takens", NewList(1))
	})
}

func Test_validate_map(t *testing.T) {
	should := require.New(t)
	schemas, err := ThriftSchemas(`
	struct Request {
		1: map<string, string> words;
	}
	`)
	should.Nil(err)
	obj := NewObject()
	obj.Schema = schemas["Request"]
	obj.Set("words", NewObject("hello", "world"))
	should.Panics(func() {
		obj.Set("words", 1)
	})
	should.Panics(func() {
		obj.Set("words", NewList(1))
	})
}