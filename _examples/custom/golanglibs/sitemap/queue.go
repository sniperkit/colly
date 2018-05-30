package main

import (
	"strings"

	"github.com/sniperkit/colly/pkg/queue"
	// "github.com/sniperkit/colly/pkg/storage"
	// res "github.com/sniperkit/colly/addons/storage/external/redis"
	// sq3 "github.com/sniperkit/colly/addons/storage/external/sqlite3"
	// baq "github.com/sniperkit/colly/addons/storage/external/badger"
	// stq "github.com/sniperkit/colly/addons/storage/external/storm"
	// myq "github.com/sniperkit/colly/addons/storage/external/mysql"
	// moq "github.com/sniperkit/colly/addons/storage/external/mongo"
	// elq "github.com/sniperkit/colly/addons/storage/external/elastic"
	// shq "github.com/sniperkit/colly/addons/storage/external/sphinx"
	// caq "github.com/sniperkit/colly/addons/storage/external/cassandra"
)

// collector's queues defaults
var (
	defaultQueueConsumerThreads int    = 2
	defaultQueueMaxSize         int    = 100000
	defaultQueueStorageEngine   string = "InMemory" // Available: InMemory, Redis, SQlite3, Badger KV or Mysql
)

func initQueue(ct int, s int, b string) (q *queue.Queue, err error) {
	if ct < 0 {
		err = errInvalidQueueThreads
		return
	}
	b = strings.ToLower(b)
	if s < 0 {
		err = errInvalidQueueMaxSize
		return
	}

	switch b {
	case "inmemory":

		// create a request queue with 2 consumer threads
		q, err = queue.New(
			ct, // Number of consumer threads
			&queue.InMemoryQueueStorage{
				MaxSize: s, // 10000, // Use default queue storage
			},
		)

	/*
		case "sqlite":
			fallthrough
		case "sqlite3":
			// create a request queue with 2 consumer threads
			q, err = queue.New(
				ct, // Number of consumer threads

				&sq3.InMemoryQueueStorage{
					MaxSize: s, // 10000, // Use default queue storage
				},
			)

		case "redis":
			// create a request queue with 2 consumer threads
			q, err = queue.New(
				ct, // Number of consumer threads
				&res.InMemoryQueueStorage{
					MaxSize: s, // 10000, // Use default queue storage
				},
			)

		case "badger":
			fallthrough
	*/
	default:
		err = errInvalidQueueBackend
		return
	}

	return
}
