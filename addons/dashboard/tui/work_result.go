package tui

import (
	"fmt"
	"net/url"
	"time"
)

type WorkResult struct {
	Err error

	ParentURL url.URL
	URL       url.URL

	NumberOfWorkers int
	WorkerID        int

	ResponseSize int
	StatusCode   int
	StartTime    time.Time
	EndTime      time.Time
	ContentType  string
}

func (w WorkResult) String() string {
	return fmt.Sprintf("#%03d: %03d %9s %15s %20s",
		w.WorkerID,
		w.StatusCode,
		fmt.Sprintf("%d", w.ResponseSize),
		fmt.Sprintf("%f ms", w.GetResponseTime().Seconds()*1000),
		w.URL.String(),
	)
}

func (w WorkResult) GetError() error {
	return w.Err
}

func (w WorkResult) GetParentURL() url.URL {
	return w.ParentURL
}

func (w WorkResult) GetURL() url.URL {
	return w.URL
}

func (w WorkResult) GetSize() int {
	return w.ResponseSize
}

func (w WorkResult) GetStatusCode() int {
	return w.StatusCode
}

func (w WorkResult) GetStartTime() time.Time {
	return w.StartTime
}

func (w WorkResult) GetEndTime() time.Time {
	return w.EndTime
}

func (w WorkResult) GetResponseTime() time.Duration {
	return w.EndTime.Sub(w.StartTime)
}

func (w WorkResult) GetContentType() string {
	return w.ContentType
}

func (w WorkResult) GetWorkerID() int {
	return w.WorkerID
}

func (w WorkResult) GetNumberOfWorkers() int {
	return w.NumberOfWorkers
}
