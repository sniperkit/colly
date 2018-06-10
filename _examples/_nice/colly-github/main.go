package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	appDebug = false

	//-- End
)

// github api vars
var (

	// githubAPIAccount sets the github user name for api requests
	githubAPIAccount = "roscopecoltran"

	// githubAPIRepository sets the github repository for api requests
	githubAPIRepository = "sniperkit/colly"

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

	// githubAPIEndpointURLUserInfo
	githubAPIEndpointURLUserInfo = fmt.Sprintf("https://api.github.com/users/%s", githubAPIAccount)

	// githubAPIEndpointURL
	githubAPIEndpointURLStarred = fmt.Sprintf("https://api.github.com/users/%s/starred?%s", githubAPIAccount, strings.Join(githubAPIPaginationParams, "&"))

	//githubAPIEndpointURLRepo
	githubAPIEndpointURLRepo = fmt.Sprintf("https://api.github.com/repos/%s", githubAPIRepository)

	//-- End
)

// collector vars
var (

	// collectorDebug sets collector's debugger
	collectorDebug = false

	// collectorDebugger stores the collector's log event listener
	collectorDebugger *debug.LogDebugger = &debug.LogDebugger{}

	// - collectorJsonParser sets the json parser package to unmarshal JSON responses.
	//   - Available parsers:
	//     - `JSON` default golang "encoding/json" package. Important: This parser does not extract/flatten nested object headers
	//     - `MXJ` decodes / encodes JSON or XML to/from map[string]interface{}; extracts values with dot-notation paths and wildcards.
	//     - `GJSON` (Not Ready Yet), decodes JSON document; performs one line retrieval, dot notation paths, iteration, and parsing json lines.
	collectorJsonParser = "json"

	// collectorTabEnabled specifies if the collector load and marshall content-types that are tabular compatible
	// - `OnTAB` supported loading formats:
	//   - JSON (Sets + Books)
	//   - YAML (Sets + Books)
	//   - TOML (Sets + Books)
	//   - XML (Sets)
	//   - CSV (Sets)
	//   - TSV (Sets)
	// - IMPORTANT:
	//   - input must be marshallable as a slice interfaces ([]interface).
	//   - map[string]interface will be converted to []map[string]interface or []interface
	collectorTabEnabled = true

	// collectorDatasetOutputPrefixPath specifies the prefix path for all saved dumps
	collectorDatasetOutputPrefixPath = "./shared/dataset"

	// collectorDatasetOutputBasename specifies the template to use to write the dataset dump
	collectorDatasetOutputBasename = "colly_github_%d"

	// collectorDatasetOutputFormat sets the ouput format of the dataset extracted by the collector
	// `OnTAB` event export/print supported formats:
	//  - JSON (Sets + Books)
	//  - YAML (Sets + Books)
	//  - XLSX (Sets + Books)
	//  - XML (Sets + Books)
	//  - TSV (Sets)
	//  - CSV (Sets)
	//  - ASCII + Markdown (Sets)
	//  - MySQL (Sets)
	//  - Postgres (Sets)
	collectorDatasetOutputFormat = "tabular-grid" // "yaml"

	//  collectorSubDatasetColumns specifies the columns to filter from the json content
	collectorSubDatasetColumns = []string{"id", "name", "full_name", "description", "language", "stargazers_count", "forks_count"}
)

// init() function is executed when the executable is started, before function main()
func init() {
	// Ensure that the output format is set in lower case.
	collectorDatasetOutputFormat = strings.ToLower(collectorDatasetOutputFormat)
}

// descriptionLength function is a callback used to append dynamic column to a `OnTab` event dataset
func descriptionLength(row []interface{}) interface{} {
	if row == nil {
		return 0
	}
	//if appDebug {
	//	fmt.Printf("\n----- Calculated the description length for row:\n")
	//	prettyPrint(row)
	//}
	if len(row) < 2 {
		return 0
	}
	return len(asString(row[2]))
}

func asString(vv interface{}) string {
	var v string
	switch vv.(type) {
	case string:
		v = vv.(string)
	case int:
		v = strconv.Itoa(vv.(int))
	case int64:
		v = strconv.FormatInt(vv.(int64), 10)
	case uint64:
		v = strconv.FormatUint(vv.(uint64), 10)
	case bool:
		v = strconv.FormatBool(vv.(bool))
	case float64:
		v = strconv.FormatFloat(vv.(float64), 'G', -1, 32)
	case json.Number:
		v = vv.(json.Number).String()
	case time.Time:
		v = vv.(time.Time).Format(time.RFC3339)
	default:
		v = fmt.Sprintf("%s", v)
	}
	return v
}

// prettyPrint wraps debug message with `github.com/k0kubun/pp` package functions.
func prettyPrint(output ...interface{}) {
	pp.Println(output...)
}

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: api.github.com
		colly.AllowedDomains("api.github.com"),
		colly.AllowTabular(collectorTabEnabled),
	)

	// UseJsonParser. Available: mxj, gjson, json (default)
	c.UseJsonParser = collectorJsonParser

	if collectorDebug {
		c.SetDebugger(collectorDebugger)
	}

	// On every a element which has tabular format data call callback
	// Notes:
	// - `OnTAB` callback event are enabled only if the `AllowTabular` attribute is set to true.
	// - `OnTAB` use a fork of the package `github.com/agrison/go-tablib`
	// - `OnTAB` query specifications are available in 'SPECS.md'

	c.OnTAB("0:0", func(e *colly.TABElement) {
		// c.OnTAB(`cols=(:5), rows=(1:7)`, func(e *colly.TABElement) {
		// c.OnTAB(`cols[0:5], rows[1:7]`, func(e *colly.TABElement) {
		// c.OnTAB(`rows=(1,10), cols=("id", "name", "full_name", "description", "language", "stargazers_count", "forks_count")`, func(e *colly.TABElement) {
		// c.OnTAB(`rows=(1,10), cols=("id", "name", "full_name", "description", "language", "stargazers_count", "forks_count")`, func(e *colly.TABElement) {

		// Debug the dataset slice
		if appDebug {
			fmt.Println("Valid=", e.Dataset.Valid(), "Height=", e.Dataset.Height(), "Width=", e.Dataset.Width())
			pp.Printf("Headers: \n %s \n\n", e.Dataset.Headers())
		}

		// checkHeader selection
		var datasetGroup string
		var hdrSelect, hdrNotFound, hdrFound []string

		// Better url matching/pattern ?! ^^
		switch {
		case strings.Contains(e.Request.URL.String(), "/starred"):
			fmt.Println("Starred dataset...")
			hdrSelect = []string{"id", "full_name", "description", "language", "stargazers_count", "owner_login", "owner_id", "updated_at"}
			datasetGroup = "starred"

		case strings.Contains(e.Request.URL.String(), "/repos/"):
			fmt.Println("Repositories dataset...")
			hdrSelect = []string{"id", "full_name", "description", "language", "stargazers_count", "watchers_count", "owner_login", "owner_id"}
			datasetGroup = "repos"

		case strings.Contains(e.Request.URL.String(), "/users/"):
			fmt.Println("Users dataset...")
			hdrSelect = []string{"id", "login", "avatar_url", "blog", "created_at", "hireable", "following", "followers"}
			datasetGroup = "users"

		}

		hdrNotFound, hdrFound = e.Dataset.HeadersExists(hdrSelect...)
		if len(hdrNotFound) > 0 {
			pp.Printf("Headers - NotFound: \n %s \n\n", hdrNotFound)
			pp.Printf("Headers - Found  \n %s \n\n", hdrFound)
		}

		ds, err := e.Dataset.Select(0, 0, hdrFound...)
		if err != nil {
			fmt.Println("error:", err)
		}

		switch datasetGroup {
		case "starred":
		case "repos":
			// Add a dynamic column, by passing a function which has access to the current row, and must return a value:
			ds.AppendDynamicColumn("description_length", descriptionLength)
		case "users":
		default:
		}

		// Export dataset
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
		case "markdown", "tabular-markdown":
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

		// output final export
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

	// 3. Start scraping repository info; json object with nested objects
	c.Visit(githubAPIEndpointURLRepo)

	// 1. Start scraping user info; json object with no nested objects or arrays
	c.Visit(githubAPIEndpointURLUserInfo)

	// 2. Start scraping starred repository; json array with nested objects
	c.Visit(githubAPIEndpointURLStarred)

}
