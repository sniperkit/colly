package config

type CollectorConfig struct {

	//////////////////////////////////////////////////
	///// Collector - info
	//////////////////////////////////////////////////

	// ID is the unique identifier of a collector
	Identifier uint32 `default:'colly' json:'identifier' yaml:'identifier' toml:'identifier' xml:'identifier' ini:'identifier'`

	// Title/name of the current crawling campaign
	Title string `default:'Colly - Web Scraper' json:'title' yaml:'title' toml:'title' xml:'title' ini:'title'`

	// UserAgent is the User-Agent string used by HTTP requests
	UserAgent string `default:'colly - https://github.com/sniperkit/colly' json:'user_agent' yaml:'user_agent' toml:'user_agent' xml:'userAgent' ini:'userAgent'`

	// RandomUserAgent specifies to generate a random User-Agent string for all HTTP requests
	RandomUserAgent bool `default:'false' json:'random_user_agent' yaml:'random_user_agent' toml:'random_user_agent' xml:'randomUserAgent' ini:'randomUserAgent'`

	//////////////////////////////////////////////////
	///// Collector - crawling parameters
	//////////////////////////////////////////////////

	// Async turns on asynchronous network communication. Use Collector.Wait() to be sure all requests have been finished.
	Async bool `default:'false' json:'async' yaml:'async' toml:'async' xml:'async' ini:'async'`

	// MaxDepth limits the recursion depth of visited URLs.
	// Set it to 0 for infinite recursion (default).
	MaxDepth int `default:'0' json:'max_depth' yaml:'max_depth' toml:'max_depth' xml:'maxDepth' ini:'maxDepth'`

	// AllowURLRevisit allows multiple downloads of the same URL
	AllowURLRevisit bool `default:'false' json:'allow_url_revisit' yaml:'allow_url_revisit' toml:'allow_url_revisit' xml:'allowURLRevisit' ini:'allowURLRevisit'`

	// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
	// the target host's robots.txt file.  See http://www.robotstxt.org/ for more information.
	IgnoreRobotsTxt bool `default:'true' json:'ignore_robots_txt' yaml:'ignore_robots_txt' toml:'ignore_robots_txt' xml:'ignoreRobotsTxt' ini:'ignoreRobotsTxt'`

	//////////////////////////////////////////////////
	///// Request - Filtering parameters
	//////////////////////////////////////////////////

	////// Not exportable attributes

	// AllowedDomains is a domain whitelist.
	// Leave it blank to allow any domains to be visited
	AllowedDomains []string `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`

	// DisallowedDomains is a domain blacklist.
	DisallowedDomains []string `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`

	// DisallowedURLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request will be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Important: Leave it blank to allow any URLs to be visited
	DisallowedURLFilters []*regexp.Regexp `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`

	// URLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request won't be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Important: Leave it blank to allow any URLs to be visited
	URLFilters []*regexp.Regexp `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`

	////// Exportable attributes

	// WhitelistDomains is a domain whitelist.
	WhitelistDomains []*DomainConfig `json:'whitelist_domains' yaml:'whitelist_domains' toml:'whitelist_domains' xml:'whitelistDomains' ini:'whitelistDomains'`

	// DisallowedDomains is a domain blacklist.
	BlacklistDomains []*DomainConfig `json:'blacklist_domains' yaml:'blacklist_domains' toml:'blacklist_domains' xml:'blacklistDomains' ini:'blacklistDomains'`

	// WhitelistURLFilters is a list...
	WhitelistURLFilters []*FilterConfig `json:'whitelist_url_filters' yaml:'whitelist_url_filters' toml:'whitelist_url_filters' xml:'whitelistURLFilters' ini:'whitelistURLFilters'`

	// BlacklistURLFilters is a list...
	BlacklistURLFilters []*FilterConfig `json:'blacklist_url_filters' yaml:'blacklist_url_filters' toml:'blacklist_url_filters' xml:'blacklistURLFilters' ini:'blacklistURLFilters'`

	//////////////////////////////////////////////////
	///// Request - Filtering parameters
	//////////////////////////////////////////////////

	// WhitelistBodyFilters  is a list...
	WhitelistBodyFilters []*FilterConfig `json:'whitelist_body_filters' yaml:'whitelist_body_filters' toml:'whitelist_body_filters' xml:'whitelistBodyFilters' ini:'whitelistBodyFilters'`

	// BlacklistBodyFilters  is a list...
	BlacklistBodyFilters []*FilterConfig `json:'blacklist_body_filters' yaml:'blacklist_body_filters' toml:'blacklist_body_filters' xml:'blacklistBodyFilters' ini:'blacklistBodyFilters'`

	// WhitelistBodyFilters  is a list...
	WhitelisHeaderFilters []*FilterConfig `json:'whitelist_header_filters' yaml:'whitelist_header_filters' toml:'whitelist_header_filters' xml:'whitelisHeaderFilters' ini:'whitelisHeaderFilters'`

	// BlacklistBodyFilters  is a list...
	BlacklistHeaderFilters []*FilterConfig `json:'blacklist_header_filters' yaml:'blacklist_header_filters' toml:'blacklist_header_filters' xml:'blacklistHeaderFilters' ini:'blacklistHeaderFilters'`

	// MaxBodySize is the limit of the retrieved response body in bytes.
	// 0 means unlimited.
	// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes).
	MaxBodySize int `default:'0' json:'max_body_size' yaml:'max_body_size' toml:'max_body_size' xml:'maxBodySize' ini:'maxBodySize'`

	//////////////////////////////////////////////////
	///// Response processing
	//////////////////////////////////////////////////

	// ParseHTTPErrorResponse allows parsing HTTP responses with non 2xx status codes.
	// By default, Colly parses only successful HTTP responses. Set ParseHTTPErrorResponse to true to enable it.
	ParseHTTPErrorResponse bool `default:'true' json:'parse_http_error_response' yaml:'parse_http_error_response' toml:'parse_http_error_response' xml:'parseHTTPErrorResponse' ini:'parseHTTPErrorResponse'`

	// DetectCharset can enable character encoding detection for non-utf8 response bodies
	// without explicit charset declaration. This feature uses https://github.com/saintfish/chardet
	DetectCharset bool `default:'true' json:'detect_charset' yaml:'detect_charset' toml:'detect_charset' xml:'detectCharset' ini:'detectCharset'`

	// DetectMimeType
	DetectMimeType bool `default:'true' json:'detect_charset' yaml:'detect_charset' toml:'detect_charset' xml:'detectCharset' ini:'detectCharset'`

	// DetectTabular
	DetectTabularData bool `default:'true' json:'detect_tabular_content' yaml:'detect_tabular_content' toml:'detect_tabular_content' xml:'detectTabularContent' ini:'detectTabularContent'`

	//////////////////////////////////////////////////
	///// Filesystem parameters
	//////////////////////////////////////////////////

	// CacheDir specifies a location where GET requests are cached as files.
	// When it's not defined, caching is disabled.
	CacheDir string `default:'./shared/storage/cache/http/backends/internal' json:'cache_dir' yaml:'cache_dir' toml:'cache_dir' xml:'cacheDir' ini:'cacheDir'`

	// ExportDir
	ExportDir string `default:'./shared/exports' json:'export_dir' yaml:'export_dir' toml:'export_dir' xml:'exportDir' ini:'exportDir'`

	// ForceDir specifies that the program will try to create missing storage directories.
	ForceDir bool `default:'true' json:'force_dir' yaml:'force_dir' toml:'force_dir' xml:'forceDir' ini:'forceDir'`

	// ForceDirRecursive specifies that the program will try to create missing storage directories recursively.
	ForceDirRecursive bool `default:'true' json:'force_dir_recursive' yaml:'force_dir_recursive' toml:'force_dir_recursive' xml:'forceDirRecursive' ini:'forceDirRecursive'`

	//////////////////////////////////////////////////
	///// Debug mode
	//////////////////////////////////////////////////

	// DebugMode
	DebugMode bool `default:'false' json:'debug' yaml:'debug' toml:'debug' xml:'debugMode' ini:'debugMode'`

	// VerboseMode
	VerboseMode bool `default:'verbose' json:'verbose' yaml:'verbose' toml:'verbose' xml:'verboseMode' ini:'verboseMode'`

	//////////////////////////////////////////////////
	///// Dashboard TUI (terminal ui only)
	//////////////////////////////////////////////////

	// IsDashboard
	IsDashboard bool `default:'true' json:'dashboard' yaml:'dashboard' toml:'dashboard' xml:'withDashboard' ini:'withDashboard'`

	//////////////////////////////////////////////////
	///// Export application's config to local file
	//////////////////////////////////////////////////

	// AllowExportConfigSchema
	AllowExportConfigSchema bool `default:'true' json:'allow_export_config_schema' yaml:'allow_export_config_schema' toml:'allow_export_config_schema' xml:'allowExportConfigSchema' ini:'allowExportConfigSchema'`

	// AllowExportConfigAll
	AllowExportConfigAutoload bool `default:'false' json:'allow_export_config_autoload' yaml:'allow_export_config_autoload' toml:'allow_export_config_autoload' xml:'allowExportConfigAutoload' ini:'allowExportConfigAutoload'`

	//////////////////////////////////////////////////
	///// Experimental stuff
	//////////////////////////////////////////////////

	// Distributed collector config...
	// Distributed *DistributedConfig `json:'distributed' yaml:'distributed' toml:'distributed' xml:'distributed' ini:'distributed'`

	// Workload stores... not ready yet
	// Workloads map[string]Workload

}
