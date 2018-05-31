package main

import (
	"net/http"
	"time"

	// cache - core
	// "github.com/gregjones/httpcache"

	"github.com/segmentio/stats"
	"github.com/segmentio/stats/datadog"
	"github.com/segmentio/stats/httpstats"
	"github.com/segmentio/stats/influxdb"
	// "github.com/segmentio/stats/prometheus"
)

// stats/metrics engine
var (
	xStatsEngine *stats.Engine
	xStatsTags   []*stats.Tag
)

// stats storage client(s)
var (
	influxClient     *influxdb.Client
	influxClientConf *influxdb.ClientConfig

	datadogClient     *datadog.Client
	datadogClientConf *datadog.ClientConfig
)

/*
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

/*
	if len(config.Dogstatsd.Address) != 0 {
		stats.Register(datadog.NewClientWith(datadog.ClientConfig{
			Address:    config.Dogstatsd.Address,
			BufferSize: config.Dogstatsd.BufferSize,
		}))
	}
*/

func newStatsTransport(rt http.RoundTripper) http.RoundTripper {
	defer funcTrack(time.Now())
	return httpstats.NewTransport(rt)
}

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

func statsWithTags() {}
