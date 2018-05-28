package main

type config struct {
	cache   bool
	depth   int
	workers int
	debug
}

type worker struct {
	max     uint
	timeout uint
	debug
}

type debug struct {
	enabled bool
	forward bool
}

type logger struct {
	enabled bool
	engine  string
	connectRemote
}

type connectRemote struct {
	address       string
	port          string
	password      string
	tls           bool
	persist       bool
	autoReconnect bool
	maxReconnect  uint
}
