package sitemap

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/sniperkit/colly/pkg"
	"github.com/sniperkit/colly/pkg/queue"

	"github.com/sniperkit/xutil/plugin/debug/pp"
)

type Sitemap struct {
	// *Config
	href         url.URL
	sub_sitemaps []url.URL
	urls         []url.URL
	converted    map[string]string
	content      []byte
	contentType  string
	contentSize  int
	statusCode   int
	prefixPath   string
	localPath    string
	total_size   int
	startTime    time.Time
	endTime      time.Time
	responseTime float64
	collector    *colly.Collector
	cqueue       *queue.Queue
	lock         *sync.RWMutex
	wg           *sync.WaitGroup
}

type URL struct {
	Location string `xml:"loc"`
}

type SitemapIndex struct {
	Sitemaps []URL `xml:"sitemap"`
}

type SitemapIndexError struct {
	message string
}

type XMLSitemap struct {
	URLs []URL `xml:"url"`
}

type XmlSitemapError struct {
	message string
}

type TXTSitemap struct {
	URLs []URL `csv:"url"`
}

type TxtSitemapError struct {
	message string
}

func New(inputURL string) (*Sitemap, error) {
	sitemapURL, err := url.Parse(inputURL)
	if err != nil {
		return nil, err
	}

	ouput, err := readURL(*sitemapURL)
	if err != nil {
		return nil, err
	}

	s := &Sitemap{
		href:         *sitemapURL,
		content:      ouput.Body,
		statusCode:   ouput.StatusCode,
		contentType:  ouput.ContentType,
		contentSize:  len(ouput.Body),
		responseTime: ouput.EndTime.Sub(ouput.StartTime).Seconds(),
		// lock:         &sync.RWMutex{},
		// wg:           &sync.WaitGroup{},
	}

	pp.Println(s)

	return s, nil
}

func (s *Sitemap) Read() error {
	return errInvalidContent
}

func (s *Sitemap) Print(format string) error {
	return errInvalidContent
}

func getXMLSitemap(xmlSitemapURL url.URL) (XMLSitemap, error) {
	response, readErr := readURL(xmlSitemapURL)
	if readErr != nil {
		return XMLSitemap{}, readErr
	}

	if !strings.Contains(string(response.GetBody()), "</urlset>") {
		return XMLSitemap{}, XmlSitemapError{"Invalid content"}
	}

	var urlSet XMLSitemap
	unmarshalError := xml.Unmarshal(response.GetBody(), &urlSet)
	if unmarshalError != nil {
		return XMLSitemap{}, unmarshalError
	}
	return urlSet, nil
}

func (sitemapIndexError XmlSitemapError) Error() string {
	return sitemapIndexError.message
}

func isInvalidXMLSitemapContent(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == "Invalid content"
}

func checkSitemap(loc string) bool {
	return true
}

func readURL(url url.URL) (colly.Response, error) {
	startTime := time.Now().UTC()
	resp, fetchErr := http.Get(url.String())
	if fetchErr != nil {
		return colly.Response{}, fetchErr
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return colly.Response{}, readErr
	}

	endTime := time.Now().UTC()

	// content type
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(body)
	}

	return colly.Response{
		Body:        body,
		StatusCode:  resp.StatusCode,
		StartTime:   startTime,
		EndTime:     endTime,
		ContentType: contentType,
	}, nil
}
