package main

import (
	queue "github.com/sniperkit/colly/pkg/queue"
	storage "github.com/sniperkit/colly/pkg/storage"
)

// collector queue
var (
	collectorQueueThreads int             = 4
	collectorQueueWorkers int             = 4
	collectorQueueMaxSize int             = 10000
	collectorQueueStorage storage.Storage // Storage interface
)

var (
	cq *queue.Queue // collector's queue instance
)
