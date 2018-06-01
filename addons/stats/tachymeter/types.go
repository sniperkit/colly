package tachymeter

// timeslice holds time.Duration values.
type timeSlice []time.Duration

// Satisfy sort for timeSlice.
func (p timeSlice) Len() int           { return len(p) }
func (p timeSlice) Less(i, j int) bool { return int64(p[i]) < int64(p[j]) }
func (p timeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Histogram is a map["low-high duration"]count of events that
// fall within the low-high time duration range.
type Histogram []map[string]uint64

type ranking struct {
	label     string
	startedAt time.Time
	endedAt   time.Time
	duration  time.Duration
	err       bool
}

// timeRank holds time.Duration values.
type timeRank []ranking

// Satisfy sort for timeRank.
func (p timeRank) Len() int           { return len(p) }
func (p timeRank) Less(i, j int) bool { return int64(p[i].duration) < int64(p[j].duration) } //  p[i].duration.Before(p[j].duration) }
func (p timeRank) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// type Rank []map[string]uint64
// type Events map[string]bool

// Metrics holds the calculated outputs
// produced from a Tachymeter sample set.
type Metrics struct {
	Time struct { // All values under Time are selected entirely from events within the sample window.
		Cumulative time.Duration // Cumulative time of all sampled events.
		HMean      time.Duration // Event duration harmonic mean.
		Avg        time.Duration // Event duration average.
		P50        time.Duration // Event duration nth percentiles ..
		P75        time.Duration
		P95        time.Duration
		P99        time.Duration
		P999       time.Duration
		Long5p     time.Duration // Average of the longest 5% event durations.
		Short5p    time.Duration // Average of the shortest 5% event durations.
		Max        time.Duration // Highest event duration.
		Min        time.Duration // Lowest event duration.
		Range      time.Duration // Event duration range (Max-Min).
	}

	Rank struct {
		Cumulative time.Duration // Cumulative time of all sampled events.
		HMean      time.Duration // Event duration harmonic mean.
		Avg        time.Duration // Event duration average.
		P50        time.Duration // Event duration nth percentiles ..
		P75        time.Duration
		P95        time.Duration
		P99        time.Duration
		P999       time.Duration
		Long5p     time.Duration // Average of the longest 5% event durations.
		Short5p    time.Duration // Average of the shortest 5% event durations.
		Max        string
		Min        string
		Range      time.Duration // Event duration range (Max-Min).
	}

	/*
		Rank struct {
			P50     []map[time.Duration]string // Event duration nth percentiles ..
			P75     []map[time.Duration]string
			P95     []map[time.Duration]string
			P99     []map[time.Duration]string
			P999    []map[time.Duration]string
			Long5p  []map[time.Duration]string
			Short5p []map[time.Duration]string
			Max     []map[time.Duration]string
			Min     []map[time.Duration]string
		}
	*/

	Rate struct {
		// Per-second rate based on event duration avg. via Metrics.Cumulative / Metrics.Samples.
		// If SetWallTime was called, event duration avg = wall time / Metrics.Count
		Second float64
	}

	Abuse struct {
		Cumulative  time.Duration // Cumulative time of all sampled events.
		HMean       time.Duration // Event duration harmonic mean.
		Avg         time.Duration // Event duration average.
		TriggeredAt time.Time
		Second      float64
		Count       int
	}

	Events              map[string]bool
	Histogram           *Histogram    // Frequency distribution of event durations in len(Histogram) buckets of HistogramBucketSize.
	HistogramBucketSize time.Duration // The width of a histogram bucket in time.
	Samples             int           // Number of events included in the sample set.
	Count               int           // Total number of events observed.
	Wall                time.Duration
}
