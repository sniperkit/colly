package colly

import (
	"io"
)

// HTML is the key for the html encoding
const HTML = "html"

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

// ContentTypeHTML...
type ContentTypeHTML struct {
	resp *Response
	err  error
}

// Check...
func (c *ContentTypeHTML) Check(resp *Response) error { return ErrNotImplementedYet }
