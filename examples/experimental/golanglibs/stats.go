package main

import (
	"net/http"
	"time"

	// stats - core
	"github.com/segmentio/stats"

	// stats - collectors
	"github.com/segmentio/stats/httpstats"

	// stats - remote clients
	"github.com/segmentio/stats/datadog"
	"github.com/segmentio/stats/influxdb"
	// "github.com/segmentio/stats/prometheus"
)

/*
	Notes:

	*** InfluxDB (API) ***
	- Install: `brew install influxdb`
	- Run:
		- To have launchd start influxdb now and restart at login: `brew services start influxdb`
		- Or, if you don't want/need a background service you can just run: `influxd -config /usr/local/etc/influxdb.conf`

	*** Chronograf (UI) ***
	- Install: `brew install chronograf`
	- Run:
		- To have launchd start chronograf now and restart at login: `brew services start chronograf`
		- Or, if you don't want/need a background service you can just run: `chronograf`
*/

var (
	isStatsTransport             bool = true
	xStatsEngine                 *stats.Engine
	xStatsTags                   []*stats.Tag
	allStatisticsHaveBeenUpdated chan bool
)

var (
	influxClient     *influxdb.Client
	influxClientConf *influxdb.ClientConfig

	datadogClient     *datadog.Client
	datadogClientConf *datadog.ClientConfig
)

func newStatsTransport(rt http.RoundTripper) http.RoundTripper {
	defer funcTrack(time.Now())
	return httpstats.NewTransport(rt)
}

func newStatsEngine(backend string) {
	switch backend {
	case "datadog":
		xStatsEngine = nil
	case "influxdb":
		fallthrough
	default:
		xStatsEngine = nil
	}
}

func addMetrics(start time.Time, incr int, failed bool) {
	callTime := time.Now().Sub(start)
	m := &funcMetrics{}
	m.calls.count = incr
	m.calls.time = callTime
	if failed {
		m.calls.failed = incr
	}
	stats.Report(m)
}

func statsWithTags() {}

/*
	if len(config.Dogstatsd.Address) != 0 {
		stats.Register(datadog.NewClientWith(datadog.ClientConfig{
			Address:    config.Dogstatsd.Address,
			BufferSize: config.Dogstatsd.BufferSize,
		}))
	}
*/

/*
func newCacheTransportWithStats(engine string, prefixPath string) (httpcache.Cache, *httpcache.Transport) {
	defer funcTrack(time.Now())

	backendCache, err := newCacheBackend(engine, prefixPath)
	if err != nil {
		log.Fatal("cache err", err.Error())
	}

	var httpTransport = http.DefaultTransport
	httpTransport = httpstats.NewTransport(httpTransport)
	http.DefaultTransport = httpTransport

	cachingTransport := httpcache.NewTransportFrom(backendCache, httpTransport) // httpcache.NewMemoryCacheTransport()
	cachingTransport.MarkCachedResponses = true

	return backendCache, cachingTransport
}
*/
