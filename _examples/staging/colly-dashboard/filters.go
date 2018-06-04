package main

import (
	"regexp"
)

// collector - default regexp filters
var (
	defaultAllowedDomains []string = []string{
		"golanglibs.com",
		"golanglibs.com:443",
	}
	defaultDisabledURLFilters []*regexp.Regexp = []*regexp.Regexp{
		regexp.MustCompile("(.*)?sort=$"),
	}

	defaultAllowedURLFilterList []string = []string{
		"^" + defaultDomain + "/?page=([0-9])+$",
		"^" + defaultDomain + "/top$",
		"^" + defaultDomain + "/categories$",
		"^" + defaultDomain + "/random$",
		"^" + defaultDomain + "/active$",
		"^" + defaultDomain + "/tagged$",
		"^" + defaultDomain + "/repo/$",
		"^" + defaultDomain + "/category/$",
		"^" + defaultDomain + "/tag/$",
		"^" + defaultDomain + "/similar/$",
		"^/?page=([0-9])+$",
		"^/top$",
		"^/categories$",
		"^/random$",
		"^/active$",
		"^/tagged$",
		"^/repo/$",
		"^/category/$",
		"^/tag/$",
		"^/similar/$",
		"^(.*)/?page=([0-9])+$",
		"^(.*)/top$",
		"^(.*)/categories$",
		"^(.*)/random$",
		"^(.*)/active$",
		"^(.*)/tagged$",
		"^(.*)/repo/$",
		"^(.*)/category/$",
		"^(.*)/tag/$",
		"^(.*)/similar/$",
		defaultDomain + "/(e.+)$",
		defaultDomain + "/b.+",
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
