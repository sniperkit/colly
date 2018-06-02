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

func (cs *Sitemap) List() []string {
	cs.urls, _ = getURLs(cs.href)
	var links []string
	for _, link := range cs.urls {
		links = append(links, link.String())
	}
	return links
}

func (cs *Sitemap) VisitAll() *colly.Collector {
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
	pp.Println("sitemapURL=", sitemapURL)

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
