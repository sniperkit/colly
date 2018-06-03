package main

import (
	"time"

	colly "github.com/sniperkit/colly/pkg"
	config "github.com/sniperkit/colly/pkg/config"
	debug "github.com/sniperkit/colly/pkg/debug"
	helper "github.com/sniperkit/colly/pkg/helper"
)

// collector -
var (
	cc *config.Config   = &cfg.Config{} // collector's config object
	mc *colly.Collector                 // main/master collector instance
	sc *colly.Collector                 // slave/details collector instance
)

// collector channels
var (
	// collectorDone       chan struct{}
	collectorStop       chan bool  = make(chan bool)
	collectorAllVisited chan bool  = make(chan bool)
	collectorResult     chan error = make(chan error)
)

func newCollector() (*colly.Collector, error) {
	// Create a Collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains("www.shopify.com"),
		colly.Debugger(&debug.LogDebugger{}), // Attach a debugger to the collector
		// colly.IgnoreRobotsTxt(),
		colly.Async(true),
		colly.CacheDir(cacheCollectorDir), // Cache responses to prevent multiple download of pages even if the collector is restarted
		// colly.Debugger(&debug.LogDebugger{}),
		// colly.Async(true),
	)
	return c, nil
}

func newCollectorWithConfig(cfg *config.Config) (*colly.Collector, error) {
	// Create a Collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains("www.shopify.com"),
		colly.Debugger(&debug.LogDebugger{}), // Attach a debugger to the collector
		// colly.IgnoreRobotsTxt(),
		colly.Async(true),
		colly.CacheDir(cacheCollectorDir), // Cache responses to prevent multiple download of pages even if the collector is restarted
		// colly.Debugger(&debug.LogDebugger{}),
		// colly.Async(true),
	)
	return c, nil
}

func newCollectorLimits(domain *string, parallelism *int, delay *time.Duration) *colly.Collector {
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

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(collectorLimitConfig)
}
