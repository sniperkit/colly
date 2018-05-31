package main

import (
	"encoding/csv"
	"strings"
	// "time"
	"net/http"
	"os"

	"github.com/sniperkit/colly/pkg/queue"
	// "github.com/sniperkit/colly/pkg/helper"
	// "github.com/sniperkit/colly/pkg/storage"
	// pp "github.com/sniperkit/xutil/plugin/debug/pp"

	res "github.com/sniperkit/colly/addons/storage/external/redis"
	sq3 "github.com/sniperkit/colly/addons/storage/external/sqlite3"
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

var (
	queueConsumerThreads int = 2
	queueMaxSize         int = 100000
	// queueStorage         storage.Storage // Storage interface
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

	case "sqlite":
		fallthrough
	case "sqlite3":
		/*
			queueStorage = &sq3.Storage{
				Filename: "./shared/datastore/queue.db",
			}
		*/

		// create a request queue with 2 consumer threads
		q, err = queue.New(
			ct, // Number of consumer threads
			&sq3.Storage{
				Filename: "./shared/datastore/queue.db",
			},
		)

	case "redis":

		/*
			queueStorage = &res.Storage{
				Address:  "127.0.0.1:6379",
				Password: "",
				DB:       0,
				Prefix:   "job01",
			}
		*/

		// create a request queue with 2 consumer threads
		q, err = queue.New(
			ct, // Number of consumer threads
			&res.Storage{
				Address:  "127.0.0.1:6379",
				Password: "",
				DB:       0,
				Prefix:   "job01",
			},
		)

	//	case "badger":
	//		fallthrough

	case "inmemory":
		fallthrough

	default:
		// create a request queue with 2 consumer threads
		q, err = queue.New(
			ct, // Number of consumer threads
			&queue.InMemoryQueueStorage{
				MaxSize: s, // 10000, // Use default queue storage
			},
		)
		return
	}

	return
}

func linksFromCSV(filePath string) ([]string, error) {

	/*
		csvStream, err := NewStreamCSV(filePath, "by_key")
		if err != nil {
			log.Println("error could not create new CSV stream")
		}

		// csvStream.Buffer(csvStreamBuffer).SetColumnsKeys(0).SplitAt(5000)
		// lines.ReadAsync() // .Wait()
	*/

	isRemote := isRemoteURL(filePath)

	var reader *csv.Reader
	if !isRemote {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return nil, err
		}
		log.Infoln("reading file:", filePath)
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalln("failed to open file, error: ", err)
			return nil, err
		}
		defer file.Close()
		reader = csv.NewReader(file)

	} else {
		log.Infoln("loading remote:", filePath)
		resp, err := http.Get(filePath)
		if err != nil {
			log.Fatalln("failed to fetch content, error: ", err)
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			log.Fatalln("failed with status", resp.Status)
			return nil, err
		}
		reader = csv.NewReader(resp.Body)

	}

	lines := streamCsv(reader, csvStreamBuffer)
	links := make([]string, len(lines))
	for line := range lines {
		// pp.Println(line)
		links = append(links, line.GetByKey(0))
		log.Infoln("[LIST-ROW] col[0]=", line.GetByKey(0))
	}

	log.Infoln("links pre-loaded:", len(links))
	return links, nil

	/*
		for line := range lines {
			expiresAt, err := parseTimeStamp(line.GetByName("expiration_timestamp"))
			if err != nil {
				log.Errorln("[TSK-ERR] taskInfo, domain=", line.GetByName("domain"), "loc=", line.GetByName("loc"), "expiresTimestamp", line.GetByName("expiration_timestamp"))
				continue
			}

			now := time.Now()
			if now.After(expiresAt.Add(cacheTTL)) {
				log.Infoln("[TSK-ADD] task info, domain=", line.GetByName("domain"), "loc=", line.GetByName("loc"), "expiresAt=", expiresAt)
				continue
			}

			cuckflt.InsertUnique([]byte(line.GetByName("loc")))
		}
	*/

	// log.Warnln("[TSK-QUEUE] Number of requests to bypass: ", cuckflt.Count())
}
