package main

import (
	"net/http"
	"time"

	"github.com/sniperkit/xstats/pkg"

	"github.com/sniperkit/xstats/client/datadog"
	"github.com/sniperkit/xstats/client/influxdb"
	// "github.com/sniperkit/xstats/client/prometheus"
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
	return stats.NewTransport(rt)
}

func newStatsEngine(backend string) {
	switch backend {
	case "datadog":
		statsEngine = nil
	case "influxdb":
		fallthrough
	default:
		statsEngine = nil
	}
}

func statsWithTags() {}
