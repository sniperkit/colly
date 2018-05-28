package main

import (
	"time"
)

type cacheEngine struct {
	ttl time.Duration
	connectRemote
}
