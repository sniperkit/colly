package main

import (
	// default packages
	"log"

	//// collector - core packages

	//// collector - queue storage
	res "github.com/sniperkit/colly/plugins/data/storage/backends/redis"
	// lru "github.com/sniperkit/colly/plugins/data/storage/backends/lru"
	// sq3 "github.com/sniperkit/colly/plugins/data/storage/backends/sqlite3"
	// baq "github.com/sniperkit/colly/plugins/data/storage/backends/badger"
	// stq "github.com/sniperkit/colly/plugins/data/storage/backends/storm"
	// myq "github.com/sniperkit/colly/plugins/data/storage/backends/mysql"
	// moq "github.com/sniperkit/colly/plugins/data/storage/backends/mongo"
	// elq "github.com/sniperkit/colly/plugins/data/storage/backends/elastic"
	// shq "github.com/sniperkit/colly/plugins/data/storage/backends/sphinx"
	// caq "github.com/sniperkit/colly/plugins/data/storage/backends/cassandra"

	//// collector - console UI
	cui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/gocui"
	tui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui"
	// dash "github.com/sniperkit/colly/plugins/cmd/dashboard"
	// cui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/gocui"
	// tvi "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/tview"

	//// collector - stats
	ta "github.com/sniperkit/colly/plugins/data/collection/stats/tachymeter"
	// hist "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui/histogram"
)

// app params
var (
	version           string = "0.0.1-alpha"
	configFile        string = "colly.yaml"
	cacheCollectorDir string = "./shared/cache/collector"
	sitemapURL        string = "https://www.shopify.com/sitemap.xml"
	logger            *log.Logger
)

// Initialize default object instances for the application.
// Components initilization list:
// - Data Collection mananger; create default datasets and hook them to the default databook
// - Collector Queue mananger; create
func init() {

	// Initialize data collections for storing data/pattern extracted
	// or the sitemap urls by the collector into datasets
	initDataCollections()

	// Initialize collector queue
	xQueue, err = initCollectorQueue(collectorWorkers, queueMaxSize, "InMemory")
	if err != nil {
		log.Fatalln("error: ", err)
	}

}

func main() {

	helper.RandomUserAgent(c)
	helper.Referrer(c)

	cs, err := sitemap.NewWithCollector(SITEMAP_URL_TXT, c)
	if err != nil {
		log.Println("invalid sitemap.")
	}

	// _, sitemaps := cs.List()
	// log.Println("sitemaps: ", strings.Join(sitemaps, ","))

	if cq != nil {
		cs.EnqueueAll()
	} else {
		cs.VisitAll()
	}
	cs.Count()

	// c.Visit(SITEMAP_INDEX_URL)

	c.Wait()

	// c.Wait()
	// Start the collector
	// c.Visit("https://www.shopify.com/sitemap.xml")

}

func initDataCollections() {
	dsURLs = tablib.NewDataset([]string{"loc", "changefreq", "priority"})
	dsExport = tablib.NewDataset([]string{"url", "price", "size", "colors"})
	dataBook = tablib.NewDatabook()
	dataBook.AddSheet("patterns", dsExport) //.Sort("price"))         // add the patterns sheets to the collector databook
	dataBook.AddSheet("known_urls", dsURLs) //.Sort("priority")) // add the kown_urls sheets to the collector databook
}

func initCollectorQueue(ct int, s int, b string) (q *queue.Queue, err error) {
	if ct < 0 {
		err = errInvalidQueueThreads
		return
	}
	b = strings.ToLower(b)
	if s < 0 {
		err = errInvalidQueueMaxSize
		return
	}

	switch cqStore {
	case "sqlite":
		fallthrough

	case "sqlite3": // Warning! Conflict with Pivot
		q, err = queue.New(ct, &sq3.Storage{Filename: "./shared/datastore/queue.db"})

	case "redis":
		q, err = queue.New(ct, &res.Storage{Address: "127.0.0.1:6379", Password: "", DB: 0, Prefix: "job01"})

	case "badger":
		log.Printf("Collector queue storage/backend '%s' is not implemented yet...\n", cqStore)

	case "diskv":
		log.Printf("Collector queue storage/backend '%s' is not implemented yet...\n", cqStore)

	case "boltdb":
		log.Printf("Collector queue storage/backend '%s' is not implemented yet...\n", cqStore)

	case "lru":
		log.Printf("Collector queue storage/backend '%s' is not implemented yet...\n", cqStore)

	case "inmemory":
		fallthrough

	default:
		q, err = queue.New(ct, &queue.InMemoryQueueStorage{MaxSize: s})

	}
	return
}
