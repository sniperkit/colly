package main

import (
	"fmt"
	"strings"

	// colly - core
	colly "github.com/sniperkit/colly/pkg"
	debug "github.com/sniperkit/colly/pkg/debug"

	// colly - plugins
	pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

// app vars
var appVersion = "0.0.1-alpha"

// collector vars
var (
	// collectorWithDebug specifies if the collector debugger is enabled or disabled
	collectorWithDebug = true
	// collectorDebugger stores the collector's log event listener
	collectorDebugger *debug.LogDebugger = &debug.LogDebugger{}
	// collectorDatasetDebug sets some debugging information
	collectorDatasetDebug = true
	// collectorDatasetOutputPrefixPath specifies the prefix path for all saved dumps
	collectorDatasetOutputPrefixPath = "./shared/dataset"
	// collectorDatasetOutputBasename specifies the template to use to write the dataset dump
	collectorDatasetOutputBasename = "colly_github_%d"
	// collectorDatasetOutputFormat sets the ouput format of the dataset extracted by the collector
	// `OnTAB` event Export formats supported:
	//  - JSON (Sets + Books)
	//  - YAML (Sets + Books)
	//  - XLSX (Sets + Books)
	//  - XML (Sets + Books)
	//  - TSV (Sets)
	//  - CSV (Sets)
	//  - ASCII + Markdown (Sets)
	//  - MySQL (Sets)
	//  - Postgres (Sets)
	// `OnTAB` event loading formats supported:
	//  - JSON (Sets + Books)
	//  - YAML (Sets + Books)
	//  - XML (Sets)
	//  - CSV (Sets)
	//  - TSV (Sets)
	collectorDatasetOutputFormat = "json"
)

// AppendDynamicColumn to the tabular dataset
func descriptionLen(row []interface{}) interface{} {
	if row == nil {
		return 0
	}
	return len(row[3].(string))
}

// PrettyPrint structs
func prettyPrint(msg ...interface{}) {
	pp.Println(interface{}...)
}

func init() {
	// Ensure that the output format is set in lower case.
	collectorDatasetOutputFormat = strings.ToLower(collectorDatasetOutputFormat)
}

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: api.github.com
		colly.AllowedDomains("api.github.com"),
	)

	c.AllowTabular = true
	if collectorWithDebug {
		c.SetDebugger(debugger)
	}

	// On every a element which has href attribute call callback
	c.OnTAB("0:0", func(e *colly.TABElement) {

		if collectorDatasetDebug {

		}
		// fmt.Println("OnTAB event processing...")
		// fmt.Printf("\nValid: %t\n", e.Dataset.Valid(), "Height: %d\n", e.Dataset.Height(), "Width: %d\n", e.Dataset.Width())
		// pp.Println(e.Dataset.Headers())

		// Select
		ds, err := e.Dataset.Select(0, 0, "id", "name", "full_name", "description", "language", "stargazers_count", "forks_count")
		if err != nil {
			fmt.Println("error:", err)
		}
		ds.AppendDynamicColumn("description_length", descriptionLen)

		switch outputFormat {
		case "yaml":
		case "json":
		case "tsv":
		case "csv":
		case "grid-simple":
		case "grid-condensed":
		case "grid-markdown", "ascii":
		case "mysql":
		case "postgresql":
		}
		// JSON
		json, err := ds.JSON()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(json)

		// YAML
		yaml, err := ds.YAML()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(yaml)

		// fmt.Println(ds.SortReverse("id").Tabular("condensed"))

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
