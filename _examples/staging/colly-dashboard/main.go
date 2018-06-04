package main

import (
	// "fmt"
	"os"
	"os/signal"

	// metric "github.com/sniperkit/colly/pkg/metric"
	// tui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui"

	// "github.com/sadlil/go-trigger"
	// "github.com/olebedev/emitter"
	// "github.com/chuckpreslar/emission"
	// "github.com/GianlucaGuarini/go-observable"

	// Logger
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	//// collector - advanced sitemap parser
	sitemap "github.com/sniperkit/colly/plugins/data/format/sitemap"
)

// app params
var (
	version       string         = "0.0.1-alpha"
	configFiles   []string       = []string{"colly.yaml"}
	defaultDomain string         = "https://www.shopify.com"
	sitemapURL    string         = "https://www.shopify.com/sitemap.xml"
	kill          bool           = false
	enableDebug   bool           = false
	enableVerbose bool           = false
	isAutoLoad    bool           = false
	log           *logrus.Logger = logrus.New()
)

// Initialize default object instances for the application.
// Components initilization list:
// - Data Collection mananger; create default datasets and hook them to the default databook
// - Collector Queue mananger; create
func init() {
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel
}

func prepareSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	stopTheCrawler <- true
	log.Println("Handling Keyboard interrupt.")
	kill = true
}

func main() {
	// go prepareSignalHandler()

	masterCollector, err := newCollectorWithConfig(configFiles...)
	if err != nil {
		log.Println("could not instanciate the master collector.")
	}

	masterCollector = addCollectorEvents(masterCollector)

	// Initialize data collections for storing data/pattern extracted
	// or the sitemap urls by the collector into custom datasets
	initDataCollections()

	if appConfig.App.DashboardMode {

		initDashboard()
		updateDashboard()
	}

	switch collectorMode {
	case "async":
		if !appConfig.Crawler.Sitemap.Disabled {
			// Attach master collector to the sitemap collector
			sitemapCollector, err := sitemap.AttachCollector(sitemapURL, masterCollector)
			if err != nil {
				log.Fatalln("could not instanciate the sitemap collector.")
			}
			sitemapCollector.VisitAll()
			sitemapCollector.Count()
		}
		masterCollector.Visit(defaultDomain)
		// Consume URLs
		masterCollector.Wait()

	case "queue":
		// Initialize collector queue
		collectorQueue, err := initCollectorQueue(collectorQueueWorkers, collectorQueueMaxSize, "InMemory")
		if err != nil {
			log.Fatalln("error: ", err)
		}

		if !appConfig.Crawler.Sitemap.Disabled {
			// Attach queue and master collector to the sitemap collector
			sitemapCollector, err := sitemap.AttachQueue(sitemapURL, collectorQueue)
			if err != nil {
				log.Fatalln("could not instanciate the sitemap collector.")
			}
			sitemapCollector.Count()
			// Enqueue all URLs found in the sitemap.txt
			sitemapCollector.EnqueueAll()
		}

		// Consume URLs
		collectorQueue.Run(masterCollector)

	default:
		if !appConfig.Crawler.Sitemap.Disabled {
			// Initalize new sitemap collector
			sitemapCollector, err := sitemap.New(sitemapURL)
			if err != nil {
				log.Fatalln("could not instanciate the sitemap collector.")
			}
			sitemapCollector.Count()
			urls, _ := sitemapCollector.List()
			for _, url := range urls {
				masterCollector.Visit(url.String())
			}
		}

	}

	// if enableUI && !masterCollector.IsDebug() {
	if appConfig.App.DashboardMode {
		stopTheUI <- true
	}

	if isTachymeter {
		err := closeTachymeter("./shared/exports/stats/tachymeter")
		if err != nil {
			log.Fatalln("error: ", err)
		}
	}

}
