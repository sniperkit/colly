package main

import (
	// Logger
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	// collector - advanced sitemap parser
	sitemap "github.com/sniperkit/colly/plugins/data/format/sitemap"
)

// app params
var (
	version     string   = "0.0.1-alpha"
	configFiles []string = []string{
		"./conf/app.yaml",
		"./conf/collector.yaml",
		"./conf/collection.yaml",
		"./conf/filters.yaml",
		"./conf/outputs.yaml",
		"./conf/debug.yaml",
		"./conf/legacy.yaml",
	}
	log *logrus.Logger = logrus.New()
)

// Initialize collector and other components
func init() {
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel
}

func main() {

	masterCollector, err := newCollectorWithConfig(configFiles...)
	if err != nil {
		log.Println("could not instanciate the master collector.")
	}

	masterCollector = addCollectorEvents(masterCollector)

	// Initialize data collections for storing data/pattern extracted
	// or the sitemap urls by the collector into custom datasets
	initDataCollections()

	if appConfig.App.DashboardMode {
		appConfig.App.DebugMode = false
		appConfig.App.VerboseMode = false
		initDashboard()
		updateDashboard()
	}

	if !appConfig.Debug.Tachymeter.Disabled {
		initTachymeter()
	}

	switch appConfig.Collector.CurrentMode {
	case "async":

		if !appConfig.Collector.Sitemap.Disabled && appConfig.Collector.Sitemap.URL != "" {
			// Attach master collector to the sitemap collector
			sitemapCollector, err := sitemap.AttachCollector(appConfig.Collector.Sitemap.URL, masterCollector)
			if err != nil {
				log.Fatalln("could not instanciate the sitemap collector.")
			}
			sitemapCollector.VisitAll()
			sitemapCollector.Count()
		}
		masterCollector.Visit(appConfig.Collector.RootURL)

		// Consume URLs
		masterCollector.Wait()

	case "queue":
		// Initialize collector queue
		collectorQueue, err := initCollectorQueue(appConfig.Collector.Modes.Queue.WorkersCount, appConfig.Collector.Modes.Queue.MaxSize, "InMemory")
		if err != nil {
			log.Fatalln("error: ", err)
		}

		if !appConfig.Collector.Sitemap.Disabled && appConfig.Collector.Sitemap.URL != "" {
			// Attach queue and master collector to the sitemap collector
			sitemapCollector, err := sitemap.AttachQueue(appConfig.Collector.Sitemap.URL, collectorQueue)
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
		if !appConfig.Collector.Sitemap.Disabled && appConfig.Collector.Sitemap.URL != "" {
			// Initalize new sitemap collector
			sitemapCollector, err := sitemap.New(appConfig.Collector.Sitemap.URL)
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

	if !appConfig.Debug.Tachymeter.Disabled {
		err := closeTachymeter("./shared/exports/stats/tachymeter")
		if err != nil {
			log.Fatalln("error: ", err)
		}
	}

}
