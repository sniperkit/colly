package json

import (
	"encoding/json"
	"io"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// JSON is the key for the json encoding
const JSON = "json"

// ContentTypeJSON...
type JSONDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *JSONDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *JSONDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewJSONDecoder return the right JSON decoder
func NewJSONDecoder(isCollection bool) Decoder {
	if isCollection {
		return JSONCollectionDecoder
	}
	return JSONDecoder
}

// JSONDecoder implements the Decoder interface
func JSONDecoder(r io.Reader, v *map[string]interface{}) error {
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(v)
}

// JSONCollectionDecoder implements the Decoder interface over a collection
func JSONCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	var collection []interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": collection}
	return nil
}
