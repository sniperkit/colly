package tabular

import (
	"io"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// TAB is the key for the tab compatible encodings
const TAB = "tab"

// TABDecoder...
type TABDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *TABDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *TABDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewTABDecoder return the right TAB decoder
func NewTABDecoder(isCollection bool) Decoder {
	if isCollection {
		return TABCollectionDecoder // Databook
	}
	return TABDecoder // Dataset
}

// TABDecoder implements the Decoder interface
func TABDecoder(r io.Reader, v *map[string]interface{}) error {
	// *v = mv
	return nil
}

// TABCollectionDecoder implements the Decoder interface over a collection
func TABCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	// *(v) = map[string]interface{}{"collection": mv}
	return nil
}
