package xml

import (
	"io"

	// external
	mxj "github.com/sniperkit/mxj/pkg"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// XML is the key for the xml encoding
const XML = "xml"

// ContentTypeXML...
type XMLDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *XMLDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *XMLDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewXMLDecoder return the right XML decoder
func NewXMLDecoder(isCollection bool) Decoder {
	if isCollection {
		return XMLCollectionDecoder
	}
	return XMLDecoder
}

// XMLDecoder implements the Decoder interface
func XMLDecoder(r io.Reader, v *map[string]interface{}) error {
	mv, err := mxj.NewMapXmlReader(r)
	if err != nil {
		return err
	}
	*v = mv
	return nil

}

// XMLCollectionDecoder implements the Decoder interface over a collection
func XMLCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	mv, err := mxj.NewMapXmlReader(r)
	if err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": mv}
	return nil
}
