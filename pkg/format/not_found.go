package decoder

import (
	// internal - core
	colly "github.com/sniperkit/colly/pkg"
)

// NOTFoundDecoder represents...
type NOTFoundDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *NOTFoundDecoder) Check(resp *colly.Response) error { return ErrNotImplementedYet }

// Check...
func (c *NOTFoundDecoder) Decoder(resp *colly.Response, v *map[string]interface{}) error {
	return ErrNotImplementedYet
}
