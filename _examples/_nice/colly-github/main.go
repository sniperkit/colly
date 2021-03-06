package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	// colly - core
	colly "github.com/sniperkit/colly/pkg"
	config "github.com/sniperkit/colly/pkg/config"
	debug "github.com/sniperkit/colly/pkg/debug"

	// colly - plugins
	pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

/*
	To do asap:
	- Tuesday:
		- [x] Dump load config into different file formats
		- [ ] Add header link parser
		- [ ] Add Plucker units
		- [ ] Add yql for tab data query parser
		- [ ] Add ranger parser
			- [ ] github.com/sniperkit/rangetype
		- [ ] Add TextQL/JsonQL
			- [ ] github.com/davecb/jxpath
			- [ ] github.com/bloglovin/obpath
		- [ ]Add custom query lexer with lexmachine
	- Thrusday:
		- [ ] Create data blocks templates
		- [ ] Create databooks form nested json arrays
			- [ ] Develop json pointers
		- [ ] Create databooks form nested csv arrays
			- [ ] Develop csv pointers
	- Friday:
		- [ ] Concurrent CSV writers
		- [ ] HttpCache Transport
		- [ ] HttpStats Transport
	- Saturday:
		- [ ] Extract meta
		- [ ] Extract open graph metadata
		- [ ] To do list manager
			- [ ] https://github.com/jasoncodingnow/todos
			- [ ] https://github.com/izqui/todos/blob/master/git.go
			- [ ] https://github.com/gerad/release-notes
*/

// collyVersion specifies the colly's version imported in the current executable.
// If the executable is built with Makefile, the collyVersion will use the actual repo's short tag version
var collyVersion = ""

// app vars
var (

	// appVersion specifies the version of the app. If the executable is built with Makefile, the appVersion will use the actual repo's short tag version
	appVersion = "0.0.1-alpha"

	// appCurrentDir specifies the directory of the currently file/binary is running
	appCurrentDir = "."

	// appXGDBDir specifies the XGDB directory of the local workstation
	appXGDBDir = "~/.colly"

	// appBaseDir specifies the directory name to set inside the appXGDBDir
	appBaseDir = "github-collector"

	// appDebug specifies if the app debug/verbose some development event logged
	appDebug = false

	// appConfigDump specifies whether the loaded configuration has to be dumped into various file formats.
	appConfigDump = true

	// appConfigDirDump specifies the prefix path for
	appConfigDirDump = "./shared/config/dumps/"

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

)

// init() function is executed when the executable is started, before function main()
func init() {
	// Ensure that the output format is set in lower case.
	collectorDatasetOutputFormat = strings.ToLower(collectorDatasetOutputFormat)
}

func panic(err error) {
	if err != nil {
		fmt.Println("Could not get the directory of the file/binary currently running...")
		os.Exit(1)
	}
}

// https://github.com/sniperkit/xaggregate/blob/sniperkit/plugin/provider/github/utils.go
func dateFormat(row []interface{}) interface{} {
	if row == nil {
		return 0
	}
	if len(row) < 7 {
		return 0
	}

	date, err := time.Parse("2006-01-02", asString(row[7]))
	if err != nil {
		return "0000-00-00"
	}
	return date.Format("2006-01-02")
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

	// appConfig stores a nested structure of settings for the collector and other components.
	// Important!
	// - You can define several configuration files and in different formats
	// - Formats available: yaml, json, xml, toml
	appConfig, err := config.NewFromFile(false, false, false, "./shared/config/global.yaml")
	panic(err)
	if appDebug {
		prettyPrint(appConfig)
	}

	// TODO: create colly.NewCollectorWithConfig() method
	// c := colly.NewCollectorWithConfig(appConfig)

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
	// Examples:
	// c.OnTAB(`cols=(:5), rows=(1:7)`, func(e *colly.TABElement) { }
	// c.OnTAB(`cols=(:5), rows=(1:7)`, func(e *colly.TABElement) { }
	// c.OnTAB(`cols[0:5], rows[1:7]`, func(e *colly.TABElement) { }
	// c.OnTAB(`rows=(1,10), cols=("id", "name", "full_name", "description", "language", "stargazers_count", "forks_count")`, func(e *colly.TABElement) { }
	// c.OnTAB(`rows=(1,10), cols=("id", "name", "full_name", "description", "language", "stargazers_count", "forks_count")`, func(e *colly.TABElement) { }
	// params ... interface{}

	// OnDATA
	tabHooks := &colly.TABHooks{
		Enabled:  true,
		Registry: make(map[string]*colly.TABHook, 3),
	}

	// TABHook - default
	// URL: https://api.github.com/
	selectorDefault := &colly.TABHook{
		Writer: &colly.TABWriter{
			Basename:   "colly",
			PrefixPath: "./shared/exports",
			Concurrent: true,
			Split:      true,
			SplitAt:    2500,
			Formats:    []string{"yaml", "json", "csv"},
		},
		Printer: &colly.TABOutput{
			Format: "tabular-grid",
		},
	}
	// selectorDefault.Valid()
	tabHooks.Registry[selectorDefault.ID()] = selectorDefault

	// TABHook - users
	// URL: https://api.github.com/users/roscopecoltran
	selectorUsers := &colly.TABHook{
		Id: "users",
		Slicer: &colly.TABSlicer{
			Headers: []string{"id", "login", "avatar_url", "blog", "created_at", "hireable", "following", "followers"},
			Cols: &colly.TABRanger{
				Expr: "[::]",
			},
			Rows: &colly.TABRanger{
				Lower: 0,
				Upper: 0,
				Cap:   0,
			},
		},
		Printer: &colly.TABOutput{
			Format: "tabular-grid",
		},
	}
	selectorUsers.PatternRegexp("/users/([a-zA-Z0-9\\-_]+)$")
	// selectorUsers.Valid()
	tabHooks.Registry[selectorUsers.ID()] = selectorUsers

	// TABHook - repos
	// URL: https://api.github.com/repos/sniperkit/colly
	selectorRepos := &colly.TABHook{
		Id: "repos",
		Slicer: &colly.TABSlicer{
			Headers: []string{"id", "full_name", "description", "language", "stargazers_count", "watchers_count", "owner_login", "owner_id"},
			Cols: &colly.TABRanger{
				Expr: "[::]",
			},
			Rows: &colly.TABRanger{
				Lower: 0,
				Upper: 0,
				Cap:   0,
			},
		},
		Printer: &colly.TABOutput{
			Format: "tabular-grid",
		},
	}
	selectorRepos.PatternRegexp("/repos/([a-zA-Z0-9\\-_]+)/([a-zA-Z0-9\\-_]+)$")
	// selectorRepos.Valid()
	tabHooks.Registry[selectorRepos.ID()] = selectorRepos

	// TABHook - starred
	// URL: https://api.github.com/users/roscopecoltran/starred?page=1&per_page=10&direction=desc&sort=updated
	selectorStarred := &colly.TABHook{
		Id: "starred",
		Slicer: &colly.TABSlicer{
			Headers: []string{"id", "full_name", "description", "language", "owner_id", "stargazers_count", "updated_at"},
			Cols: &colly.TABRanger{
				Expr: "[::]",
			},
			Rows: &colly.TABRanger{
				Lower: 0,
				Upper: 0,
				Cap:   0,
			},
		},
		Printer: &colly.TABOutput{
			Format: "tabular-grid",
		},
	}

	selectorStarred.AddDynamicColumn(descriptionLength, "desc_length", 2)
	// selectorStarred.AddDynamicColumn(dateFormat, "updated_at", 7)

	selectorStarred.PatternRegexp("/users/([a-zA-Z0-9\\-_]+)/starred")
	// selectorStarred.Valid()
	tabHooks.Registry[selectorStarred.ID()] = selectorStarred

	// dumpFormats specifies...
	dumpFormats := []string{"json", "yaml"}

	// dumpNodes specifies the config sections to dump
	// - if empty; only a merged config file will be written.
	// - if set to `all`; a global and a sub-set of config files per components will be written.
	// - if set to `collector,export,dataset`; only config files for the defined components will be written.
	// - to check the sections available; use the config.Sections() method.
	dumpNodes := []string{"inspect", "collector", "app", "debug", "filters", "collection", "hooks", "dirs"}

	// TABHook - header-link
	// selectorHeaderLink.PatternRegexp("<([^>]+)>;\\s+rel=\"([^i\"]+)\"")

	// SetHooks
	c.SetHooks(tabHooks)

	/*
		// pluckers defines
		pluckers := []*colly.Pluck{
			{
				Activators:  []string{"a", "b"},
				Deactivator: "f",
			},
			{
				Activators:  []string{"a", "c", "d"},
				Deactivator: "z",
			},
		}

		// SetPlucks
		c.SetPluckers(pluckers)
	*/

	// SetExtractors (not implemented yet...)
	// c.SetExtractors(extractors)

	// DumpConfig
	c.DumpConfig(dumpNodes, dumpFormats, "./shared/config/schema") // use string slices

	// OnDATA (either hooks matched by RequestURI() or )
	c.OnDATA(tabHooks, func(e *colly.TABElement) {

		// datasetGroup
		var datasetGroup string = "colly"

		// outputFormat
		outputFormat := collectorDatasetOutputFormat

		// check if a hook is attached
		if e.Hook != nil {
			if appDebug {
				pp.Printf("Hook.Id: \n %s \n\n", e.Hook.Id)
				pp.Printf("Hook.PatternURL: \n %s \n\n", e.Hook.PatternURL)
			}
			if len(e.Hook.Headers) > 0 {
				if appDebug {
					pp.Printf("Hook.Headers: \n %s \n\n", e.Hook.Headers)
				}
			}
			if e.Hook.Slicer != nil {
				if appDebug {
					pp.Printf("Hook.Slicer: \n %s \n\n", e.Hook.Slicer)
				}
			}
			if e.Hook.Writer != nil {
				if appDebug {
					pp.Printf("Hook.Writer: \n %s \n\n", e.Hook.Writer)
				}
			}
			if e.Hook.Printer != nil {
				if appDebug {
					pp.Printf("Hook.Printer: \n %s \n\n", e.Hook.Printer)
				}
				if e.Hook.Printer.Format != "" {
					outputFormat = e.Hook.Printer.Format
				}
			}
			datasetGroup = e.Hook.Id
		}

		// Debug the dataset slice
		if appDebug {
			pp.Println("Dataset=", datasetGroup, "Format=", outputFormat, "Valid=", e.Dataset.Valid(), "Height=", e.Dataset.Height(), "Width=", e.Dataset.Width())
			pp.Printf("Headers: \n %s \n\n", e.Dataset.Headers())
		}

		// dataset pre-process
		ds := e.Dataset

		// dataset post-process
		/*
			// checkHeader selection
			var hdrSelect, hdrNotFound, hdrFound []string

			// datasetGroup (default: colly)
			switch datasetGroup {
			case "starred":
				fmt.Println("Starred dataset...")
				// hdrSelect = []string{"id", "full_name", "description", "language", "stargazers_count", "owner_login", "owner_id", "updated_at"}

			case "repos":
				fmt.Println("Repositories dataset...")
				// hdrSelect = []string{"id", "full_name", "description", "language", "stargazers_count", "watchers_count", "owner_login", "owner_id"}

			case "users":
				fmt.Println("Users dataset...")
				// hdrSelect = []string{"id", "login", "avatar_url", "blog", "created_at", "hireable", "following", "followers"}

			}

			// e.Hook.Headers
			// Dataset.HeadersExists
			hdrNotFound, hdrFound = e.Dataset.HeadersExists(hdrSelect...)
			if len(hdrNotFound) > 0 {
				pp.Printf("Headers - NotFound: \n %s \n\n", hdrNotFound)
				pp.Printf("Headers - Found  \n %s \n\n", hdrFound)
			}

			// Dataset.Select
			ds, err := e.Dataset.Select(0, 0, hdrFound...)
			if err != nil {
				fmt.Println("error:", err)
			}
		*/

		// Hook.DynamicColumns
		for _, dc := range e.Hook.DynamicColumns {
			ds.AppendDynamicColumn(dc.Header, dc.Func)
		}

		// Export dataset
		// ds.EXPORT_FORMAT().String() 					--> returns the contents of the exported dataset as a string.
		// ds.EXPORT_FORMAT().Bytes() 					--> returns the contents of the exported dataset as a byte array.
		// ds.EXPORT_FORMAT().WriteTo(writer) 			--> writes the exported dataset to w.
		// ds.EXPORT_FORMAT().WriteFile(filename, perm) --> writes the databook or dataset content to a file named by filename.
		var output string
		switch outputFormat {
		// YAML
		case "yaml", "yml":
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

	// OnResponse
	c.OnResponse(func(r *colly.Response) {

		// Pluck content ?!

		// Parse headers ?!

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
