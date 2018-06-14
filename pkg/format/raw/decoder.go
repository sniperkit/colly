package raw

import (
	"io"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// RAW is the key for the html encoding
const RAW = "raw"

// RAWDecoder...
type RAWDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *RAWDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *RAWDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewJSONDecoder return the right JSON decoder
func NewRAWDecoder(isCollection bool) Decoder {
	if isCollection {
		return JSONCollectionDecoder
	}
	return RAWDecoder
}

// RAWDecoder implements the Decoder interface
func RAWDecoder(r io.Reader, v *map[string]interface{}) error {
	// *v = doc
	return nil
}

// XMLCollectionDecoder implements the Decoder interface over a collection
func RAWCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	// *(v) = map[string]interface{}{"collection": doc}
	return nil
}
