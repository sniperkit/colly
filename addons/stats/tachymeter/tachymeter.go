// Package tachymeter yields summarized data describing a series of timed events.
package tachymeter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Config holds tachymeter initialization parameters. Size defines the sample capacity.
// Note: Tachymeter is thread safe.
type Config struct {
	SampleSize       int  `default:'50'`
	HistogramBuckets int  `default:'10'`   // Histogram buckets.
	Safe             bool `default:'true'` // Deprecated. Flag held on to as to not break existing users.
}

// Tachymeter holds event durations
// and counts.
type Tachymeter struct {
	sync.Mutex
	size     uint64
	times    timeSlice
	ranks    timeRank
	count    uint64
	wallTime time.Duration
	hBuckets int
}

// New initializes a new Tachymeter.
func New(c *Config) *Tachymeter {
	var hSize int

	if c == nil {
		c = &Config{
			HBuckets: 10,
			Size:     50,
			Safe:     true,
		}
	}

	if c.HBuckets != 0 {
		hSize = c.HBuckets
	} else {
		hSize = 10
	}

	return &Tachymeter{
		size:     uint64(c.Size),
		ranks:    make(timeRank, c.Size),
		hBuckets: hSize,
	}
}

func Clone(t *Tachymeter) *Tachymeter {
	m.Lock()
	defer m.Unlock()
	return t
}

func (m *Tachymeter) WallTime(t time.Duration) {
	m.Lock()
	defer m.Unlock()
	return m.WallTime
}

func (m *Tachymeter) Size() uint64, int {
	m.Lock()
	defer m.Unlock()
	return m.size, int(m.size)
}

// Reset resets a Tachymeter instance for reuse.
func (m *Tachymeter) Reset() {
	// This lock is obviously not needed for
	// the m.Count update, rather to prevent a
	// Tachymeter reset while Calc is being called.
	m.Lock()
	atomic.StoreUint64(&m.Count, 0)
	m.Unlock()
}

// AddTime adds a time.Duration to Tachymeter.
func (m *Tachymeter) AddTime(label string, t time.Duration) {
	//	m.Times[(atomic.AddUint64(&m.Count, 1)-1)%m.Size] = t
	m.Ranks[(atomic.AddUint64(&m.Count, 1)-1)%m.Size] = ranking{duration: t, label: label}
}

// SetWallTime optionally sets an elapsed wall time duration.
// This affects rate output by using total events counted over time.
// This is useful for concurrent/parallelized events that overlap
// in wall time and are writing to a shared Tachymeter instance.
func (m *Tachymeter) SetWallTime(t time.Duration) {
	m.WallTime = t
}

// WriteHTML writes a histograph
// html file to the cwd.
func (m *Metrics) WriteHTML(p string) error {
	w := Timeline{}
	w.AddEvent(m)
	return w.WriteHTML(p)
}

// Dump prints a formatted Metrics output to console.
func (m *Metrics) Dump() {
	fmt.Println(m.String())
}

// String returns a formatted Metrics string.
func (m *Metrics) String() string {
	return fmt.Sprintf(`%d samples of %d events
Wall:		%s
Cumulative:	%s
HMean:		%s
Avg.:		%s
p50: 		%s
p75:		%s
p95:		%s
p99:		%s
p999:		%s
Long 5%%:	%s
Short 5%%:	%s
Max:		%s (%s)
Min:		%s (%s)
Range:		%s
Rate/sec.:	%.2f`,
		m.Samples,
		m.Count,
		m.Wall.String(),
		m.Rank.Cumulative,
		m.Rank.HMean,
		m.Rank.Avg,
		m.Rank.P50,
		m.Rank.P75,
		m.Rank.P95,
		m.Rank.P99,
		m.Rank.P999,
		m.Rank.Long5p,
		m.Rank.Short5p,
		m.Rank.Max,
		m.Rank.Max,
		m.Rank.Min,
		m.Rank.Min,
		m.Rank.Range,
		m.Rate.Second)
}

// JSON returns a *Metrics as
// a JSON string.
func (m *Metrics) JSON() string {
	j, _ := json.Marshal(m)

	return string(j)
}

// MarshalJSON defines the output formatting
// for the JSON() method. This is exported as a
// requirement but not intended for end users.
func (m *Metrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Time struct {
			Cumulative string
			HMean      string
			Avg        string
			P50        string
			P75        string
			P95        string
			P99        string
			P999       string
			Long5p     string
			Short5p    string
			Max        string
			Min        string
			Range      string
		}
		Rate struct {
			Second float64
		}
		Samples   int
		Count     int
		Histogram *Histogram
		Wall      string
	}{
		Time: struct {
			Cumulative string
			HMean      string
			Avg        string
			P50        string
			P75        string
			P95        string
			P99        string
			P999       string
			Long5p     string
			Short5p    string
			Max        string
			Min        string
			Range      string
		}{
			Cumulative: m.Time.Cumulative.String(),
			HMean:      m.Time.HMean.String(),
			Avg:        m.Time.Avg.String(),
			P50:        m.Time.P50.String(),
			P75:        m.Time.P75.String(),
			P95:        m.Time.P95.String(),
			P99:        m.Time.P99.String(),
			P999:       m.Time.P999.String(),
			Long5p:     m.Time.Long5p.String(),
			Short5p:    m.Time.Short5p.String(),
			Max:        m.Time.Max.String(),
			Min:        m.Time.Min.String(),
			Range:      m.Time.Range.String(),
		},
		Rate: struct{ Second float64 }{
			Second: m.Rate.Second,
		},
		Histogram: m.Histogram,
		Samples:   m.Samples,
		Count:     m.Count,
		Wall:      m.Wall.String(),
	})
}

// Dump prints a formatted histogram output to console
// scaled to a width of s.
func (h *Histogram) Dump(s int) {
	fmt.Println(h.String(s))
}

// String returns a formatted Metrics string scaled
// to a width of s.
func (h *Histogram) String(s int) string {
	if h == nil {
		return ""
	}

	var min, max uint64 = math.MaxUint64, 0
	// Get the histogram min/max counts.
	for _, bucket := range *h {
		for _, v := range bucket {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
	}

	var b bytes.Buffer

	// Build histogram string.
	for _, bucket := range *h {
		for k, v := range bucket {
			// Get the bar length.
			blen := scale(float64(v), float64(min), float64(max), 1, float64(s))
			line := fmt.Sprintf("%20s %s\n", k, strings.Repeat("-", int(blen)))
			b.WriteString(line)
		}
	}

	return b.String()
}

// Scale scales the input x with the input-min a0,
// input-max a1, output-min b0, and output-max b1.
func scale(x float64, a0, a1, b0, b1 float64) float64 {
	return (x-a0)/(a1-a0)*(b1-b0) + b0
}
