package main

import (
	"fmt"

	colly "github.com/sniperkit/colly/pkg"
	// debug "github.com/sniperkit/colly/pkg/debug"
)

var version = "0.0.1-alpha"

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: api.github.com
		colly.AllowedDomains("api.github.com"),
	)

	// c.SetDebugger(&debug.LogDebugger{})

	// On every a element which has href attribute call callback
	c.OnJSON("//description", func(e *colly.JSONElement) {
		fmt.Printf("Values found: %s\n", e.Text)
		// link := e.Node("description")
		// Print link
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://api.github.com/users/roscopecoltran/starred?sort=updated&direction=desc&page=1")
}
