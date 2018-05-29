package config

import (
	"regexp"
)

type CollectorConfig struct {
	// UserAgent is the User-Agent string used by HTTP requests
	UserAgent string `default:'colly - https://github.com/sniperkit/colly'`

	// RandomUserAgent specifies to generate a random User-Agent string for all HTTP requests
	RandomUserAgent bool `default:'false'`

	// MaxDepth limits the recursion depth of visited URLs.
	// Set it to 0 for infinite recursion (default).
	MaxDepth int `default:'0'`

	// AllowedDomains is a domain whitelist.
	// Leave it blank to allow any domains to be visited
	AllowedDomains []string

	// DisallowedDomains is a domain blacklist.
	DisallowedDomains []string

	// DisallowedURLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request will be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Important: Leave it blank to allow any URLs to be visited
	DisallowedURLFilters []*regexp.Regexp

	// URLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request won't be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Important: Leave it blank to allow any URLs to be visited
	URLFilters []*regexp.Regexp

	// AllowURLRevisit allows multiple downloads of the same URL
	AllowURLRevisit bool `default:'false'`

	// MaxBodySize is the limit of the retrieved response body in bytes.
	// 0 means unlimited.
	// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes).
	MaxBodySize int `default:'10485760'`

	// CacheDir specifies a location where GET requests are cached as files.
	// When it's not defined, caching is disabled.
	CacheDir string `default:'./shared/storage/http/raw'`

	// ForceDir specifies that the program will try to create missing storage directories.
	ForceDir bool `default:'true'`

	// ForceDirRecursive specifies that the program will try to create missing storage directories recursively.
	ForceDirRecursive bool `default:'true'`

	// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
	// the target host's robots.txt file.  See http://www.robotstxt.org/ for more information.
	IgnoreRobotsTxt bool `default:'true'`

	// Async turns on asynchronous network communication. Use Collector.Wait() to be sure all requests have been finished.
	Async bool `default:'false'`

	// ParseHTTPErrorResponse allows parsing HTTP responses with non 2xx status codes.
	// By default, Colly parses only successful HTTP responses. Set ParseHTTPErrorResponse to true to enable it.
	ParseHTTPErrorResponse bool `default:'true'`

	// ID is the unique identifier of a collector
	ID uint32 // `default:'1'`

	// DetectCharset can enable character encoding detection for non-utf8 response bodies
	// without explicit charset declaration. This feature uses https://github.com/saintfish/chardet
	DetectCharset bool `default:'true'`

	// DetectMimeType
	DetectMimeType bool `default:'true'`

	// DetectTabular
	DetectTabular bool `default:'true'`

	// AnalyzeContent
	AnalyzeContent bool `default:'false'`

	// SummarizeContent
	SummarizeContent bool `default:'false'`

	// TopicModelling
	TopicModelling bool `default:'false'`

	// DebugMode
	DebugMode bool `default:'false'`

	// VerboseMode
	VerboseMode bool `default:'false'`

	// `required:"true"`
}
