package main

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// cache storage errors
var (
	errInvalidCacheDuration = errors.New("Invalid cache duration...")
	errInvalidCacheTTL      = errors.New("Invalid cache TTL...")
	errInvalidCacheTimeUnit = errors.New("Invalid cache time unit. available: second, hour, day")
	errInvalidCacheBackend  = errors.New("Invalid cache backend. Available: " + strings.Join(defaultCacheBackends, ",") + ".")
)

// colly queue processing errors
var (
	errInvalidQueueThreads     = errors.New("Invalid queue consumer threads count. Must be superior or equal to 0.")
	errInvalidQueueBackend     = errors.New("Unkown queue storage backend name. Available: inmemory, redis, sqlite3, badger, mysql, postgres.")
	errInvalidQueueMaxSize     = errors.New("Invalid queue max size value. Must be superior or equal to 0.")
	errLocalFileStat           = errors.New("File not found.")
	errLocalFileOpen           = errors.New("Could not open the filepath")
	errInvalidRemoteStatusCode = errors.New("errInvalidRemoteStatusCode")
)

// e returns an error, prefixed with the name of the function that triggered it. Originally by StackOverflow user svenwltr:
// http://stackoverflow.com/a/38551362/199475
func e(err error) error {
	pc, _, _, _ := runtime.Caller(2)

	fr := runtime.CallersFrames([]uintptr{pc})
	namer, _ := fr.Next()
	name := namer.Function

	if !FullyQualifiedPath {
		fn := strings.Split(name, "/")
		if len(fn) > 0 {
			return fmt.Errorf("%s: %s", fn[len(fn)-1], err.Error())
		}
	}

	return fmt.Errorf("%s: %s", name, err.Error())
}

// Err consumes an error, a string, or nil, and produces an error message prefixed with the name of the function that called it (or nil).
func Err(err interface{}) error {
	switch o := err.(type) {
	case string:
		return e(fmt.Errorf("%s", o))
	case error:
		return e(o)
	default:
		return nil
	}
}
