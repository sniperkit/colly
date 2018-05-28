package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go/test"
)

func Test_decode_bool(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteBool(true)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(true, iter.ReadBool())
	}
}

func Test_unmarshal_bool(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteBool(true)
		var val bool
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(true, val)
	}
}

func Test_encode_bool(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteBool(true)
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(true, iter.ReadBool())
	}
}

func Test_marshal_bool(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(true)
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(true, iter.ReadBool())
	}
}
