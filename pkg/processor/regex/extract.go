package regex

import (
	// internal - core
	colly "github.com/sniperkit/colly/pkg"
)

// REGEXProcessor defines...
type REGEXProcessor struct {
}

// Detect...
func (p *REGEXProcessor) Detect(resp *colly.Response) string {
	return ""
}
