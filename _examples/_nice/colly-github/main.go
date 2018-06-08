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
	// githubAPIAccount sets the github user name to request for its starred repositories list
	githubAPIAccount = "roscopecoltran"
	// githubAPIPaginationPage
	githubAPIPaginationPage = 1
	// githubAPIPaginationPerPage
	githubAPIPaginationPerPage = 10
	// githubAPIPaginationDirection
	githubAPIPaginationDirection = "desc"
	// githubAPIPaginationSort
	githubAPIPaginationSort = "updated"
	// githubAPIPaginationParams
	githubAPIPaginationParams = []string{
		fmt.Sprintf("page=%d", githubAPIPaginationPage),
		fmt.Sprintf("per_page=%d", githubAPIPaginationPerPage),
		fmt.Sprintf("direction=%s", githubAPIPaginationDirection),
		fmt.Sprintf("sort=%s", githubAPIPaginationSort),
	}
	// githubAPI_EndpointURL
	githubAPI_EndpointURL = fmt.Sprintf("https://api.github.com/users/%s/starred?%s", githubAPIAccount, strings.Join(githubAPIPaginationParams, "&"))
)

// collector vars
var (
	// collectorDebugger stores the collector's log event listener
	collectorDebugger *debug.LogDebugger = &debug.LogDebugger{}
	// collectorDatasetDebug sets some debugging information
	collectorDatasetDebug = true
	// collectorTabEnabled sets some debugging information
	collectorTabEnabled = true
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
	pp.Println(msg...)
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

	if appDebug {
		c.SetDebugger(collectorDebugger)
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

		// ds.EXPORT_FORMAT().String() 					--> returns the contents of the exported dataset as a string.
		// ds.EXPORT_FORMAT().Bytes() 					--> returns the contents of the exported dataset as a byte array.
		// ds.EXPORT_FORMAT().WriteTo(writer) 			--> writes the exported dataset to w.
		// ds.EXPORT_FORMAT().WriteFile(filename, perm) --> writes the databook or dataset content to a file named by filename.
		var output string
		switch collectorDatasetOutputFormat {
		// YAML
		case "yaml":
			if export, err := ds.YAML(); err == nil {
				output = export.String()
			} else {
				fmt.Println("error:", err)
			}

		// JSON
		case "json":
			if export, err := ds.JSON(); err == nil {
				output = export.String()
			} else {
				fmt.Println("error:", err)
			}

		// TSV
		case "tsv":
			if export, err := ds.TSV(); err == nil {
				output = export.String()
			} else {
				fmt.Println("error:", err)
			}

		// CSV
		case "csv":
			if export, err := ds.CSV(); err == nil {
				output = export.String()
			} else {
				fmt.Println("error:", err)
			}

		// Markdown
		case "markdown":
			output = ds.Markdown().String()

		// HTML
		case "html":
			output = ds.HTML().String()

		// MySQL
		case "mysql":
			output = ds.MySQL("github_starred").String()

		// Postgres
		case "postgresql":
			output = ds.Postgres("github_starred").String()

		// ASCII - TabularGrid
		case "grid-default", "ascii-grid", "tabular-grid":
			output = ds.Tabular("grid" /* tablib.TabularGrid */).String()

		// ASCII - TabularSimple
		case "grid-simple", "ascii-simple", "tabular-simple":
			output = ds.Tabular("simple" /* tablib.TabularSimple */).String()

		// ASCII - TabularSiTabularCondensedmple
		case "grid-condensed", "ascii-condensed", "tabular-condensed":
			output = ds.Tabular("condensed" /* tablib.TabularCondensed */).String()

		}

		fmt.Println(output)

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
