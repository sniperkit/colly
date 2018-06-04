package main

import (
	"runtime"
	"time"

	colly "github.com/sniperkit/colly/pkg"
	config "github.com/sniperkit/colly/pkg/config"
	debug "github.com/sniperkit/colly/pkg/debug"
	helper "github.com/sniperkit/colly/pkg/helper"
	metric "github.com/sniperkit/colly/pkg/metric"

	// helpers
	// "github.com/davecgh/go-spew/spew"
	// "github.com/k0kubun/pp"
	pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

var appConfig *config.CollectorConfig

// collector
var (
	// collectorCacheDir   string             = "./shared/cache/collector"
	collectorStop       chan bool          = make(chan bool)
	collectorAllVisited chan bool          = make(chan bool)
	collectorResult     chan error         = make(chan error)
	collectorDebug      *debug.LogDebugger = &debug.LogDebugger{}
)

func initCollectorHelpers(c *colly.Collector) *colly.Collector {
	helper.RandomUserAgent(c)
	helper.Referrer(c)
	return c
}

func newCollectorWithConfig(files ...string) (*colly.Collector, error) {

	// Enable debug mode or set env `CONFIGOR_DEBUG_MODE` to true when running your application
	var err error
	appConfig, err = config.NewFromFile(false, false, files...)
	if err != nil {
		return nil, err
	}

	// Dump config file for dev purpise
	dumpFormats := []string{"yaml", "json", "toml", "xml"} // "ini"
	dumpNodes := []string{}
	config.Dump(appConfig, dumpFormats, dumpNodes, "./shared/exports/config/dump/colly_dashboard") // use string slices

	// pp.WithLineInfo = false
	// spew.Sdump(appConfig)
	pp.Println("Config=", appConfig)

	// Create a Collector
	collector := colly.NewCollector()

	// Set User-Agent
	if appConfig.UserAgent != "" {
		collector.UserAgent = `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36`
	}

	// Set Allowed Domains
	if len(appConfig.Filters.Whitelists.Domains) > 0 {
		collector.AllowedDomains = appConfig.Filters.Whitelists.Domains
	}

	if len(appConfig.Filters.Blacklists.Domains) > 0 {
		collector.DisallowedDomains = appConfig.Filters.Whitelists.Domains
	}

	// Set Verbose Mode components
	/*
		if appConfig.App.VerboseMode {
			collector = addCollectorVerboseEvents(collector)
		}
	*/

	// Set Debugger
	if appConfig.App.DebugMode {
		collector.SetDebugger(&debug.LogDebugger{})
	}

	// Set Custom HttpCache Transport
	// if appConfig.Collector.Transport.Http.Cache.Enabled {
	// }

	// Set Custom HttpStats Transport
	// if appConfig.Collector.Transport.Http.Stats.Enabled {
	// }

	collector.AllowURLRevisit = appConfig.Collector.AllowURLRevisit
	collector.IgnoreRobotsTxt = appConfig.Collector.IgnoreRobotsTxt
	collector.CacheDir = appConfig.Collector.Cache.Directory

	if appConfig.Collector.CurrentMode == "async" {
		collector.Async = true
		collectorLimits := &colly.LimitRule{}
		collectorLimits.DomainGlob = appConfig.Collector.Modes.Async.DomainGlob
		if appConfig.Collector.Modes.Async.Parallelism > 0 {
			collectorLimits.Parallelism = appConfig.Collector.Modes.Async.Parallelism
		} else {
			collectorLimits.Parallelism = runtime.NumCPU() - 1
		}
		if appConfig.Collector.Modes.Async.RandomDelay > 0 {
			collectorLimits.RandomDelay = appConfig.Collector.Modes.Async.RandomDelay * time.Second
		}
		collector.Limit(collectorLimits)
	}

	return collector, nil
}

func newCollectorLimits(domain *string, parallelism *int, delay *time.Duration) *colly.LimitRule {
	collectorLimitConfig := &colly.LimitRule{}

	//
	collectorLimitConfig.DomainGlob = "*"
	if domain != nil {
		collectorLimitConfig.DomainGlob = *domain
	}

	// must be superior to 1
	if parallelism != nil {
		if *parallelism < 1 {
			collectorLimitConfig.Parallelism = 2
		}
	}

	// must be superior to 1
	if delay != nil {
		collectorLimitConfig.Delay = 5 * time.Second
		if *delay > 0 {
			collectorLimitConfig.Delay = *delay
		}
	}

	return collectorLimitConfig
}

func addCollectorVerboseEvents(scraper *colly.Collector) *colly.Collector {
	// Set error handler
	scraper.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Before making a request print "Visiting ..."
	scraper.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})
	return scraper
}

func addCollectorEvents(scraper *colly.Collector) *colly.Collector {
	scraper.OnResponse(func(r *colly.Response) {
		if !appConfig.App.DashboardMode {
			log.Infoln("[REQUEST] url=", r.Request.URL.String())
		} else {
			collectorResponseMetrics <- metric.Response{
				URL:             *r.Request.URL,
				NumberOfWorkers: appConfig.Collector.Modes.Queue.WorkersCount,
				ResponseSize:    r.GetSize(),
				StatusCode:      r.GetStatusCode(),
				StartTime:       r.GetStartTime(),
				EndTime:         r.GetEndTime(),
				ContentType:     r.GetContentType(),
			}
		}
	})

	scraper.OnError(func(r *colly.Response, e error) {
		if !appConfig.App.DashboardMode {
			log.Println("[ERROR] msg=", e, ", url=", r.Request.URL, ", body=", string(r.Body))
		} else {
			collectorResponseMetrics <- metric.Response{
				Err:             e,
				URL:             *r.Request.URL,
				NumberOfWorkers: appConfig.Collector.Modes.Queue.WorkersCount,
				ResponseSize:    r.GetSize(),
				StatusCode:      r.GetStatusCode(),
				StartTime:       r.GetStartTime(),
				EndTime:         r.GetEndTime(),
				ContentType:     r.GetContentType(),
			}
		}
	})
	return scraper
}
