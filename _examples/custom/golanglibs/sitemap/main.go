package main

import (
	"fmt"
	"strings"
	"time"

	// "mvdan.cc/xurls"

	"github.com/sniperkit/colly/pkg"
	// "github.com/sniperkit/colly/pkg/debug"
	"github.com/sniperkit/colly/pkg/queue"

	sm "github.com/sniperkit/colly/addons/sitemap"
	cfg "github.com/sniperkit/colly/pkg/config"
)

var (
	isDebug         bool   = false
	isStrict        bool   = true
	isVerbose       bool   = true
	sitemapURL      string = "https://golanglibs.com/sitemap.txt"
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

	q, err = initQueue(2, 1000000, "InMemory")
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
			// colly.Debugger(&debug.LogDebugger{}),
			colly.CacheDir(defaultStorageCacheDir),
			// colly.CacheHTTP(defaultStorageCacheDir),
			// colly.Async(true),
			// MaxDepth is 2, so only the links on the scraped page and links on those pages are visited
			// colly.MaxDepth(2),
		)
	}

	writer, err := newSafeCsvWriter("export.csv")
	if err != nil {
		log.Fatal("Failed to make data file")
	}
	defer writer.Flush()

	writer.Delimiter('|').Write([]string{"url", "package_uri", "name", "description", "stars", "tags"})

	/*
		scraper.Limit(&colly.LimitRule{
			Parallelism: 4,
			DomainGlob:  "*",
			// RandomDelay: 2 * time.Second,
		})
	*/

	// On every a element which has href attribute call callback
	scraper.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		cds.Append("links", e.Request.AbsoluteURL(link))
		// scraper.Visit(e.Request.AbsoluteURL(link))
	})

	// On every a HTML element which has name attribute call callback
	scraper.OnHTML(`div.col-md-8`, func(e *colly.HTMLElement) {

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
			log.Infof("[PKG] uri='%s', pkg='%s', name='%s', stars='%s', desc='%s', tags='%s'\n", url, pkg, name, stars, desc, tagsStr)
			writer.Write([]string{url, pkg, name, desc, stars, tagsStr})

			// Similar PKGs ?!

		})

	})

	scraper.OnRequest(func(r *colly.Request) {
		log.Infoln("[REQUEST] url=", r.URL.String())
	})

	scraper.OnError(func(r *colly.Response, e error) {
		log.Println("[ERROR] msg=", e, ", url=", r.Request.URL, ", body=", string(r.Body))
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

				///?page=2
				// url = r.Request.URL.String()
				// if strings.Index(url, "/?page=") > -1 {

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

	for i := 1; i <= 11560; i++ {
		q.AddURL(fmt.Sprintf("https://golanglibs.com/?page=%d", i)) // Add URLs to the queue
	}

	// Async
	// scraper.Wait()

	/*
		links, err := linksFromCSV(sitemapURL)
		check(err)

		for _, link := range links {
			q.AddURL(link)
		}
	*/

	// Consume URLs
	q.Run(scraper)

	/*
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")

		// Dump json to the standard output
		enc.Encode(libraries)
	*/

}
