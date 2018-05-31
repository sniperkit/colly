package histogram

import (
	"fmt"
	"sort"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type LatencyHistogram struct {
	ch_values chan time.Duration
	hist      map[int64]int
	count     int64
	size      int64
	wg        sync.WaitGroup
}

func (self *LatencyHistogram) Add(v time.Duration) {
	self.ch_values <- v
	self.size++
}

func (self *LatencyHistogram) Close() {
	close(self.ch_values)
}

func (self *LatencyHistogram) place(v int64) {
	self.hist[v/100]++
}

func (self *LatencyHistogram) New() chan time.Duration {
	log.Debugln("new latency hist")
	self.hist = make(map[int64]int)
	self.wg.Add(1)

	self.ch_values = make(chan time.Duration, 10000)

	go func() {
		defer self.wg.Done()
		for v := range self.ch_values {
			self.count++
			self.place(v.Nanoseconds() / 1000)
		}
	}()
	return self.ch_values
}

func (self *LatencyHistogram) GetResults() ([]string, []float64) {
	log.Debugln("get latency hist")
	self.wg.Wait()
	var keys []int
	for k := range self.hist {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	res_strings := []string{}
	res_values := []float64{}
	for _, k := range keys {
		v := self.hist[int64(k)]
		res_strings = append(res_strings, fmt.Sprintf("%5d - %5d",
			k*100, (k+1)*100))
		value := float64(v*100) / float64(self.count)
		res_values = append(res_values, value)
	}
	return res_strings, res_values
}

func (self *LatencyHistogram) GetHistMap() map[int64]int {
	self.wg.Wait()
	return self.hist
}
