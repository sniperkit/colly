package main

import (
	"regexp"
)

// date/time related constants
const (
	APP_NAME            = "Colly - GolangLibs"
	DEFAULT_DATE_FORMAT = "Jan 02, 2006" // DATE_FORMAT default format date
)

// sitemap related constants
const (
	DEFAULT_SITEMAP_URL             string = "https://golanglibs.com/sitemap.txt"
	DEFAULT_SITEMAP_BASENAME        string = "sitemap.txt"
	DEFAULT_SITEMAP_EXTENSION       string = "txt"
	DEFAULT_SITEMAP_ENCODING        string = "plain/text"
	DEFAULT_SITEMAP_TXT_COLUMN_NAME string = "urls"
	DEFAULT_SITEMAP_TXT_COLUMN_ID   int    = 1
	DEFAULT_SITEMAP_XML_XPATH_LOC   string = "//urlset/url/loc"
)

var (
	defaultDomain                string   = "https://golanglibs\\.com"
	defaultStorageExportDir      string   = "./shared/storage/export"
	defaultStorageCacheDir       string   = "./shared/storage/cache"
	defaultStorageLogDir         string   = "./shared/logs"
	defaultStorageSitemapDirname string   = defaultStorageCacheDir + "/sitemaps"
	defaultStorageDirs           []string = []string{
		defaultStorageCacheDir,
		defaultStorageLogDir,
		defaultStorageExportDir,
		defaultStorageSitemapDirname,
	}
	defaultAllowedDomains []string = []string{
		"golanglibs.com",
		"golanglibs.com:443",
	}

	defaultDisabledURLFilters []*regexp.Regexp = []*regexp.Regexp{
		regexp.MustCompile("(.*)?sort=$"),
	}

	defaultAllowedURLFilters []*regexp.Regexp = []*regexp.Regexp{
		regexp.MustCompile("^" + defaultDomain + "/?page=([0-9])+$"),
		regexp.MustCompile("^" + defaultDomain + "/top$"),
		regexp.MustCompile("^" + defaultDomain + "/categories$"),
		regexp.MustCompile("^" + defaultDomain + "/random$"),
		regexp.MustCompile("^" + defaultDomain + "/active$"),
		regexp.MustCompile("^" + defaultDomain + "/tagged$"),
		regexp.MustCompile("^" + defaultDomain + "/repo/$"),
		regexp.MustCompile("^" + defaultDomain + "/category/$"),
		regexp.MustCompile("^" + defaultDomain + "/tag/$"),
		regexp.MustCompile("^" + defaultDomain + "/similar/$"),
		regexp.MustCompile("^/?page=([0-9])+$"),
		regexp.MustCompile("^/top$"),
		regexp.MustCompile("^/categories$"),
		regexp.MustCompile("^/random$"),
		regexp.MustCompile("^/active$"),
		regexp.MustCompile("^/tagged$"),
		regexp.MustCompile("^/repo/$"),
		regexp.MustCompile("^/category/$"),
		regexp.MustCompile("^/tag/$"),
		regexp.MustCompile("^/similar/$"),
		regexp.MustCompile("^(.*)/?page=([0-9])+$"),
		regexp.MustCompile("^(.*)/top$"),
		regexp.MustCompile("^(.*)/categories$"),
		regexp.MustCompile("^(.*)/random$"),
		regexp.MustCompile("^(.*)/active$"),
		regexp.MustCompile("^(.*)/tagged$"),
		regexp.MustCompile("^(.*)/repo/$"),
		regexp.MustCompile("^(.*)/category/$"),
		regexp.MustCompile("^(.*)/tag/$"),
		regexp.MustCompile("^(.*)/similar/$"),
		// regexp.MustCompile(defaultDomain + "/(e.+)$"),
		// regexp.MustCompile(defaultDomain + "/b.+"),
	}
)
