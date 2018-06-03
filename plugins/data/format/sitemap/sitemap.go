package sitemap

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	// "encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/sniperkit/colly/pkg"
	"github.com/sniperkit/colly/pkg/queue"
	// "github.com/sniperkit/xutil/plugin/debug/pp"
)

type Sitemap struct {
	Name         xml.Name `xml:"urlset,sitemapindex"`
	NS           string   `xml:"xmlns,attr"`
	Indices      []url.URL
	URLs         []url.URL
	href         url.URL
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
	// Indices []Index `json:"loc" xml:"sitemap"`
	// URLs    []URL   `json:"url" xml:"url"`
}

type URL struct {
	href       url.URL
	Loc        string `json:"loc" xml:"loc"`
	LastMod    string `json:"lastmod" xml:"lastmod"`
	ChangeFreq string `json:"changefreq" xml:"changefreq"`
	Priority   string `json:"priority" xml:"priority"`
}

type Index struct {
	Loc     string `json:"loc" xml:"loc"`
	LastMod string `json:"lastmod" xml:"lastmod"`
}

type Indices struct {
	Sitemaps []Index `xml:"sitemap" json:"sitemap"`
}

type IndexError struct {
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

	s.getURLs()
	// pp.Println(s)
	return s, nil
}

func (s *Sitemap) IsValid() bool {
	return s.href.String() != ""
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

func getTXTSitemap(txtSitemapURL url.URL) (TXTSitemap, error) {
	_, readErr := readURL(txtSitemapURL)
	if readErr != nil {
		return TXTSitemap{}, readErr
	}

	var urlSet TXTSitemap
	//unmarshalError := csv.Unmarshal(response.GetBody(), &urlSet)
	//if unmarshalError != nil {
	//	return TXTSitemap{}, unmarshalError
	//}
	return urlSet, nil
}

func (sitemapIndexError TxtSitemapError) Error() string {
	return sitemapIndexError.message
}

func isInvalidSitemapContent(err error) bool {
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

	var body []byte
	var errReader error

	// content type
	contentType := resp.Header.Get("Content-Type")

	body, errReader = ioutil.ReadAll(resp.Body)
	if errReader != nil {
		log.Fatalln("error.ReadAll:", contentType, ", msg=", errReader)
		return colly.Response{}, errReader
	}

	//if contentType == "" {
	contentType = http.DetectContentType(body)
	//}

	pathExtension := path.Ext(url.String())
	contentEncoding := resp.Header.Get("Content-Encoding")

	log.Println("URL:", resp.Request.URL.String(), "Content-Type:", contentType, "StatusCode:", resp.StatusCode)

	if resp.StatusCode == 404 {
		return colly.Response{}, errReader
	}

	switch pathExtension {
	case ".gz":
		contentType = "application/x-gzip"
	case ".txt":
		contentType = "text/plain"
	}

	switch contentEncoding {
	case "gzip":
		contentType = "application/x-gzip"
	case "deflate":
		contentType = "application/x-deflate"
	case "zlib":
		contentType = "application/x-zlib"
	}

	switch contentType {
	// "application/octet-stream", "application/x-tar"
	case "application/x-gzip", "application/gzip":
		gr, _ := gzip.NewReader(bytes.NewBuffer(body))
		defer gr.Close()

		body, errReader = ioutil.ReadAll(gr)

	case "application/x-deflate", "application/deflate":
		rdata := flate.NewReader(bytes.NewBuffer(body))
		body, errReader = ioutil.ReadAll(rdata)

	case "application/x-zlib", "application/zlib":
		var readCloser io.ReadCloser
		readCloser, errReader = zlib.NewReader(bytes.NewBuffer(body))
		if errReader != nil {
			log.Fatalln("readCloser.error:", contentType, ", msg=", errReader)
			return colly.Response{}, errReader
		}
		body, errReader = ioutil.ReadAll(readCloser)

	}

	if errReader != nil {
		log.Fatalln("error.ReadAll:", contentType, ", msg=", errReader)
		return colly.Response{}, errReader
	}

	endTime := time.Now().UTC()

	return colly.Response{
		Body:        body,
		StatusCode:  resp.StatusCode,
		StartTime:   startTime,
		EndTime:     endTime,
		ContentType: contentType,
	}, nil
}

// String return the string format of the sitemap
func (s *Sitemap) String() string {
	var items []string
	for _, item := range s.URLs {
		items = append(items, item.String())
	}
	return fmt.Sprintf(SitemapXML, strings.Join(items, `
`))
}

// ToFile saves a sitemap to a file with either extension .xml or .gz.
// If extension is .gz, the file will be gzipped.
func (s *Sitemap) ToFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := filepath.Ext(file.Name())
	if ext != ".xml" && ext != ".gz" {
		return fmt.Errorf("filename %s does not have extension .xml or .gz, extension %s given", file.Name(), ext)
	}

	// Gzip
	if ext == ".gz" {
		zip := gzip.NewWriter(file)
		defer zip.Close()

		_, err = zip.Write([]byte(s.String()))
		if err != nil {
			return err
		}
	} else {
		_, err = file.Write([]byte(s.String()))
		if err != nil {
			return err
		}
	}

	return nil
}
