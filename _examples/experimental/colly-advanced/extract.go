package main

var (
	// extracted patterns
	libraries []library
	entries   map[string]bool
	links     []string = []string{} // Array containing all the known URLs in a sitemap
)

// library stores information about indexed golang library in golanglibs.com
type library struct {
	Title       string
	Description string
	Categories  []string
	URI         string
	URL         string
	Stars       int
}

func init() {
	libraries = make([]library, 0)
	entries = make(map[string]bool, 0)
}
