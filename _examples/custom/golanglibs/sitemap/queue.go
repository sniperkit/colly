package main

import (
	"runtime"
	"strings"

	"github.com/sniperkit/colly/pkg/queue"
	// "github.com/sniperkit/colly/pkg/storage"

	res "github.com/sniperkit/colly/plugins/storage/external/redis"
	sq3 "github.com/sniperkit/colly/plugins/storage/external/sqlite3"
	// baq "github.com/sniperkit/colly/plugins/storage/external/badger"
	// stq "github.com/sniperkit/colly/plugins/storage/external/storm"
	// myq "github.com/sniperkit/colly/plugins/storage/external/mysql"
	// moq "github.com/sniperkit/colly/plugins/storage/external/mongo"
	// elq "github.com/sniperkit/colly/plugins/storage/external/elastic"
	// shq "github.com/sniperkit/colly/plugins/storage/external/sphinx"
	// caq "github.com/sniperkit/colly/plugins/storage/external/cassandra"
)

// collector's queues defaults
var (
	defaultQueueConsumerThreads uint   = 2
	defaultQueueStorageEngine   string = "InMemory" // Available: InMemory, Redis, SQlite3, Badger KV or Mysql
)

func initQueue(ct uint, s uint, b string) (q *queue.Queue, err error) {
	if ct < 0 {
		err = errInvalidQueueThreads
		return
	}
	b = strings.ToLower(b)

	switch b {
	case "inmemory":
		if errInvalidQueueMaxSize < 0 {
			err = errInvalidQueueMaxSize
			return
		}
		// create a request queue with 2 consumer threads
		q, err = queue.New(
			ct, // Number of consumer threads
			&queue.InMemoryQueueStorage{
				MaxSize: s, // 10000, // Use default queue storage
			},
		)

	case "sqlite":
		fallthrough
	case "sqlite3":
		// create a request queue with 2 consumer threads
		q, err = queue.New(
			consumerThreads, // Number of consumer threads

			&sq3.InMemoryQueueStorage{
				MaxSize: s, // 10000, // Use default queue storage
			},
		)

	case "redis":
		// create a request queue with 2 consumer threads
		q, err = queue.New(
			consumerThreads, // Number of consumer threads
			&res.InMemoryQueueStorage{
				MaxSize: s, // 10000, // Use default queue storage
			},
		)

	case "badger":
		fallthrough
	default:
		err = errInvalidQueueBackend
		return
	}

	return
}
