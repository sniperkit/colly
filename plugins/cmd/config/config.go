package config

import (
	"fmt"
	"os"
	"time"
)

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

// collectors timeline vars
var (
	endAt   time.Time = time.Now()
	startAt time.Time = time.Now()
	upTime  time.Duration
)

type Global struct {
	*Setting
}

type Config struct {
	Environment string
	ENVPrefix   string
	Debug       bool
	Verbose     bool

	// Supported only for toml and yaml files.
	// json does not currently support this: https://github.com/golang/go/issues/15314
	// This setting will be ignored for json files.
	ErrorOnUnmatchedKeys bool
}

// New initialize a Configor
func New(config *Config) *Global {
	if config == nil {
		config = &Setting{}
	}

	if os.Getenv("CONFIGOR_DEBUG_MODE") != "" {
		config.Debug = true
	}

	if os.Getenv("CONFIGOR_VERBOSE_MODE") != "" {
		config.Verbose = true
	}

	return &Global{Config: config}
}

// GetEnvironment get environment
func (g *Global) GetEnvironment() string {
	if g.Environment == "" {
		if env := os.Getenv("CONFIGOR_ENV"); env != "" {
			return env
		}
		if testRegexp.MatchString(os.Args[0]) {
			return "test"
		}
		return "development"
	}
	return g.Environment
}

// GetErrorOnUnmatchedKeys returns a boolean indicating if an error should be
// thrown if there are keys in the config file that do not correspond to the
// config struct
func (g *Global) GetErrorOnUnmatchedKeys() bool {
	return g.ErrorOnUnmatchedKeys
}

// Load will unmarshal configurations to struct from files that you provide
func (g *Global) Load(config interface{}, files ...string) error {
	defer func() {
		if g.Setting.Debug || g.Setting.Verbose {
			fmt.Printf("Configuration:\n  %#v\n", config)
		}
	}()

	for _, file := range g.getConfigurationFiles(files...) {
		if g.Setting.Debug || g.Setting.Verbose {
			fmt.Printf("Loading configurations from file '%v'...\n", file)
		}
		if err := processFile(config, file, g.GetErrorOnUnmatchedKeys()); err != nil {
			return err
		}
	}

	prefix := g.getENVPrefix(config)
	if prefix == "-" {
		return g.processTags(config)
	}
	return g.processTags(config, prefix)
}

// ENV return environment
func ENV() string {
	return New(nil).GetEnvironment()
}

// Load will unmarshal configurations to struct from files that you provide
func Load(config interface{}, files ...string) error {
	return New(nil).Load(config, files...)
}
