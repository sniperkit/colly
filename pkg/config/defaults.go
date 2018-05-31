package config

import (
	"time"
)

const (
	APP_NAME    string = "X-Colly - Web Crawler"
	APP_VERSION string = "1.0.0"
)

const (
	DefaultDetectMimeType         bool   = true
	DefaultDetectCharset          bool   = true
	DefaultParseHTTPErrorResponse bool   = true
	DefaultForceDir               bool   = true
	DefaultForceDirRecursive      bool   = true
	DefaultIgnoreRobotsTxt        bool   = true
	DefaultAllowURLRevisit        bool   = false
	DefaultRandomUserAgent        bool   = false
	DefaultSummarizeContent       bool   = false
	DefaultTopicModelling         bool   = false
	DefaultAnalyzeContent         bool   = false
	DefaultDebugMode              bool   = false
	DefaultVerboseMode            bool   = false
	DefaultMaxDepth               int    = 0
	DefaultMaxBodySize            int    = 10 * 1024 * 1024
	DefaultConfigFilepath         string = "./colly.yaml"
	DefaultSitemapXpath           string = "//urlset/url/loc"
	DefaultSitemapFilename        string = "sitemap.xml"
	DefaultUserAgent              string = "X-Colly - Alpha"
	DefaultStorageDir             string = "./shared/storage"
	DefaultCacheDir               string = "http/raw"
	DefaultEnvPrefix              string = "Colly"
)

var (
	startTime time.Time = time.Now()
	upTime    time.Duration
)

var (
	DefaultStorageCacheDir string   = DefaultStorageDir + "/" + DefaultCacheDir
	DefaultAllowedDomains  []string = []string{
		"golanglibs.com", "github.com", "gitlab.com", "bitbucket.com", "stackoverflow.com", "reddit.com", "medium.com",
	}
	DefaultDisallowedDomains []string = []string{
		"google.com", "bing.com", "yahoo.com",
	}
)
