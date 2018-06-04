package main

import (
	"fmt"
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
	version           string         = "0.0.1-alpha"
	configFiles       []string       = []string{"colly.yaml"}
	cacheCollectorDir string         = "./shared/cache/collector"
	defaultDomain     string         = "https://www.shopify.com"
	sitemapURL        string         = "https://www.shopify.com/sitemap.xml"
	kill              bool           = false
	enableDebug       bool           = true
	enableVerbose     bool           = true
	isAutoLoad        bool           = false
	log               *logrus.Logger = logrus.New()
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
	fmt.Println("Handling Keyboard interrupt.")
	kill = true
}

func main() {
	go prepareSignalHandler()

	/*
		trigger.On("first-event", func() {
			// Do Some Task Here.
			fmt.Println("Done")
		})
		trigger.Fire("first-event")

		// trigger.On("second-event", tui.TestFunc)
		// trigger.Fire("second-event")

		// Call them using
		trigger.On("third-event", tui.TestFunc)
		values, err := trigger.Fire("third-event", 5, 6)
		if err != nil {
			log.Println("could not trigger third-event.")
		} else {
			log.Printf("trigger values=%s \n", values)
		}
	*/

	masterCollector, err := newCollectorWithConfig(configFiles...)
	if err != nil {
		log.Println("could not instanciate the master collector.")
	}

	masterCollector = addCollectorEvents(masterCollector)
	// masterCollector = addCollectorDebugEvents(masterCollector)

	// Initialize data collections for storing data/pattern extracted
	// or the sitemap urls by the collector into datasets
	initDataCollections()

	if enableUI {

		initDashboard()
		updateDashboard()
		/*
			stopTheUI = make(chan bool)
			collectorResponseMetrics = make(chan metric.Response)
			go func() {
				tui.Dashboard(stopTheUI, stopTheCrawler)
			}()

			go func() {
				for {
					select {
					case <-allURLsHaveBeenVisited:
						allStatisticsHaveBeenUpdated <- true
						return

					case <-stopTheCrawler:
						log.Println("stopTheUI")
						stopTheUI <- true

					case snapshot := <-collectorResponseMetrics:
						log.Println("new collectorResponseMetrics")
						if collectorStats != nil {
							collectorStats = metric.NewStatsCollector()
						}
						collectorStats.UpdateStatistics(snapshot)

					}
				}
			}()
		*/
	}

	switch collectorMode {
	case "async":

		// Attach master collector to the sitemap collector
		sitemapCollector, err := sitemap.AttachCollector(sitemapURL, masterCollector)
		if err != nil {
			log.Fatalln("could not instanciate the sitemap collector.")
		}
		sitemapCollector.VisitAll()

		sitemapCollector.Count()

		masterCollector.Visit(defaultDomain)

		// Consume URLs
		masterCollector.Wait()

	case "queue":

		// Initialize collector queue
		collectorQueue, err := initCollectorQueue(collectorQueueWorkers, collectorQueueMaxSize, "InMemory")
		if err != nil {
			log.Fatalln("error: ", err)
		}

		// Attach queue and master collector to the sitemap collector
		sitemapCollector, err := sitemap.AttachQueue(sitemapURL, collectorQueue)
		if err != nil {
			log.Fatalln("could not instanciate the sitemap collector.")
		}
		sitemapCollector.Count()

		// Enqueue all URLs found in the sitemap.txt
		sitemapCollector.EnqueueAll()

		// Consume URLs
		collectorQueue.Run(masterCollector)

	default:

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

	// if enableUI && !masterCollector.IsDebug() {
	if enableUI {
		stopTheUI <- true
	}

	if isTachymeter {
		err := closeTachymeter("./shared/exports/stats/tachymeter")
		if err != nil {
			log.Fatalln("error: ", err)
		}
	}

}
