package main

import (
	"runtime"
	"time"

	colly "github.com/sniperkit/colly/pkg"
	config "github.com/sniperkit/colly/pkg/config"
	debug "github.com/sniperkit/colly/pkg/debug"
	helper "github.com/sniperkit/colly/pkg/helper"
	metric "github.com/sniperkit/colly/pkg/metric"
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

func newCollector(domain string, async bool) (*colly.Collector, error) {

	// Create a Collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains(domain),
		colly.Async(async),
		colly.CacheDir("./shared/cache/collector"), // Cache responses to prevent multiple download of pages even if the collector is restarted
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second,
	})

	return c, nil
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

	// Create a Collector
	collector := colly.NewCollector()

	if appConfig.UserAgent != "" {
		collector.UserAgent = `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36`
	}

	if len(appConfig.AllowedDomains) > 0 {
		collector.AllowedDomains = appConfig.AllowedDomains
	}

	if appConfig.App.VerboseMode {
		collector = addCollectorVerboseEvents(collector)
	}

	if appConfig.App.DebugMode {
		collector.SetDebugger(&debug.LogDebugger{})
	}

	if appConfig.AllowURLRevisit {
		collector.AllowURLRevisit = true
	}

	if appConfig.IgnoreRobotsTxt {
		collector.IgnoreRobotsTxt = true
	}

	if appConfig.Collector.Cache.Directory != "" {
		collector.CacheDir = appConfig.Collector.Cache.Directory
	}

	if appConfig.Collector.CurrentMode == "async" {

		collector.Async = true

		collectorLimits := &colly.LimitRule{}
		if appConfig.Collector.Modes.Async.DomainGlob != "" {
			collectorLimits.DomainGlob = appConfig.Collector.Modes.Async.DomainGlob
		}
		if appConfig.Collector.Modes.Async.Parallelism > 0 {
			collectorLimits.Parallelism = appConfig.Collector.Modes.Async.Parallelism
		} else {
			collectorLimits.Parallelism = runtime.NumCPU()
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
