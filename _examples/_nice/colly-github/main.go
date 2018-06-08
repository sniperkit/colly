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
var (
	// appVersion specifies the version of the app. If the executable is built with Makefile, the appVersion will use the actual repo's short tag version
	appVersion = "0.0.1-alpha"
	// appDebug specifies if the app debug/verbose some development event logged
	appDebug = true
)

// github api vars
var (
	// githubAPI_Account sets the github user name to request for its starred repositories list
	githubAPI_Account = "roscopecoltran"
	// githubAPI_Pagination_Page
	githubAPI_Pagination_Page = 1
	// githubAPI_Pagination_PerPage
	githubAPI_Pagination_PerPage = 10
	// githubAPI_Pagination_Direction
	githubAPI_Pagination_Direction = "desc"
	// githubAPI_Pagination_Sort
	githubAPI_Pagination_Sort = "updated"
	// githubAPI_Account
	githubAPI_Pagination_Params = []string{
		fmt.Sprintf("page=%d", githubAPI_Pagination_Page),
		fmt.Sprintf("per_page=%s", githubAPI_Pagination_PerPage),
		fmt.Sprintf("direction=%s", githubAPI_Pagination_Direction),
		fmt.Sprintf("sort=%s", githubAPI_Pagination_Sort),
	}
	// githubAPI_EndpointURL
	githubAPI_EndpointURL = fmt.Println("https://api.github.com/users/%s/starred?%s", githubAPI_Account, strings.Join(githubAPI_Pagination_Params, "&"))
)

// collector vars
var (
	// collectorDebugger stores the collector's log event listener
	collectorDebugger *debug.LogDebugger = &debug.LogDebugger{}
	// collectorDatasetDebug sets some debugging information
	collectorDatasetDebug = true
	// collectorDatasetDebug sets some debugging information
	collectorDatasetEnable = true
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
	//  collectorSubDatasetColumns specifies the columns to filter from the json content
	collectorSubDatasetColumns = []string{"id", "name", "full_name", "description", "language", "stargazers_count", "forks_count"}
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
		colly.AllowTabular(true),
	)

	c.AllowTabular = true
	if collectorWithDebug {
		c.SetDebugger(debugger)
	}

	// On every a element which has tabular format data call callback
	// Notes:
	// - `OnTAB` callback event are enabled only if the `AllowTabular` attribute is set to true.
	// - `OnTAB` use a fork of the package `github.com/agrison/go-tablib`
	// - `OnTAB` query specifications are available in 'SPECS.md'
	c.OnTAB("0:0", func(e *colly.TABElement) {

		// Debug the dataset slice
		if appDebug {
			fmt.Printf("\nValid: %t\n", e.Dataset.Valid(), "Height: %d\n", e.Dataset.Height(), "Width: %d\n", e.Dataset.Width())
			pp.Println(e.Dataset.Headers())
		}

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

	// On every a element which has json content-type or file extension call callback
	// Notes:
	// - If `AllowTabular` is true, OnJSON is overrided by the OnTAB callback event
	// - OnJSON use a fork of the package `github.com/antchfx/jsonquery`
	c.OnJSON("//description", func(e *colly.JSONElement) {
		if appDebug {
			fmt.Printf("Values found: %s\n", e.Text)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		if appDebug {
			fmt.Println("Visiting", r.URL.String())
		}
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(githubAPI_EndpointURL)
}
