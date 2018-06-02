package main

import (
	"log"

	// colly core
	"github.com/sniperkit/colly/pkg"
	"github.com/sniperkit/colly/pkg/debug"
	"github.com/sniperkit/colly/pkg/helper"
	"github.com/sniperkit/colly/pkg/queue"

	// colly plugins
	sitemap "github.com/sniperkit/colly/plugins/data/format/sitemap"

	// datastructure helpers
	cmmap "github.com/sniperkit/colly/plugins/data/structure/map/multi"
	tablib "github.com/sniperkit/colly/plugins/data/transform/tabular"
)

var (
	version                                    = "0.0.1-alpha"
	cacheCollectorDir                          = "./shared/cache/collector"
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
		//Delay:      5 * time.Second,
	})

	helper.RandomUserAgent(c)
	helper.Referrer(c)

	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		log.Println("url=", e.Text)
	})

	cs, err := sitemap.NewWithCollector("https://www.shopify.com/sitemap.xml", c)
	if err != nil {
		log.Println("invalid sitemap.")
	}
	if cq != nil {
		cs.EnqueueAll()
	} else {
		cs.VisitAll()
	}

	c.Wait()

	// c.Wait()

	// Start the collector
	// c.Visit("https://www.shopify.com/sitemap.xml")

}
