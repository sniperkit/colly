package main

import (
	"fmt"
	"os"
	// "plugin"
	"strings"
	"time"

	"mvdan.cc/xurls"

	"github.com/sniperkit/colly/pkg"
	"github.com/sniperkit/colly/pkg/debug"
	"github.com/sniperkit/colly/pkg/queue"

	sm "github.com/sniperkit/colly/addons/sitemap"
	cfg "github.com/sniperkit/colly/pkg/config"
)

var (
	isDebug         bool = false
	isStrict        bool = true
	startedAt       time.Time
	sitemap         *sm.Sitemap
	q               *queue.Queue
	scraper         *colly.Collector
	detailCollector *colly.Collector
	libraries       []library
	collyConfig     *cfg.CollectorConfig
	entries         map[string]bool
	links           []string = []string{} // Array containing all the known URLs in a sitemap
)

// library stores information about a golang library
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

	if err != nil {
		log.Fatalln("error: ", err)
	}

	// log.Println("Loading cache engine")
	// cacheTTL, cacheStatus, cacheError = initCacheHTTP(defaultCacheEngine, defaultCacheFreqDuration, defaultCacheFreqUnit)
	// log.Printf("Result: TTL=%d, Status=%b, Error=%s", cacheTTL, cacheStatus, cacheError)

}

func main() {

	if collyConfig != nil {
		collyConfig = &cfg.CollectorConfig{}
		scraper = colly.NewCollectorWithConfig(collyConfig)
	} else {
		// Create a Collector specifically for Shopify
		scraper = colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
			colly.AllowedDomains(defaultAllowedDomains...),
			colly.DisallowedURLFilters(defaultDisabledURLFilters...),
			colly.URLFilters(defaultAllowedURLFilters...),
			colly.IgnoreRobotsTxt(),
			// colly.AllowURLRevisit(),
			// Cache responses to prevent multiple download of pages even if the collector is restarted
			colly.Debugger(&debug.LogDebugger{}),
			colly.CacheDir(defaultStorageCacheDir),
			// colly.Async(true),
			// MaxDepth is 2, so only the links on the scraped page and links on those pages are visited
			// colly.MaxDepth(2),
		)
	}

	log.Println("scraper.UserAgent=", scraper.UserAgent)
	log.Println("scraper.AllowedDomains=", scraper.AllowedDomains)
	log.Println("scraper.DisallowedURLFilters=", scraper.DisallowedURLFilters)
	log.Println("scraper.URLFilters=", scraper.URLFilters)
	log.Println("scraper.IgnoreRobotsTxt=", scraper.IgnoreRobotsTxt)
	log.Println("scraper.CacheDir=", scraper.CacheDir)

	/*
		writer, err := SafeCsvWriter("data.csv")
		if err != nil {
			log.Fatal("Failed to make data file")
		}
		defer writer.Flush()

		writer.Write([]string{"Date", "Headline"})
	*/

	// Create another collector to scrape page details
	// detailCollector = scraper.Clone()

	// Limit the maximum parallelism to 2
	// This is necessary if the goroutines are dynamically
	// created to control the limit of simultaneous requests.
	//
	// Parallelism can be controlled also by spawning fixed
	// number of go routines.
	scraper.Limit(&colly.LimitRule{
		Parallelism: 4,
		DomainGlob:  "*",
		// RandomDelay: 2 * time.Second,
	})

	// Find and visit next page links
	scraper.OnHTML(`li.page-item a[href]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(link)
		log.Println("`li.page-item a[href]` URL=", link, ", AbsURL=", e.Request.AbsoluteURL(link))

	})

	// On every a element which has href attribute call callback
	scraper.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// log.Println("pattern `a[href]` link:", link)
		/*
			if !strings.HasPrefix(link, "/repo/") || !strings.HasPrefix(link, "/category/") || !strings.HasPrefix(link, "/tag/") || !strings.HasPrefix(link, "/similar/") {
				log.Println("[SKIP] link:", link)
				return
			}
		*/

		// Print link
		log.Println("`a[href]` URL=", link, ", AbsURL=", e.Request.AbsoluteURL(link))
		entries[e.Request.AbsoluteURL(link)] = false // = append(entries, link)

		// Visit link found on page
		// Only those links are visited which are matched by  any of the URLFilter regexps
		// scraper.Visit(e.Request.AbsoluteURL(link))
		// e.Request.Visit(link)
	})

	// On every a HTML element which has name attribute call callback
	scraper.OnHTML(`div.col-md-8`, func(e *colly.HTMLElement) {
		log.Println("[LIST] link=", e.Request.URL)

		e.ForEach(".row", func(_ int, el *colly.HTMLElement) {
			// var stars, desc, name, uri string
			var uri string
			// name = strings.Replace(el.ChildText("div.description > a[href]"), "\n", " ", -1)
			// stars = strings.Replace(el.ChildText(".float-right"), "\n", " ", -1)
			// desc = strings.Replace(el.ChildText(".description"), "\n", " ", -1)
			if link, ok := el.DOM.Find("div.description > a").Attr("href"); ok {
				uri = strings.Replace(link, "\n", " ", -1)
				uri = strings.Replace(uri, "/repo/", "github.com/", 1)
			} else {
				return
			}

			log.Println("`div.col-md-8 > .row` PKG=", uri)

		})

	})

	scraper.OnHTML(`a.page-link`, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		log.Println("[VISIT] link=", link)
		// start scaping the page under the link found
		// e.Request.Visit(link)
		// Visit link found on page
		// Only those links are visited which are matched by  any of the URLFilter regexps

		// scraper.Visit(e.Request.AbsoluteURL(link))
		entries[e.Request.AbsoluteURL(link)] = false
		//e.Request.Visit(link)
	})

	// Before making a request print "Visiting ..."
	/*
		scraper.OnRequest(func(r *colly.Request) {
			fmt.Println("visiting", r.URL)
			if r.ID < 15 {
				r2, err := r.New("GET", fmt.Sprintf("%s?x=%v", url, r.ID), nil)
				if err == nil {
					q.AddRequest(r2)
				}
			}
		})
	*/

	scraper.OnRequest(func(r *colly.Request) {
		log.Println("[REQUEST] url=", r.URL.String())

	})

	scraper.OnError(func(r *colly.Response, e error) {
		log.Println("[ERROR] msg=", e, ", url=", r.Request.URL, ", body=", string(r.Body))
	})

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

			///?page=2
			// url = r.Request.URL.String()
			// if strings.Index(url, "/?page=") > -1 {

			log.Println("[ADD] URL=", url, ", AbsoluteURL=", r.Request.AbsoluteURL(url), ", contentType=", contentType)
			// Visit link found on page
			// Only those links are visited which are matched by  any of the URLFilter regexps
			scraper.Visit(r.Request.AbsoluteURL(url))
			// e.Request.Visit(link)
			//entries = append(entries, url)
			entries[url] = false

		}
	})

	for i := 1; i <= 11560; i++ {
		// Add URLs to the queue
		q.AddURL(fmt.Sprintf("https://golanglibs.com/?page=%d", i))
	}

	// Start the collector
	// scraper.Visit(defaultDomain)
	// scraper.Visit(defaultSitemapURL)

	// Async
	// scraper.Wait()

	// Consume URLs
	q.Run(scraper)

	log.Println("All URLs found:")
	for link, status := range entries {
		log.Println("\t link=", link, "status=", status)
	}

	log.Println("Collected", len(entries), "URLs")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(libraries)

}
