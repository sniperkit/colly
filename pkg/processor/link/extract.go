package pluck

import (
	// internal - core
	colly "github.com/sniperkit/colly/pkg"
)

// LINKProcessor defines...
type LINKProcessor struct {
}

// Detect...
func (p *LINKProcessor) Detect(resp *colly.Response) string {
	return ""
}
