package main

import (
	"fmt"
	"net/url"
	"time"
)

type workResult struct {
	err error

	parentURL url.URL
	url       url.URL
	// parentURL string
	// url       string

	numberOfWorkers int
	workerID        int

	responseSize int
	statusCode   int
	startTime    time.Time
	endTime      time.Time
	contentType  string
}

func (w workResult) String() string {
	return fmt.Sprintf("#%03d: %03d %9s %15s %20s",
		w.workerID,
		w.statusCode,
		fmt.Sprintf("%d", w.responseSize),
		fmt.Sprintf("%f ms", w.ResponseTime().Seconds()*1000),
		w.url.String(),
	)
}

func (w workResult) Error() error {
	return w.err
}

func (w workResult) ParentURL() url.URL {
	return w.parentURL
}

func (w workResult) URL() url.URL {
	return w.url
}

func (w workResult) Size() int {
	return w.responseSize
}

func (w workResult) StatusCode() int {
	return w.statusCode
}

func (w workResult) StartTime() time.Time {
	return w.startTime
}

func (w workResult) EndTime() time.Time {
	return w.endTime
}

func (w workResult) ResponseTime() time.Duration {
	return w.endTime.Sub(w.startTime)
}

func (w workResult) ContentType() string {
	return w.contentType
}

func (w workResult) WorkerID() int {
	return w.workerID
}

func (w workResult) NumberOfWorkers() int {
	return w.numberOfWorkers
}
