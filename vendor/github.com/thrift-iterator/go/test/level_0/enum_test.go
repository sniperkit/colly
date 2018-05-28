package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go/test"
	"github.com/thrift-iterator/go/test/level_0/enum_test"
)

func Test_unmarshal_enum(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI32(1)
		var val enum_test.Player
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(enum_test.Player_FLASH, val)
	}
}

func Test_marshal_enum(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(enum_test.Player_FLASH)
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(int32(1), iter.ReadInt32())
	}
}