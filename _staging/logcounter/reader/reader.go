package main

import (
	"os"

	"github.com/timjchin/logcounter"
	"github.com/timjchin/logcounter/reader/server"
)

func main() {
	l := logcounter.NewLogCounter(&logcounter.LogCounterConfig{
		NumResults: 10,
		PrintInput: true,
	})
	s := server.NewServer(&server.ServerConfig{}, l)
	go func() {
		s.Start()
	}()

	l.ParseReader(os.Stdin)
}
