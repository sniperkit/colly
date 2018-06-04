package main

import (
	"regexp"
)

// date/time related constants
const (
	APP_NAME            = "Colly - GolangLibs"
	APP_VERSION         = "0.0.1-alpha"
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

	defaultProxyTypes []string = []string{
		"default|socks5",
		"onion",
	}
)
