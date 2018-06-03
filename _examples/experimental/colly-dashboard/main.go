package main

import (
	"log"

	// colly core
	"github.com/sniperkit/colly/pkg"
	"github.com/sniperkit/colly/pkg/debug"
	"github.com/sniperkit/colly/pkg/helper"
	"github.com/sniperkit/colly/pkg/queue"

	// colly plugins

	//// console UI
	cui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/gocui"
	tui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui"
	// dash "github.com/sniperkit/colly/plugins/cmd/dashboard"
	// cui "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/gocui"
	// tvi "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/tview"

	//// Stats
	ta "github.com/sniperkit/colly/plugins/data/collection/stats/tachymeter"
	// hist "github.com/sniperkit/colly/plugins/cmd/dashboard/tui/termui/histogram"

	//// Sitemap
	sitemap "github.com/sniperkit/colly/plugins/data/format/sitemap"

	// datastructure helpers
	cmmap "github.com/sniperkit/colly/plugins/data/structure/map/multi"
	tablib "github.com/sniperkit/colly/plugins/data/transform/tabular"
)

const SITEMAP_URL string = "https://www.shopify.com/sitemap.xml"

var (
	xConsoleUI *cui.TermUI
	xTermUI    *cui.TermUI
	xResults   chan tui.WorkResult
	stopTheUI  chan bool
)

var (
	// tachymeter
	startedAt            time.Time
	isTachymeter         bool = true
	isTachymeterParallel bool = false
	cTachymeter          chan *ta.Tachymeter
	xTachymeter          *ta.Tachymeter
	xTachyResults        *ta.Metrics
	xTachymeterTL        ta.Timeline
	wallTimeStart        time.Time
)

var (
	version           string                   = "0.0.1-alpha"
	cacheCollectorDir string                   = "./shared/cache/collector"
	sheets            map[string][]interface{} = make(map[string][]interface{}, 0)
	dsExport          *tablib.Dataset
	dsURLs            *tablib.Dataset
	dataBook          *tablib.Databook
	mapKnownURLs      = cmmap.NewConcurrentMultiMap()
	logger            *log.Logger
	cq                *queue.Queue
)

func init() {
	initCollections()
}

func initCollections() {
	dsURLs = tablib.NewDataset([]string{"loc", "changefreq", "priority"})
	dsExport = tablib.NewDataset([]string{"url", "price", "size", "colors"})
	dataBook = tablib.NewDatabook()
	dataBook.AddSheet("patterns", dsExport) //.Sort("price"))         // add the patterns sheets to the collector databook
	dataBook.AddSheet("known_urls", dsURLs) //.Sort("priority")) // add the kown_urls sheets to the collector databook
}

func main() {

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

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		// Delay:      5 * time.Second,
	})

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
