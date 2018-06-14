package colly

import (
	// internal - core
	colly "github.com/sniperkit/colly/pkg"

	// internal - plugins
	link "github.com/sniperkit/colly/pkg/processor/link"
	link_media "github.com/sniperkit/colly/pkg/processor/link/media"
	pluck "github.com/sniperkit/colly/pkg/processor/pluck"
	regex "github.com/sniperkit/colly/pkg/processor/regex"
)

type ProcessorOn interface {
	Detect(resp *colly.Response) string
}
