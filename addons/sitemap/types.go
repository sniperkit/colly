package sitemap

import (
	"sync"
)

type Config struct {
	Location      string `default:''`
	CacheDir      string `default:'./shared/storage/cache/sitemaps'`
	ExportDir     string `default:'./shared/storage/export/sitemaps'`
	IsCache       bool   `default:'true'`
	ExportEntries bool   `default:'false'`
	DryMode       bool   `default:'false'`
	EnsureDirs    bool   `default:'true'`
}

type Sitemap struct {
	// *Config
	Location      string `default:''`
	CacheDir      string `default:'./shared/storage/cache/sitemaps'`
	ExportDir     string `default:'./shared/storage/export/sitemaps'`
	IsCache       bool   `default:'true'`
	ExportEntries bool   `default:'false'`
	DryMode       bool   `default:'false'`
	EnsureDirs    bool   `default:'true'`
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
