package sitemap

import (
	"encoding/xml"
	"net/url"
	"os"
	"strings"
	"sync"
)

func NewWithConfig(s *Sitemap) (*Sitemap, error) {
	s.lock = &sync.RWMutex{}
	s.wg = &sync.WaitGroup{}
	return s, nil // errInvalidSitemapWithConfig
}

func New(location string, enabled bool, isCache bool, dryMode bool) (*Sitemap, error) {
	isValid := checkSitemap(location)
	if !isValid {
		return nil, errInvalidSitemap
	}

	s := &Sitemap{
		Location:  location,
		IsCache:   isCache,
		DryMode:   dryMode,
		converted: make(map[string]string, 0),
		// entries: make([]string, 0),
		lock: &sync.RWMutex{},
		wg:   &sync.WaitGroup{},
	}
	return s, nil
}

func (s *Sitemap) Read() error {
	return errInvalidContent
}

func (s *Sitemap) Print(format string) error {
	return errInvalidContent
}

func GetXMLSitemap(xmlSitemapURL url.URL) (XMLSitemap, error) {
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

func CreateDirs(dirs []string) (res map[string]bool) {
	res = make(map[string]bool, len(dirs))
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0777); err == nil {
			res[dir] = true
		}
	}
	return
}
