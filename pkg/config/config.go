package config

import (
	"fmt"
	"os"
	"regexp"
	"time"

	configor "github.com/sniperkit/colly/plugins/data/import/configor"
)

const (
	DEFAULT_BASE_DIR string = "./shared"
)

// private variables
var (
	collectorBaseDir string
	collectorWorkDir string
	collectorAppName string = DEFAULT_APP_NAME
)

// Public variables
var (
	AutoLoad bool = true
)

func init() {
	if AutoLoad {
		autoLoad()
	}
}

type CollectorConfig struct {

	// createdAt is set when...
	createdAt time.Time

	// startedAt is set when...
	startedAt time.Time

	// App
	App struct {
		// ID is the unique identifier of a collector
		ID uint32 `default:"colly" json:"identifier" yaml:"identifier" toml:"identifier" xml:"identifier" ini:"identifier" csv:"identifier"`

		// Title/name of the current crawling campaign
		Title string `default:"Colly - Web Scraper" json:"title" yaml:"title" toml:"title" xml:"title" ini:"title" csv:"title"`

		// DebugMode
		DebugMode bool `default:"false" json:"debug" yaml:"debug" toml:"debug" xml:"debugMode" ini:"debugMode" csv:"DebugMode"`

		// VerboseMode
		VerboseMode bool `default:"false" json:"verbose" yaml:"verbose" toml:"verbose" xml:"verboseMode" ini:"verboseMode" csv:"VerboseMode"`

		//////////////////////////////////////////////////
		///// Dashboard TUI (terminal ui only)
		//////////////////////////////////////////////////

		// IsDashboard
		DashboardMode bool `default:"true" json:"dashboard" yaml:"dashboard" toml:"dashboard" xml:"dashboard" ini:"dashboardMode" csv:"dashboardMode"`
	} `json:"app" yaml:"app" toml:"app" xml:"app" ini:"app" csv:"App"`

	Crawler struct {

		// UserAgent is the User-Agent string used by HTTP requests
		UserAgent string `default:"colly - https://github.com/sniperkit/colly" json:"user_agent" yaml:"user_agent" toml:"user_agent" xml:"userAgent" ini:"userAgent" csv:"userAgent"`

		// RandomUserAgent specifies to generate a random User-Agent string for all HTTP requests
		RandomUserAgent bool `default:"false" json:"random_user_agent" yaml:"random_user_agent" toml:"random_user_agent" xml:"randomUserAgent" ini:"randomUserAgent" csv:"randomUserAgent"`

		// MaxDepth limits the recursion depth of visited URLs.
		// Set it to 0 for infinite recursion (default).
		MaxDepth int `default:"0" json:"max_depth" yaml:"max_depth" toml:"max_depth" xml:"maxDepth" ini:"maxDepth" csv:"maxDepth"`

		// AllowURLRevisit allows multiple downloads of the same URL
		AllowURLRevisit bool `default:"false" json:"allow_url_revisit" yaml:"allow_url_revisit" toml:"allow_url_revisit" xml:"allowURLRevisit" ini:"allowURLRevisit" csv:"allowURLRevisit"`

		// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
		// the target host"s robots.txt file.  See http://www.robotstxt.org/ for more information.
		IgnoreRobotsTxt bool `default:"true" json:"ignore_robots_txt" yaml:"ignore_robots_txt" toml:"ignore_robots_txt" xml:"ignoreRobotsTxt" ini:"ignoreRobotsTxt" csv:"ignoreRobotsTxt"`

		// CurrentMode
		CurrentMode string `default:"async" json:"current_mode" yaml:"current_mode" toml:"current_mode" xml:"CurrentMode" ini:"CurrentMode" csv:"CurrentMode"`

		// Modes
		Modes struct {
			Default struct {
				RandomDelay time.Duration `default:"5" json:"random_delay" yaml:"random_delay" toml:"random_delay" xml:"randomDelay" ini:"randomDelay" csv:"RandomDelay"`
			} `json:"default" yaml:"default" toml:"default" xml:"default" ini:"default" csv:"default"`
			// Async
			Async struct {
				Parallelism int           `default:"2" json:"parallelism" yaml:"parallelism" toml:"parallelism" xml:"parallelism" ini:"parallelism" csv:"Parallelism"`
				DomainGlob  string        `default:"*" json:"domain_glob" yaml:"domain_glob" toml:"domain_glob" xml:"domainGlob" ini:"domainGlob" csv:"DomainGlob"`
				RandomDelay time.Duration `default:"5" json:"random_delay" yaml:"random_delay" toml:"random_delay" xml:"randomDelay" ini:"randomDelay" csv:"RandomDelay"`
				MaxQueue    int           `default:"10000" json:"max_queue" yaml:"max_queue" toml:"max_queue" xml:"maxQueue" ini:"maxQueue" csv:"MaxQueue"`
			} `json:"async" yaml:"async" toml:"async" xml:"async" ini:"async" csv:"async"`
			// Queue
			Queue struct {
				Workers     int           `default:"2" json:"workers" yaml:"workers" toml:"workers" xml:"workers" ini:"workers" csv:"Workers"`
				MaxQueue    int           `default:"10000" json:"max_queue" yaml:"max_queue" toml:"max_queue" xml:"maxQueue" ini:"maxQueue" csv:"MaxQueue"`
				RandomDelay time.Duration `default:"5" json:"random_delay" yaml:"random_delay" toml:"random_delay" xml:"randomDelay" ini:"randomDelay" csv:"RandomDelay"`
			} `json:"queue" yaml:"queue" toml:"queue" xml:"queue" ini:"queue" csv:"queue"`
		} `json:"modes" yaml:"modes" toml:"modes" xml:"modes" ini:"modes" csv:"modes"`

		// Cache
		Cache struct {
			// Enabled          bool     `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`
			Disabled bool   `default:"true" json:"disabled" yaml:"disabled" toml:"disabled" xml:"disabled" ini:"disabled" csv:"disabled"`
			Backend  string `default:"inMemory" json:"backend" yaml:"backend" toml:"backend" xml:"backend" ini:"backend" csv:"backend"`
			// CacheDir specifies a location where GET requests are cached as files.  When it"s not defined, caching is disabled.
			CustomDir string `default:"./shared/storage/cache/http/backends/internal" json:"custom_dir" yaml:"custom_dir" toml:"custom_dir" xml:"custom_dir" ini:"custom_dir" csv:"custom_dir"`
		} `json:"cache" yaml:"cache" toml:"cache" xml:"cache" ini:"cache" csv:"cache"`

		// Transport
		Transport struct {
			Disabled bool `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`
			Http     struct {
				Cache struct {
					Disabled bool           `default:"true" json:"disabled" yaml:"disabled" toml:"disabled" xml:"disabled" ini:"disabled" csv:"disabled"`
					Backend  string         `default:"badger" json:"backend" yaml:"backend" toml:"backend" xml:"backend" ini:"backend" csv:"backend"`
					TTL      string         `default:"ttl" json:"ttl" yaml:"ttl" toml:"ttl" xml:"ttl" ini:"ttl" csv:"TTL"`
					Store    []*StoreConfig `json:"store" yaml:"store" toml:"store" xml:"store" ini:"store" csv:"store"`
				} `json:"cache" yaml:"cache" toml:"cache" xml:"cache" ini:"cache" csv:"cache"`
				Stats struct {
					Disabled bool           `default:"true" json:"disabled" yaml:"disabled" toml:"disabled" xml:"disabled" ini:"disabled" csv:"disabled"`
					Client   string         `default:"" json:"client" yaml:"client" toml:"client" xml:"client" ini:"client" csv:"client"`
					Store    []*StoreConfig `json:"store" yaml:"store" toml:"store" xml:"store" ini:"store" csv:"store"`
				} `json:"stats" yaml:"stats" toml:"stats" xml:"stats" ini:"stats" csv:"stats"`
			} `json:"http" yaml:"http" toml:"http" xml:"http" ini:"http" csv:"http"`
		} `json:"transport" yaml:"transport" toml:"transport" xml:"transport" ini:"transport" csv:"transport"`

		// Proxy
		Proxy struct {
			// Enabled          bool     `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`
			Disabled         bool     `default:"true" json:"disabled" yaml:"disabled" toml:"disabled" xml:"disabled" ini:"disabled" csv:"disabled"`
			FetchRemoteList  bool     `default:"true" json:"fetch_remote_list" yaml:"fetch_remote_list" toml:"fetch_remote_list" xml:"fetchRemoteList" ini:"fetchRemoteList" csv:"FetchRemoteList"`
			PoolMode         bool     `default:"true" json:"pool_mode" yaml:"pool_mode" toml:"pool_mode" xml:"poolMode" ini:"poolMode" csv:"PoolMode"`
			AllowedProtocols []string `json:"allowed_protocols" yaml:"allowed_protocols" toml:"allowed_protocols" xml:"allowedProtocols" ini:"allowedProtocols" csv:"AllowedProtocols"`
			List             []string `json:"list" yaml:"list" toml:"list" xml:"list" ini:"list" csv:"list"`
		} `json:"proxy" yaml:"proxy" toml:"proxy" xml:"proxy" ini:"proxy" csv:"proxy"`

		// Sitemap
		Sitemap struct {
			// Enabled          bool     `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`
			Disabled   bool   `default:"true" json:"disabled" yaml:"disabled" toml:"disabled" xml:"disabled" ini:"disabled" csv:"disabled"`
			URL        string `json:"url" yaml:"url" toml:"url" xml:"url" ini:"URL" csv:"URL"`
			AutoDetect bool   `default:"false" json:"auto_detect" yaml:"auto_detect" toml:"auto_detect" xml:"autoDetect" ini:"autoDetect" csv:"AutoDetect"`
			LimitURLs  int    `default:"0" json:"limit_urls" yaml:"limit_urls" toml:"limit_urls" xml:"limitURLs" ini:"limitURLs" csv:"limitURLs"`
		} `json:"sitemap" yaml:"sitemap" toml:"sitemap" xml:"sitemap" ini:"sitemap" csv:"sitemap"`
	} `json:"crawler" yaml:"crawler" toml:"crawler" xml:"crawler" ini:"crawler" csv:"crawler"`

	// Filters
	Filters struct {
		Response struct {
			// ParseHTTPErrorResponse allows parsing HTTP responses with non 2xx status codes.
			// By default, Colly parses only successful HTTP responses. Set ParseHTTPErrorResponse to true to enable it.
			ParseHTTPErrorResponse bool `default:"true" json:"parse_http_error_response" yaml:"parse_http_error_response" toml:"parse_http_error_response" xml:"parseHTTPErrorResponse" ini:"parseHTTPErrorResponse" csv:"parseHTTPErrorResponse"`

			// DetectCharset can enable character encoding detection for non-utf8 response bodies
			// without explicit charset declaration. This feature uses https://github.com/saintfish/chardet
			DetectCharset bool `default:"true" json:"detect_charset" yaml:"detect_charset" toml:"detect_charset" xml:"detectCharset" ini:"detectCharset" csv:"DetectCharset"`

			// DetectMimeType
			DetectMimeType bool `default:"true" json:"detect_mime_type" yaml:"detect_mime_type" toml:"detect_mime_type" xml:"detectMimeType" ini:"detectMimeType" csv:"detectMimeType"`

			// DetectTabular
			DetectTabular bool `default:"true" json:"detect_tabular_data" yaml:"detect_tabular_data" toml:"detect_tabular_data" xml:"detectTabularData" ini:"detectTabularData" csv:"DetectTabularData"`
			// MaxBodySize is the limit of the retrieved response body in bytes.
			// 0 means unlimited.
			// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes).
			MaxBodySize int `default:"0" json:"max_body_size" yaml:"max_body_size" toml:"max_body_size" xml:"maxBodySize" ini:"maxBodySize" csv:"maxBodySize"`
		} `json:"response" yaml:"response" toml:"response" xml:"response" ini:"response" csv:"Response"`

		Blacklists struct {
			Domains    []*DomainConfig `json:"domains" yaml:"domains" toml:"domains" xml:"domains" ini:"domains" csv:"Domains"`
			URLs       []*FilterConfig `json:"urls" yaml:"urls" toml:"urls" xml:"urls" ini:"urls" csv:"urls"`
			Extensions []string        `json:"extensions" yaml:"extensions" toml:"extensions" xml:"extensions" ini:"extensions" csv:"Extensions"`
			Headers    []*FilterConfig `json:"headers" yaml:"headers" toml:"headers" xml:"headers" ini:"headers" csv:"headers"`
			MimeTypes  []string        `json:"mime_types" yaml:"mime_types" toml:"mime_types" xml:"mimeTypes" ini:"mimeTypes" csv:"MimeTypes"`
			Responses  []*FilterConfig `json:"responses" yaml:"responses" toml:"responses" xml:"responses" ini:"responses" csv:"responses"`
		} `json:"blacklists" yaml:"blacklists" toml:"blacklists" xml:"blackLists" ini:"blackLists" csv:"BlackLists"`

		Whitelists struct {
			Domains    []*DomainConfig `json:"domains" yaml:"domains" toml:"domains" xml:"domains" ini:"domains" csv:"Domains"`
			URLs       []*FilterConfig `json:"urls" yaml:"urls" toml:"urls" xml:"urls" ini:"urls" csv:"urls"`
			Extensions []string        `json:"extensions" yaml:"extensions" toml:"extensions" xml:"extensions" ini:"extensions" csv:"Extensions"`
			Headers    []*FilterConfig `json:"headers" yaml:"headers" toml:"headers" xml:"headers" ini:"headers" csv:"headers"`
			MimeTypes  []string        `json:"mime_types" yaml:"mime_types" toml:"mime_types" xml:"mimeTypes" ini:"mimeTypes" csv:"MimeTypes"`
			Responses  []*FilterConfig `json:"responses" yaml:"responses" toml:"responses" xml:"responses" ini:"responses" csv:"responses"`
		} `json:"whitelists" yaml:"whitelists" toml:"whitelists" xml:"whiteLists" ini:"whiteLists" csv:"Whitelists"`
	} `json:"filters" yaml:"filters" toml:"filters" xml:"filters" ini:"filters" csv:"filters"`

	Collection struct {
		Disabled  bool              `default:"true" json:"disabled" yaml:"disabled" toml:"disabled" xml:"disabled" ini:"disabled" csv:"disabled"`
		Databooks []*DatabookConfig `json:"databooks" yaml:"databooks" toml:"databooks" xml:"databooks" ini:"databooks" csv:"databooks"`
		Datasets  []*DatasetConfig  `json:"datasets" yaml:"datasets" toml:"datasets" xml:"datasets" ini:"datasets" csv:"datasets"`
	} `json:"collection" yaml:"collection" toml:"collection" xml:"collection" ini:"collection" csv:"collection"`

	// Stores struct {} `json:"stores" yaml:"stores" toml:"stores" xml:"stores" ini:"stores" csv:"stores"`

	//////////////////////////////////////////////////
	///// Outputs parameters
	//////////////////////////////////////////////////

	Dirs struct {
		// XDGBaseDir
		XDGBaseDir string `json:"xdg_base_dir" yaml:"xdg_base_dir" toml:"xdg_base_dir" xml:"xdgBaseDir" ini:"xdgBaseDir" csv:"XDGBaseDir"`

		// BaseDirectory
		BaseDir string `json:"base_dir" yaml:"base_dir" toml:"base_dir" xml:"baseDir" ini:"baseDir" csv:"BaseDir"`

		// LogsDirectory
		LogsDir string `json:"logs_dir" yaml:"logs_dir" toml:"logs_dir" xml:"logsDir" ini:"logsDir" csv:"LogsDir"`

		// CacheDir specifies a location where GET requests are cached as files.
		// When it"s not defined, caching is disabled.
		CacheDir string `default:"./shared/storage/cache/http/backends/internal" json:"cache_dir" yaml:"cache_dir" toml:"cache_dir" xml:"cacheDir" ini:"cacheDir" csv:"CacheDir"`

		// ExportDir
		ExportDir string `default:"./shared/exports" json:"export_dir" yaml:"export_dir" toml:"export_dir" xml:"exportDir" ini:"exportDir" csv:"ExportDir"`

		// ForceDir specifies that the program will try to create missing storage directories.
		ForceDir bool `default:"true" json:"force_dir" yaml:"force_dir" toml:"force_dir" xml:"forceDir" ini:"forceDir" csv:"ForceDir"`

		// ForceDirRecursive specifies that the program will try to create missing storage directories recursively.
		ForceDirRecursive bool `default:"true" json:"force_dir_recursive" yaml:"force_dir_recursive" toml:"force_dir_recursive" xml:"forceDirRecursive" ini:"forceDirRecursive" csv:"ForceDirRecursive"`
	} `json:"outputs" yaml:"outputs" toml:"outputs" xml:"outputs" ini:"outputs" csv:"outputs"`

	//////////////////////////////////////////////////
	///// Collector - info
	//////////////////////////////////////////////////

	// ID is the unique identifier of a collector
	ID uint32 `default:"colly" json:"identifier" yaml:"identifier" toml:"identifier" xml:"identifier" ini:"identifier" csv:"identifier"`

	// Title/name of the current crawling campaign
	Title string `default:"Colly - Web Scraper" json:"title" yaml:"title" toml:"title" xml:"title" ini:"title" csv:"title"`

	// UserAgent is the User-Agent string used by HTTP requests
	UserAgent string `default:"colly - https://github.com/sniperkit/colly" json:"user_agent" yaml:"user_agent" toml:"user_agent" xml:"userAgent" ini:"userAgent" csv:"userAgent"`

	// RandomUserAgent specifies to generate a random User-Agent string for all HTTP requests
	RandomUserAgent bool `default:"false" json:"random_user_agent" yaml:"random_user_agent" toml:"random_user_agent" xml:"randomUserAgent" ini:"randomUserAgent" csv:"randomUserAgent"`

	//////////////////////////////////////////////////
	///// Collector - crawling parameters
	//////////////////////////////////////////////////

	// Async turns on asynchronous network communication. Use Collector.Wait() to be sure all requests have been finished.
	Async bool `default:"false" json:"async" yaml:"async" toml:"async" xml:"async" ini:"async" csv:"async"`

	// MaxDepth limits the recursion depth of visited URLs.
	// Set it to 0 for infinite recursion (default).
	MaxDepth int `default:"0" json:"max_depth" yaml:"max_depth" toml:"max_depth" xml:"maxDepth" ini:"maxDepth" csv:"maxDepth"`

	// AllowURLRevisit allows multiple downloads of the same URL
	AllowURLRevisit bool `default:"false" json:"allow_url_revisit" yaml:"allow_url_revisit" toml:"allow_url_revisit" xml:"allowURLRevisit" ini:"allowURLRevisit" csv:"allowURLRevisit"`

	// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
	// the target host"s robots.txt file.  See http://www.robotstxt.org/ for more information.
	IgnoreRobotsTxt bool `default:"true" json:"ignore_robots_txt" yaml:"ignore_robots_txt" toml:"ignore_robots_txt" xml:"ignoreRobotsTxt" ini:"ignoreRobotsTxt" csv:"ignoreRobotsTxt"`

	//////////////////////////////////////////////////
	///// Request - Filtering parameters
	//////////////////////////////////////////////////

	////// Not exportable attributes

	// AllowedDomains is a domain whitelist.
	// Leave it blank to allow any domains to be visited
	AllowedDomains []string `json:"allowed_domains" yaml:"allowed_domains" toml:"allowed_domains" xml:"allowedDomains" ini:"allowedDomains" csv:"AllowedDomains"`

	// DisallowedDomains is a domain blacklist.
	DisallowedDomains []string `json:"disallowed_domains" yaml:"disallowed_domains" toml:"disallowed_domains" xml:"disallowedDomains" ini:"disallowedDomains" csv:"DisallowedDomains"`

	// DisallowedURLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request will be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Important: Leave it blank to allow any URLs to be visited
	DisallowedURLFilters []*regexp.Regexp `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// URLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request won"t be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Important: Leave it blank to allow any URLs to be visited
	URLFilters []*regexp.Regexp `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	////// Exportable attributes

	// WhitelistDomains is a domain whitelist.
	// WhitelistDomains []*DomainConfig `json:"whitelist_domains" yaml:"whitelist_domains" toml:"whitelist_domains" xml:"whitelistDomains" ini:"whitelistDomains" csv:"whitelistDomains"`

	// DisallowedDomains is a domain blacklist.
	// BlacklistDomains []*DomainConfig `json:"blacklist_domains" yaml:"blacklist_domains" toml:"blacklist_domains" xml:"blacklistDomains" ini:"blacklistDomains" csv:"blacklistDomains"`

	// WhitelistURLFilters is a list...
	// WhitelistURLFilters []*FilterConfig `json:"whitelist_url_filters" yaml:"whitelist_url_filters" toml:"whitelist_url_filters" xml:"whitelistURLFilters" ini:"whitelistURLFilters" csv:"whitelistURLFilters"`

	// BlacklistURLFilters is a list...
	// BlacklistURLFilters []*FilterConfig `json:"blacklist_url_filters" yaml:"blacklist_url_filters" toml:"blacklist_url_filters" xml:"blacklistURLFilters" ini:"blacklistURLFilters" csv:"blacklistURLFilters"`

	//////////////////////////////////////////////////
	///// Request - Filtering parameters
	//////////////////////////////////////////////////

	// WhitelistBodyFilters  is a list...
	// WhitelistBodyFilters []*FilterConfig `json:"whitelist_body_filters" yaml:"whitelist_body_filters" toml:"whitelist_body_filters" xml:"whitelistBodyFilters" ini:"whitelistBodyFilters" csv:"whitelistBodyFilters"`

	// BlacklistBodyFilters  is a list...
	// BlacklistBodyFilters []*FilterConfig `json:"blacklist_body_filters" yaml:"blacklist_body_filters" toml:"blacklist_body_filters" xml:"blacklistBodyFilters" ini:"blacklistBodyFilters" csv:"blacklistBodyFilters"`

	// WhitelistBodyFilters  is a list...
	// WhitelisHeaderFilters []*FilterConfig `json:"whitelist_header_filters" yaml:"whitelist_header_filters" toml:"whitelist_header_filters" xml:"whitelisHeaderFilters" ini:"whitelisHeaderFilters" csv:"whitelisHeaderFilters"`

	// BlacklistBodyFilters  is a list...
	// BlacklistHeaderFilters []*FilterConfig `json:"blacklist_header_filters" yaml:"blacklist_header_filters" toml:"blacklist_header_filters" xml:"blacklistHeaderFilters" ini:"blacklistHeaderFilters" csv:"blacklistHeaderFilters"`

	// MaxBodySize is the limit of the retrieved response body in bytes.
	// 0 means unlimited.
	// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes).
	MaxBodySize int `default:"0" json:"max_body_size" yaml:"max_body_size" toml:"max_body_size" xml:"maxBodySize" ini:"maxBodySize" csv:"maxBodySize"`

	//////////////////////////////////////////////////
	///// Response processing
	//////////////////////////////////////////////////

	// ParseHTTPErrorResponse allows parsing HTTP responses with non 2xx status codes.
	// By default, Colly parses only successful HTTP responses. Set ParseHTTPErrorResponse to true to enable it.
	ParseHTTPErrorResponse bool `default:"true" json:"parse_http_error_response" yaml:"parse_http_error_response" toml:"parse_http_error_response" xml:"parseHTTPErrorResponse" ini:"parseHTTPErrorResponse" csv:"parseHTTPErrorResponse"`

	// DetectCharset can enable character encoding detection for non-utf8 response bodies
	// without explicit charset declaration. This feature uses https://github.com/saintfish/chardet
	DetectCharset bool `default:"true" json:"detect_charset" yaml:"detect_charset" toml:"detect_charset" xml:"detectCharset" ini:"detectCharset" csv:"DetectCharset"`

	// DetectMimeType
	DetectMimeType bool `default:"true" json:"detect_mime_type" yaml:"detect_mime_type" toml:"detect_mime_type" xml:"detectMimeType" ini:"detectMimeType" csv:"detectMimeType"`

	// DetectTabular
	DetectTabular bool `default:"true" json:"detect_tabular_data" yaml:"detect_tabular_data" toml:"detect_tabular_data" xml:"detectTabularData" ini:"detectTabularData" csv:"DetectTabularData"`

	// XDGBaseDir
	XDGBaseDir string `json:"xdg_base_dir" yaml:"xdg_base_dir" toml:"xdg_base_dir" xml:"xdgBaseDir" ini:"xdgBaseDir" csv:"XDGBaseDir"`

	// BaseDirectory
	BaseDir string `json:"base_dir" yaml:"base_dir" toml:"base_dir" xml:"baseDir" ini:"baseDir" csv:"BaseDir"`

	// LogsDirectory
	LogsDir string `json:"logs_dir" yaml:"logs_dir" toml:"logs_dir" xml:"logsDir" ini:"logsDir" csv:"LogsDir"`

	// CacheDir specifies a location where GET requests are cached as files.
	// When it"s not defined, caching is disabled.
	CacheDir string `default:"./shared/storage/cache/http/backends/internal" json:"cache_dir" yaml:"cache_dir" toml:"cache_dir" xml:"cacheDir" ini:"cacheDir" csv:"CacheDir"`

	// ExportDir
	ExportDir string `default:"./shared/exports" json:"export_dir" yaml:"export_dir" toml:"export_dir" xml:"exportDir" ini:"exportDir" csv:"ExportDir"`

	// ForceDir specifies that the program will try to create missing storage directories.
	ForceDir bool `default:"true" json:"force_dir" yaml:"force_dir" toml:"force_dir" xml:"forceDir" ini:"forceDir" csv:"ForceDir"`

	// ForceDirRecursive specifies that the program will try to create missing storage directories recursively.
	ForceDirRecursive bool `default:"true" json:"force_dir_recursive" yaml:"force_dir_recursive" toml:"force_dir_recursive" xml:"forceDirRecursive" ini:"forceDirRecursive" csv:"ForceDirRecursive"`

	//////////////////////////////////////////////////
	///// Debug mode
	//////////////////////////////////////////////////

	// DebugMode
	DebugMode bool `default:"false" json:"debug" yaml:"debug" toml:"debug" xml:"debugMode" ini:"debugMode" csv:"DebugMode"`

	// VerboseMode
	VerboseMode bool `default:"verbose" json:"verbose" yaml:"verbose" toml:"verbose" xml:"verboseMode" ini:"verboseMode" csv:"VerboseMode"`

	//////////////////////////////////////////////////
	///// Dashboard TUI (terminal ui only)
	//////////////////////////////////////////////////

	// IsDashboard
	DashboardMode bool `default:"true" json:"dashboard_mode" yaml:"dashboard_mode" toml:"dashboard_mode" xml:"dashboardMode" ini:"dashboardMode" csv:"dashboardMode"`

	//////////////////////////////////////////////////
	///// Export application"s config to local file
	//////////////////////////////////////////////////

	// AllowExportConfigSchema
	AllowExportConfigSchema bool `default:"true" json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// AllowExportConfigAll
	AllowExportConfigAutoload bool `default:"false" json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	//////////////////////////////////////////////////
	///// Experimental stuff
	//////////////////////////////////////////////////

}

type LimitConfig struct {
	Parallelism int           `default:"2" json:"parallelism" yaml:"parallelism" toml:"parallelism" xml:"parallelism" ini:"parallelism" csv:"Parallelism"`
	DomainGlob  string        `default:"*" json:"domain_glob" yaml:"domain_glob" toml:"domain_glob" xml:"domainGlob" ini:"domainGlob" csv:"DomainGlob"`
	RandomDelay time.Duration `json:"random_delay" yaml:"random_delay" toml:"random_delay" xml:"randomDelay" ini:"randomDelay" csv:"RandomDelay"`
	MaxQueue    int           `default:"10000" json:"max_queue" yaml:"max_queue" toml:"max_queue" xml:"maxQueue" ini:"maxQueue" csv:"MaxQueue"`
}

//	configor.New(&configor.Config{Debug: true, Verbose: true}).Load(&Config, "config.json")

func NewFromFile(debug, verbose bool, files ...string) (*CollectorConfig, error) {
	collectorConfig := &CollectorConfig{}
	xdgPath, err := getDefaultXDGBaseDirectory()
	if err != nil {
		return nil, err
	}
	collectorConfig.XDGBaseDir = xdgPath
	// configor.Load(&collectorConfig, files...)
	configor.New(&configor.Config{Debug: debug, Verbose: verbose}).Load(&collectorConfig, files...)

	// if c.DebugMode {}
	// fmt.Printf("config: %#v", collectorConfig)
	// fmt.Println("XDGBaseDir=", xdgPath)

	return collectorConfig, nil
}

func (c *CollectorConfig) Dump(formats, nodes []string, prefixPath string) error {
	return configor.Dump(c, nodes, formats, prefixPath)
}

func Dump(c interface{}, formats, nodes []string, prefixPath string) error {
	return configor.Dump(c, nodes, formats, prefixPath)
}

func autoLoad() {
	var err error
	collectorBaseDir, err = configor.XDGBaseDir()
	if err != nil {
		fmt.Println("Can't find XDG BaseDirectory")
		os.Exit(1)
	}
}

func getDefaultXDGBaseDirectory() (string, error) {
	xdgPath, err := configor.XDGBaseDir()
	if err != nil {
		return DEFAULT_BASE_DIR, err
	}
	return xdgPath, nil
}
