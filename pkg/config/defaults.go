package config

const (
	DEFAULT_APP_NAME    string = "X-Colly - Web Crawler"
	DEFAULT_APP_VERSION string = "1.0.0"
)

var (
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
