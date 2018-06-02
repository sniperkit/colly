package sitemap

import (
	"fmt"
	"log"
	"net/url"

	"github.com/sniperkit/colly/pkg"
	"github.com/sniperkit/colly/pkg/queue"

	"github.com/sniperkit/xutil/plugin/debug/pp"
)

func NewWithCollector(inputURL string, c *colly.Collector) (*Sitemap, error) {
	cs, err := New(inputURL)
	if err != nil {
		return nil, err
	}
	cs.collector = c
	return cs, nil
}

func (cs *Sitemap) Count() int {
	return len(cs.URLs)
}

func (cs *Sitemap) List() ([]url.URL, []string) {
	cs.URLs, _ = getURLs(cs.href)
	var urls []string
	for _, u := range cs.URLs {
		urls = append(urls, u.String())
	}
	return cs.URLs, urls
}

func (cs *Sitemap) Index() ([]url.URL, []string) {
	var sitemaps []string
	for _, sitemap := range cs.Indices {
		sitemaps = append(sitemaps, sitemap.String())
	}
	return cs.Indices, sitemaps
}

func (cs *Sitemap) Sitemaps() []url.URL {
	return cs.Indices
}

func (cs *Sitemap) getURLs() {
	if !cs.IsValid() {
		return
	}
	// var urls []url.URL
	urlsFromIndex, indexError := getURLsFromSitemapIndex(cs.href)
	if indexError == nil {
		cs.Indices = urlsFromIndex
	}

	urlsFromSitemap, sitemapError := getURLsFromSitemap(cs.href)
	if sitemapError == nil {
		cs.URLs = append(cs.URLs, urlsFromSitemap...)
	}

	// if isInvalidSitemapIndexContent(indexError) && isInvalidXMLSitemapContent(sitemapError) {
	// 	return nil, fmt.Errorf("%q is neither a sitemap index nor a XML sitemap", xmlSitemapURL.String())
	// }
}

func (cs *Sitemap) All() ([]url.URL, error) {
	if !cs.IsValid() {
		return nil, fmt.Errorf("sitemap at url='%q' is not reachable", cs.href.String())
	}
	var urls []url.URL
	urlsFromIndex, indexError := getURLsFromSitemapIndex(cs.href)
	if indexError == nil {
		urls = urlsFromIndex
		cs.Indices = urlsFromIndex
	}

	urlsFromSitemap, sitemapError := getURLsFromSitemap(cs.href)
	if sitemapError == nil {
		urls = append(urls, urlsFromSitemap...)
		cs.URLs = append(cs.URLs, urlsFromSitemap...)
	}

	if isInvalidSitemapIndexContent(indexError) && isInvalidXMLSitemapContent(sitemapError) {
		return nil, fmt.Errorf("%q is neither a sitemap index nor a XML sitemap", cs.href.String())
	}
	return urls, nil
}

func (cs *Sitemap) VisitAll() *colly.Collector {
	if !cs.IsValid() {
		return cs.collector
	}
	pp.Println("sitemapURL=", cs.href.String())
	links, err := getURLs(cs.href)
	if err != nil {
		log.Fatalln("error: ", err)
		return cs.collector
	}
	log.Println("links found:", len(links))
	for _, link := range links {
		log.Println("add -", link.String())
		cs.collector.Visit(fmt.Sprintf("%s", link.String())) // Request visit URL by Collector
	}
	return cs.collector
}

func (cs *Sitemap) EnqueueAll() {
	if !cs.IsValid() {
		return
	}
	pp.Println("sitemapURL=", cs.href.String())
	links, err := getURLs(cs.href)
	if err != nil {
		log.Fatalln("error: ", err)
		return
	}
	log.Println("links found:", len(links))
	for _, link := range links {
		log.Println("enqueue -", link.String())
		cs.cqueue.AddURL(fmt.Sprintf("%s", link.String())) // Enqueue new URL
	}
	return
}

func VisitAll(inputURL string, c *colly.Collector) *colly.Collector {
	sitemapURL, err := url.Parse(inputURL)
	if err != nil {
		log.Fatalln("error: ", err)
		return c
	}

	links, err := getURLs(*sitemapURL)
	if err != nil {
		log.Fatalln("error: ", err)
		return c
	}
	log.Println("links found:", len(links))

	for _, link := range links {
		log.Println("add -", link.String())
		c.Visit(fmt.Sprintf("%s", link.String())) // Request visit URL by Collector
	}
	return c
}

func EnqueueAll(inputURL string, q *queue.Queue) *queue.Queue {
	sitemapURL, err := url.Parse(inputURL)
	if err != nil {
		log.Fatalln("error: ", err)
		return q
	}
	pp.Println("sitemapURL=", sitemapURL)
	links, err := getURLs(*sitemapURL)
	if err != nil {
		log.Fatalln("error: ", err)
		return q
	}
	log.Println("links found:", len(links))
	for _, link := range links {
		log.Println("enqueue -", link.String())
		q.AddURL(fmt.Sprintf("%s", link.String())) // Enqueue new URL
	}
	return q
}
