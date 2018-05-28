package runtime

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_searchLine(t *testing.T) {
	should := require.New(t)
	line, start, end := searchLine("ab\ncd\nef", 0)
	should.Equal("ab", line)
	should.Equal(0, start)
	should.Equal(2, end)
	line, start, end = searchLine("ab\ncd\nef", 1)
	should.Equal("cd", line)
	should.Equal(3, start)
	should.Equal(5, end)
	line, start, end = searchLine("ab\ncd\nef", 2)
	should.Equal("ef", line)
	should.Equal(6, start)
	should.Equal(8, end)
}