package sitemap

import (
	"io"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// SITEMAP is the key for the sitemap encoding
const SITEMAP = "sitemap"

// SITEMAPDecoder...
type SITEMAPDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *SITEMAPDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *SITEMAPDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewSITEMAPDecoder return the right SITEMAP decoder
func NewSITEMAPDecoder(isCollection bool) Decoder {
	if isCollection {
		return SITEMAPCollectionDecoder // Databook
	}
	return SITEMAPDecoder // Dataset
}

// SITEMAPDecoder implements the Decoder interface
func SITEMAPDecoder(r io.Reader, v *map[string]interface{}) error {
	// *v = mv
	return nil
}

// SITEMAPCollectionDecoder implements the Decoder interface over a collection/sitemap index
func SITEMAPCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	// *(v) = map[string]interface{}{"collection": mv}
	return nil
}
