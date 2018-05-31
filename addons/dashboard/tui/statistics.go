package tui

import (
	"fmt"
	"sync"
	"time"
	// pp "github.com/sniperkit/xutil/plugin/debug/pp"
)

var (
	stats                        Statistics
	AllStatisticsHaveBeenUpdated chan bool
	AllURLsHaveBeenVisited       chan bool
)

func init() {
	stats = Statistics{
		lock: sync.RWMutex{},
		numberOfRequestsByStatusCode:  make(map[int]int),
		numberOfRequestsByContentType: make(map[string]int),
	}
}

func UpdateStatistics(w WorkResult) {
	go stats.Add(w)
}

type Snapshot struct {
	// time
	timestamp           time.Time
	timeSinceStart      time.Duration
	averageResponseTime time.Duration

	// counters
	numberOfWorkers              int
	totalNumberOfRequests        int
	numberOfSuccessfulRequests   int
	numberOfUnsuccessfulRequests int
	numberOfRequestsPerSecond    float64

	numberOfRequestsByStatusCode  map[int]int
	numberOfRequestsByContentType map[string]int

	// size
	totalSizeInBytes   int
	averageSizeInBytes int
}

func (s Snapshot) Timestamp() time.Time {
	return s.timestamp
}

func (s Snapshot) NumberOfWorkers() int {
	return s.numberOfWorkers
}

func (s Snapshot) NumberOfErrors() int {
	return s.numberOfUnsuccessfulRequests
}

func (s Snapshot) TotalNumberOfRequests() int {
	return s.totalNumberOfRequests
}

func (s Snapshot) TotalSizeInBytes() int {
	return s.totalSizeInBytes
}

func (s Snapshot) AverageSizeInBytes() int {
	return s.averageSizeInBytes
}

func (s Snapshot) AverageResponseTime() time.Duration {
	return s.averageResponseTime
}

func (s Snapshot) RequestsPerSecond() float64 {
	return s.numberOfRequestsPerSecond
}

type Statistics struct {
	lock sync.RWMutex

	rawResults  []WorkResult
	snapShots   []Snapshot
	logMessages []string

	startTime time.Time
	endTime   time.Time

	totalResponseTime time.Duration

	numberOfWorkers               int
	numberOfRequests              int
	numberOfSuccessfulRequests    int
	numberOfUnsuccessfulRequests  int
	numberOfRequestsByStatusCode  map[int]int
	numberOfRequestsByContentType map[string]int

	totalSizeInBytes int
}

func (s *Statistics) Add(w WorkResult) Snapshot {
	// update the raw results
	s.lock.Lock()
	defer s.lock.Unlock()

	/*
		log.Println("Add new WorkResult URL=", w.GetURL(),
			"NumberOfWorkers=", w.GetNumberOfWorkers(),
			"WorkerID=", w.GetWorkerID(),
			"ResponseSize=", w.GetSize(),
			"StatusCode=", w.GetStatusCode(),
			"StartTime=", w.GetStartTime(),
			"EndTime=", w.GetEndTime(),
			"ContentType=", w.GetContentType(),
		)
	*/

	s.rawResults = append(s.rawResults, w)

	// initialize start and end time
	if s.numberOfRequests == 0 {
		s.startTime = w.GetStartTime()
		s.endTime = w.GetEndTime()
	}

	// start time
	if w.GetStartTime().Before(s.startTime) {
		s.startTime = w.GetStartTime()
	}

	// end time
	if w.GetEndTime().After(s.endTime) {
		s.endTime = w.GetEndTime()
	}

	// update the total number of requests
	s.numberOfRequests = len(s.rawResults)

	// is successful
	if w.GetStatusCode() > 199 && w.GetStatusCode() < 400 {
		s.numberOfSuccessfulRequests += 1
	} else {
		s.numberOfUnsuccessfulRequests += 1
	}

	// number of workers
	s.numberOfWorkers = w.GetNumberOfWorkers()

	// number of requests by status code
	s.numberOfRequestsByStatusCode[w.GetStatusCode()] += 1

	// number of requests by content type
	s.numberOfRequestsByContentType[w.GetContentType()] += 1

	// update the total duration
	responseTime := w.GetEndTime().Sub(w.GetStartTime())
	s.totalResponseTime += responseTime

	// size
	s.totalSizeInBytes += w.GetSize()
	averageSizeInBytes := s.totalSizeInBytes / s.numberOfRequests

	// average response time
	averageResponseTime := time.Duration(s.totalResponseTime.Nanoseconds() / int64(s.numberOfRequests))

	// number of requests per second
	requestsPerSecond := float64(s.numberOfRequests) / s.endTime.Sub(s.startTime).Seconds()

	// log messages
	s.logMessages = append(s.logMessages, w.String())

	// create a snapshot
	snapShot := Snapshot{
		// times
		timestamp:           w.GetEndTime(),
		averageResponseTime: averageResponseTime,

		// counters
		numberOfWorkers:               s.numberOfWorkers,
		totalNumberOfRequests:         s.numberOfRequests,
		numberOfSuccessfulRequests:    s.numberOfSuccessfulRequests,
		numberOfUnsuccessfulRequests:  s.numberOfUnsuccessfulRequests,
		numberOfRequestsPerSecond:     requestsPerSecond,
		numberOfRequestsByStatusCode:  s.numberOfRequestsByStatusCode,
		numberOfRequestsByContentType: s.numberOfRequestsByContentType,

		// size
		totalSizeInBytes:   s.totalSizeInBytes,
		averageSizeInBytes: averageSizeInBytes,
	}

	// pp.Println(snapShot)

	s.snapShots = append(s.snapShots, snapShot)
	return snapShot
}

func (s *Statistics) LastSnapshot() Snapshot {
	s.lock.RLock()
	defer s.lock.RUnlock()

	lastSnapshotIndex := len(s.snapShots) - 1
	if lastSnapshotIndex < 0 {
		return Snapshot{}
	}

	return s.snapShots[lastSnapshotIndex]
}

func (s *Statistics) LastLogMessages(count int) []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	messages, err := getLatestLogMessages(s.logMessages, count)
	if err != nil {
		panic(err)
	}

	return messages
}

func getLatestLogMessages(messages []string, count int) ([]string, error) {
	if count < 0 {
		return nil, fmt.Errorf("The count cannot be negative")
	}

	numberOfMessges := len(messages)
	if count == numberOfMessges {
		return messages, nil
	}

	if count < numberOfMessges {
		return messages[numberOfMessges-count:], nil
	}

	if count > numberOfMessges {
		fillLines := make([]string, count-numberOfMessges)
		return append(fillLines, messages...), nil
	}
	panic("Unreachable")
}
