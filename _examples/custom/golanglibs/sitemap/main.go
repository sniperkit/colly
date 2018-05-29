package main

import (
	"log"

	"github.com/sniperkit/colly/pkg"
	cfg "github.com/sniperkit/colly/pkg/config"
	// "github.com/sniperkit/go-tablib"
	// "github.com/sniperkit/xutil/plugin/debug/pp"
)

var (
	// logger  *logger.Logger
	scraper     *colly.Collector
	collyConfig *cfg.CollectorConfig
	urls        []string = []string{} // Array containing all the known URLs in a sitemap
)

func main() {

	if collyConfig != nil {
		collyConfig = &cfg.CollectorConfig{}
		collyConfig.DebugMode = true
		collyConfig.VerboseMode = true
		scraper = colly.NewCollectorWithConfig(collyConfig)

	} else {
		// Create a Collector specifically for Shopify
		scraper = colly.NewCollector(
			colly.AllowedDomains(defaultAllowedDomains),
		)

	}

	// Create a callback on the Column name to get all URLs to scrape
	scraper.OnTAB(defaultSitenapXML_XPath, func(e *colly.TABElement) {
		urls = append(urls, e.Text)
		log.Println("content: ", e.Text)
	})

	// Start the collector
	scraper.Visit(defaultSitemapURL)

	log.Println("All known URLs:")
	for _, url := range urls {
		log.Println("\t", url)
	}

	log.Println("Collected", len(urls), "URLs")

}
