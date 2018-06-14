package csv

import (
	"io"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// CSV is the key for the html encoding
const CSV = "csv"

// CSVDecoder...
type CSVDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *CSVDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *CSVDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewJSONDecoder return the right JSON decoder
func NewCSVDecoder(isCollection bool) Decoder {
	if isCollection {
		return JSONCollectionDecoder
	}
	return CSVDecoder
}

// CSVDecoder implements the Decoder interface
func CSVDecoder(r io.Reader, v *map[string]interface{}) error {
	// *v = doc
	return nil
}

// XMLCollectionDecoder implements the Decoder interface over a collection
func CSVCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	// *(v) = map[string]interface{}{"collection": doc}
	return nil
}
