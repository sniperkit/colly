package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	// core
	colly "github.com/sniperkit/colly/pkg"
	cfg "github.com/sniperkit/colly/pkg/config"
	proxy "github.com/sniperkit/colly/pkg/proxy/default"
	queue "github.com/sniperkit/colly/pkg/queue"

	// experimental addons

	//// console UI
	cui "github.com/sniperkit/colly/addons/dashboard/gocui"
	tui "github.com/sniperkit/colly/addons/dashboard/termui"
	// dash "github.com/sniperkit/colly/addons/dashboard"
	// tvi "github.com/sniperkit/colly/addons/dashboard/tview"

	//// proxies
	onion "github.com/sniperkit/colly/addons/proxy/onion"

	//// sitemaps
	sm "github.com/sniperkit/colly/addons/sitemap"

	//// stats
	ta "github.com/sniperkit/colly/addons/stats/tachymeter"
)

var version = APP_VERSION

/*
	Refs:
	- https://github.com/zdavep/dozer
	- https://github.com/yuansudong/msg_center/blob/master/node/sentinel/workpool.go
	- https://github.com/xor-gate/snippets/blob/master/golang/channel-sync/main.go
	- https://github.com/v3io/http_blaster/blob/master/http_blaster.go
	- https://github.com/prgrm0x1/cmh/blob/master/cmh.go
	- https://github.com/test-circle-provisioner/go-sample-template/blob/master/cmd/service/main.go
	- https://github.com/jamiealquiza/tachymeter
*/

var (
	// features activation
	isPolite  bool = true
	isProxy   bool = false
	isDebug   bool = false
	isStrict  bool = true
	isVerbose bool = true
	worker_qd int  = 10000

	// collector
	scraper              *colly.Collector
	detailCollector      *colly.Collector
	collyConfig          *cfg.Config = &cfg.Config{}
	currentCrawlerMode   string      = "queue"
	availableCrawlerMode []string    = []string{"queue", "async", "distributed", "default"}
	q                    *queue.Queue

	// sitemaps
	sitemapURL string = "https://golanglibs.com/sitemap.txt"
	exportFile string = "./shared/storage/exports/reports/latest.csv"
	sitemap    *sm.Sitemap

	// tui/cui
	xConsoleUI *cui.TermUI
	xTermUI    *cui.TermUI

	// tachymeter
	startedAt            time.Time
	isTachymeter         bool = true
	isTachymeterParallel bool = false
	cTachymeter          chan *ta.Tachymeter
	xTachymeter          *ta.Tachymeter
	xTachyResults        *ta.Metrics
	xTachymeterTL        ta.Timeline
	wallTimeStart        time.Time

	// proxy list
	xProxyList *onion.ProxyList

	// channels
	ch_done                chan struct{}
	stopTheCrawler         chan bool
	allURLsHaveBeenVisited chan bool
	crawlResult            chan error
	xResults               chan tui.WorkResult

	// extracted patterns
	libraries []library
	entries   map[string]bool
	links     []string = []string{} // Array containing all the known URLs in a sitemap
)

// library stores information about indexed golang library in golanglibs.com
type library struct {
	Title       string
	Description string
	Categories  []string
	URI         string
	URL         string
	Stars       int
}

func init() {

	var err error
	startedAt = time.Now()
	sm.CreateDirs(defaultStorageDirs)
	libraries = make([]library, 0)
	entries = make(map[string]bool, 0)

	q, err = initQueue(numberOfWorkers, queueMaxSize, "InMemory")
	if err != nil {
		log.Fatalln("error: ", err)
	}

	const (
		default_conf         = "example.toml"
		usage_conf           = "conf file path"
		usage_version        = "show version"
		default_showversion  = false
		usage_results_file   = "results file path"
		default_results_file = "example.results"
		usage_log_file       = "enable stdout to log"
		default_log_file     = true
		default_worker_qd    = 10000
		usage_worker_qd      = "queue depth for worker requests"

		usage_verbose   = "print debug logs"
		default_verbose = false

		usage_memprofile   = "write mem profile to file"
		default_memprofile = false

		usage_cpuprofile   = "write cpu profile to file"
		default_cpuprofile = false

		usage_enable_ui   = "enable terminal ui"
		default_enable_ui = true

		usage_dump_failures   = "enable 4xx status requests dump to file"
		defaule_dump_failures = false

		usage_dump_location   = "location of dump requests"
		default_dump_location = "."
	)

	flag.StringVar(&conf_file, "conf", default_conf, usage_conf)
	flag.StringVar(&conf_file, "c", default_conf, usage_conf+" (shorthand)")
	flag.StringVar(&results_file, "o", default_results_file, usage_results_file+" (shorthand)")
	flag.BoolVar(&showVersion, "version", default_showversion, usage_version)
	flag.BoolVar(&cpu_profile, "p", default_cpuprofile, usage_cpuprofile)
	flag.BoolVar(&mem_profile, "m", default_memprofile, usage_memprofile)
	flag.BoolVar(&enable_log, "d", default_log_file, usage_log_file)
	flag.BoolVar(&verbose, "v", default_verbose, usage_verbose)
	flag.IntVar(&worker_qd, "q", default_worker_qd, usage_worker_qd)
	flag.BoolVar(&enable_ui, "u", default_enable_ui, usage_enable_ui)
	flag.BoolVar(&dump_failures, "f", defaule_dump_failures, usage_dump_failures)
	flag.StringVar(&dump_location, "l", default_dump_location, usage_dump_location)

}

func main() {

	// Create a channels for the collector results
	allStatisticsHaveBeenUpdated = make(chan bool)
	allURLsHaveBeenVisited = make(chan bool)
	stopTheCrawler = make(chan bool)
	crawlResult = make(chan error)

	// Create a Collector
	scraper = colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains(defaultAllowedDomains...),
		colly.DisallowedURLFilters(defaultDisabledURLFilters...),
		colly.URLFilters(defaultAllowedURLFilters...),
		colly.IgnoreRobotsTxt(),
		// colly.AllowURLRevisit(),
		// Cache responses to prevent multiple download of pages even if the collector is restarted
		// colly.Debugger(&debug.LogDebugger{}),
		colly.CacheDir(defaultStorageCacheDir),
		// colly.CacheHTTP(defaultStorageCacheDir),
		// colly.Async(true),
		// MaxDepth is 2, so only the links on the scraped page and links on those pages are visited
		// colly.MaxDepth(2),
	)

	// Create a Tachymeter
	if isTachymeter {
		cTachymeter = make(chan *ta.Tachymeter)
		xTachymeterTL = ta.Timeline{}
		xTachymeter = ta.New(
			&ta.Config{
				// Tachymeter
				SafeMode:   true, // deprecated
				SampleSize: 50,
				HBins:      10,
				Export: &ta.Export{
					// Exports
					Encoding:   "tsv",
					Basename:   "golanglibs_tachymter_%d",
					PrefixPath: "./shared/exports/stats/tachymeter/",
					SplitLimit: 2500,
					BufferSize: 20000,
					Overwrite:  true,
					BackupMode: true,
				},
			},
		)
		if currentCrawlerMode == "" || currentCrawlerMode == "" {
			wallTimeStart = time.Now()
			isTachymeterParallel = true
		}
	}

	if cpu_profile || mem_profile {
		isDebug = true
	}

	// For tests
	collyConfig.Title = APP_NAME
	collyConfig.IsModeTUI = enable_ui
	collyConfig.VerboseMode = enable_log
	collyConfig.DebugMode = isDebug

	// Create channels to check statuses of TermUI components
	if enable_ui {
		xResults = make(chan tui.WorkResult)
		stopTheUI = make(chan bool)
		// ch_done = enable_tui()
		// ch_done = tui.Dashboard()
		// uiWaitGroup.Add(1)
		go func() {
			tui.Dashboard(stopTheUI, stopTheCrawler)
			// uiWaitGroup.Done()
		}()
		// defer stop_cpu_profile()
		// defer write_mem_profile()
		// start_cpu_profile()
	}

	// Export golanglibs' indexed packages
	ensurePathExists(exportFile)
	writer, err := newSafeCsvWriter(exportFile)
	if err != nil {
		log.Fatal("Failed to make data file")
	}
	defer writer.Flush()
	writer.Delimiter('|').Write([]string{"package_uri", "referrer", "name", "description", "stars", "tags"})

	// Set a rate limiter to the collector instanciated
	if isPolite {
		scraper.Limit(&colly.LimitRule{
			Parallelism: 4,
			DomainGlob:  "*",
			RandomDelay: 5 * time.Second,
		})
	}

	// Set a custom httpTransport for requests using httpcache
	if isCacheTransport {
		xCache, xTransport = newCacheWithTransport("badger", "./shared/storage/cache/http")
		scraper.WithTransport(xTransport)
	}

	// Set a list of proxies for scraping/crawling the web content
	if isProxy {
		// Rotate two socks5 proxies (Add Tor Proxies)
		rp, err := proxy.RoundRobinProxySwitcher(
			"socks5://127.0.0.1:1337",
			"socks5://127.0.0.1:1338",
		)
		if err != nil {
			log.Fatal(err)
		}
		scraper.SetProxyFunc(rp)
	}

	// On every a element which has href attribute call callback
	/*
		scraper.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			absLink := e.Request.AbsoluteURL(link)
			cds.Append("links", e.Request.AbsoluteURL(link))
			scraper.Visit(e.Request.AbsoluteURL(link))

			if !enable_ui {
				log.Infof("[LINK] link='%s', absLink='%s'\n", link, absLink)
			}
		})
	*/

	// On every a HTML element which has name attribute call callback
	scraper.OnHTML(`div.col-md-8`, func(e *colly.HTMLElement) {
		start := time.Now()
		parentURL := e.Request.AbsoluteURL(e.Request.URL.String())

		var meta []string
		e.ForEach(".row", func(_ int, el *colly.HTMLElement) {
			var url, pkg, name, stars, desc, tagsStr string
			name = strings.Replace(el.ChildText("div.description > a[href]"), "\n", " ", -1)

			stars = strings.Replace(el.ChildText(".float-right"), "\n", " ", -1)
			stars = rNumber.FindString(stars)

			desc, _ = el.DOM.Find(".description").Attr("title")
			desc = strings.Replace(desc, "\n", " ", -1)

			if link, ok := el.DOM.Find("div.description > a").Attr("href"); ok {
				pkg = strings.Replace(link, "\n", " ", -1)
				url = el.Request.AbsoluteURL(pkg)
				pkg = strings.Replace(pkg, "/repo/", "github.com/", 1)
			} else {
				return
			}

			var tags []string
			el.ForEach("span.badge.badge-secondary", func(_ int, et *colly.HTMLElement) {
				if et.Text != "" {
					tag := strings.Replace(et.Text, "\n", "", -1)
					tag = strings.TrimSpace(tag)
					tags = append(tags, tag)
				}
			})
			tagsStr = strings.Join(tags, ",")

			if !enable_ui {
				log.Infof("[PKG] uri='%s', pkg='%s', name='%s', stars='%s', desc='%s', tags='%s'\n", url, pkg, name, stars, desc, tagsStr)
			}
			meta = []string{pkg, url, name, desc, stars, tagsStr}
			writer.Write(meta)
		})

		xTachymeter.AddTimeWithLabel(parentURL, time.Since(start))

		// Add each loop tachymeter to the event timeline.
		// xTachymeterTL.AddEvent(xTachymeter.Snapshot())
	})

	scraper.OnResponse(func(r *colly.Response) {
		if !enable_ui {
			log.Infoln("[REQUEST] url=", r.Request.URL.String())
		} else {
			xResults <- tui.WorkResult{
				URL:             *r.Request.URL, //.String(), //*r.Request.URL,
				NumberOfWorkers: numberOfWorkers,
				ResponseSize:    r.GetSize(),
				StatusCode:      r.GetStatusCode(),
				StartTime:       r.GetStartTime(),
				EndTime:         r.GetEndTime(),
				ContentType:     r.GetContentType(),
			}
		}
	})

	scraper.OnError(func(r *colly.Response, e error) {
		if !enable_ui {
			log.Println("[ERROR] msg=", e, ", url=", r.Request.URL, ", body=", string(r.Body))
		} else {
			xResults <- tui.WorkResult{
				Err:             e,
				URL:             *r.Request.URL,
				NumberOfWorkers: numberOfWorkers,
				ResponseSize:    r.GetSize(),
				StatusCode:      r.GetStatusCode(),
				StartTime:       r.GetStartTime(),
				EndTime:         r.GetEndTime(),
				ContentType:     r.GetContentType(),
			}
		}
	})

	/*
		// Before making a request print "Visiting ..."
		scraper.OnResponse(func(r *colly.Response) {
			contentType := r.Headers.Get("Content-Type")
			var urls []string
			if isStrict {
				urls = xurls.Relaxed().FindAllString(string(r.GetBody()), -1)
			} else {
				urls = xurls.Strict().FindAllString(string(r.GetBody()), -1)
			}

			for _, url := range urls {

				// url = r.Request.AbsoluteURL(url)
				// url = r.URL.String()

				// url = r.Request.URL.String()
				// if strings.Index(url, "/?page=") > -1 {

				// TUI_EVENT
				// log.Println("[ADD] URL=", url, ", AbsoluteURL=", r.Request.AbsoluteURL(url), ", contentType=", contentType)
				// Visit link found on page
				// Only those links are visited which are matched by  any of the URLFilter regexps
				scraper.Visit(r.Request.AbsoluteURL(url))
				// e.Request.Visit(link)
				//entries = append(entries, url)
				entries[url] = false

			}
		})
	*/

	numPages := 30 // 11560
	for i := 1; i <= numPages; i++ {
		q.AddURL(fmt.Sprintf("https://golanglibs.com/?page=%d", i)) // Add URLs to the queue
	}

	/*
		// enqueue a list urls to visit manually from a csv file
		links, err := linksFromCSV(sitemapURL)
		check(err)
		for _, link := range links {
			q.AddURL(link)
		}
	*/

	// update the statistics with the results
	if enable_ui {
		go func() {
			for {
				select {
				case <-allURLsHaveBeenVisited:
					allStatisticsHaveBeenUpdated <- true
					return

				case <-stopTheCrawler:
					stopTheUI <- true
					// allURLsHaveBeenVisited <- true

				case result := <-xResults:
					tui.UpdateStatistics(result)
					// cds.Append("urls", url.String())

				}
			}
		}()
	}

	switch currentCrawlerMode {
	case "async":
		scraper.Wait() // Async
	case "queue":
		q.Run(scraper) // Consume URLs
	default:
		scraper.Visit(defaultDomain)
	}

	if enable_ui {
		allURLsHaveBeenVisited <- true
		stopTheUI <- true
	}

	if isTachymeter {
		// Write out an HTML page with the histogram for all iterations.
		err := xTachymeterTL.WriteHTML("./shared/exports/stats/tachymeter")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Results written")

		// Print JSON format to console.
		// fmt.Printf("%s\n\n", xTachyResults.JSON())

		// Print pre-formatted console output.
		// fmt.Printf("%s\n\n", xTachyResults.String())

		// Print text histogram.
		// fmt.Println(xTachyResults.Histogram.String(15))

	}

}