package html

import (
	"io"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// HTML is the key for the html encoding
const HTML = "html"

// HTMLDecoder...
type HTMLDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *HTMLDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *HTMLDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewJSONDecoder return the right JSON decoder
func NewHTMLDecoder(isCollection bool) Decoder {
	if isCollection {
		return JSONCollectionDecoder
	}
	return HTMLDecoder
}

// HTMLDecoder implements the Decoder interface
func HTMLDecoder(r io.Reader, v *map[string]interface{}) error {
	// *v = doc
	return nil
}

// XMLCollectionDecoder implements the Decoder interface over a collection
func HTMLCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	// *(v) = map[string]interface{}{"collection": doc}
	return nil
}
