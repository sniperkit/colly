package metric

import (
	"sync"
)

type MetricCollector struct {
	Pause   bool `default:'true'`
	Payload string
}

// NewMetricCollector creates a new Collector instance with cfg.Default configuration
func NewMetricCollector(options ...func(*MetricCollector)) *MetricCollector {
	mc := &MetricCollector{}
	mc.Init()
	for _, f := range options {
		f(mc)
	}
	// mc.parseSettingsFromEnv()
	return mc
}

// Init initializes the MetricCollector's private variables and sets default configuration for the MetricCollector
func (c *MetricCollector) Init() {
	c.wg = &sync.WaitGroup{}
	c.lock = &sync.RWMutex{}
	c.Pause = true
}

// SetPayload enables ...
func SetPayload(payload string) func(*MetricCollector) {
	return func(c *MetricCollector) {
		c.Payload = payload
	}
}

// Pause enables ...
func Pause() func(*Collector) {
	return func(c *Collector) {
		c.Pause = true
	}
}

// DebugMode enables ...
func DebugMode() func(*Collector) {
	return func(c *Collector) {
		c.DebugMode = cfg.DefaultDebugMode
		// c.Config.DebugMode = cfg.DebugMode
	}
}

// VerboseMode enables ...
func VerboseMode() func(*Collector) {
	return func(c *Collector) {
		c.VerboseMode = cfg.DefaultVerboseMode
		// c.Config.VerboseMode = cfg.VerboseMode
	}
}

/*
// ForwardTo enables ...
func ForwardTo(host string) func(*MetricCollector) {
	return func(c *MetricCollector) {
		c.UserAgent = ua
	}
}

// Debugger sets the debugger used by the Collector.
func Debugger(d debug.Debugger) func(*Collector) {
	return func(c *Collector) {
		d.Init()
		c.debugger = d
	}
}
*/
