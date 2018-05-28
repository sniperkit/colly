package main

import (
	"fmt"

	"github.com/sniperkit/colly/pkg"
)

func main() {

	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}

	// Create a Collector specifically for Shopify
	c := colly.NewCollector(
		colly.AllowedDomains("golanglibs.com"),
	)

	/*
		// Create a callback on the XPath query searching for the URLs
		c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
			knownUrls = append(knownUrls, e.Text)
		})
	*/

	// Create a callback on the XPath query searching for the URLs
	c.OnTXT("//urlset/url/loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	// Start the collector
	c.Visit("https://golanglibs.com/sitemap.txt")

	fmt.Println("All known URLs:")
	for _, url := range knownUrls {
		fmt.Println("\t", url)
	}

	fmt.Println("Collected", len(knownUrls), "URLs")

}
