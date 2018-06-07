package main

import (
	"fmt"

	// core
	colly "github.com/sniperkit/colly/pkg"
	debug "github.com/sniperkit/colly/pkg/debug"

	// plugins
	pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

var version = "0.0.1-alpha"
var debugger *debug.LogDebugger = &debug.LogDebugger{}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: api.github.com
		colly.AllowedDomains("api.github.com"),
	)

	c.AllowTabular = true
	// c.SetDebugger(debugger)

	// On every a element which has href attribute call callback
	c.OnTAB("0:5", func(e *colly.TABElement) {
		fmt.Println("OnTAB event processing...")
		fmt.Printf("\nValid: %t\n", e.Dataset.Valid())
		fmt.Printf("Height: %d\n", e.Dataset.Height())
		fmt.Printf("Width: %d\n", e.Dataset.Width())
		pp.Println("Headers: ", e.Dataset.Headers())

		// ascii := e.Dataset.Tabular("grid")
		// fmt.Println(ascii)

	})

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
	c.Visit("https://api.github.com/users/roscopecoltran/starred?sort=updated&direction=desc&page=1&per_page=10")
}
