package tui

import (
	"time"

	"github.com/sasile/gohistogram"
)

type CollectorLatency struct {
	WeighHist *gohistogram.NumericHistogram
	ch_values chan time.Duration
}

func (self *CollectorLatency) New(n int, alpha float64) chan time.Duration {
	self.WeighHist = gohistogram.NewHistogram(50)
	self.ch_values = make(chan time.Duration, 400000)
	go func() {
		for v := range self.ch_values {
			self.WeighHist.Add(float64(v.Nanoseconds() / 1000))
		}
	}()
	return self.ch_values
}

func (self *CollectorLatency) Add(v time.Duration) {
	self.ch_values <- v
}

func (self *CollectorLatency) Get() ([]string, []int) {
	return self.WeighHist.BarArray()
}

func (self *CollectorLatency) GetResults() ([]string, []float64) {
	return self.WeighHist.GetHistAsArray()

}

func (self *CollectorLatency) GetQuantile(q float64) float64 {
	return self.WeighHist.CDF(q)

}

func (self *CollectorLatency) GetCount() float64 {
	return self.WeighHist.Count()

}

func (self *CollectorLatency) String() string {
	return self.WeighHist.String()
}
