package colly

import (
	xurls "github.com/sniperkit/xurls/pkg"
)

// IsProcessorLinkStrict sets the flag if the parser is in strict or relaxed mode
var IsProcessorLinkStrict = false

// extractLink function...
func extractLink(content string) (link string) {
	switch IsProcessorLinkStrict {
	case true:
		link = xurls.Strict().FindString(content)
	case false:
		link = xurls.Relaxed().FindString(content)
	}
	return
}

// extractLinks function...
func extractLinks(content string, limit int) (links []string) {
	if limit == 0 {
		limit = -1
	}
	switch IsProcessorLinkStrict {
	case true:
		links = xurls.Strict().FindAllString(content, limit)
	case false:
		links = xurls.Relaxed().FindAllString(content, limit)
	}
	return
}
