package sitemap

import (
	"sync"
)

type Sitemap struct {
	Location      string `default:''`
	CacheDir      string `default:''`
	ExportDir     string `default:''`
	IsCache       bool   `default:'true'`
	ExportEntries bool   `default:'false'`
	DryMode       bool   `default:'false'`
	entries       []string
	converted     map[string]string
	prefixPath    string
	localPath     string
	size          int
	isValid       bool
	isIndex       bool // index or sitemap
	isCompressed  bool
	isLocalFile   bool // if false, expecting sitemap as []byte or string
	isDone        bool
	lock          *sync.RWMutex
	wg            *sync.WaitGroup
}

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
