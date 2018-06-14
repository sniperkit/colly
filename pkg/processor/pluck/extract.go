package pluck

import (
	// internal - core
	colly "github.com/sniperkit/colly/pkg"
)

// PLUCKProcessor defines...
type PLUCKProcessor struct {
}

// Detect...
func (p *PLUCKProcessor) Detect(resp *colly.Response) string {
	return ""
}
