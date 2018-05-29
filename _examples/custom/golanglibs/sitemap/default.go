package main

const (
	defaultAllowedDomains        string = "golanglibs.com"
	defaultSitemapURL            string = "https://golanglibs.com/sitemap.txt"
	defaultSitemapEncoding       string = "plain/text"
	defaultSitenapTXT_ColumnID   int    = 1
	defaultSitenapTXT_ColumnName string = "urls"
	defaultSitenapXML_XPath      string = "//urlset/url/loc"
)

var (
	defaultStorageDirs []string = []string{
		"./shared/storage/cache/sitemaps",
		"./shared/storage/export/sitemaps",
	}
)
