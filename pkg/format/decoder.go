package decoder

import (
	"io"

	// internal - encodings
	"github.com/sniperkit/colly/pkg/decoder/csv"
	"github.com/sniperkit/colly/pkg/decoder/html"
	"github.com/sniperkit/colly/pkg/decoder/json"
	"github.com/sniperkit/colly/pkg/decoder/raw"
	"github.com/sniperkit/colly/pkg/decoder/xml"

	// internal - formats
	"github.com/sniperkit/colly/pkg/decoder/rss"
	"github.com/sniperkit/colly/pkg/decoder/sitemap"
	"github.com/sniperkit/colly/pkg/decoder/tabular"
)

/*
Package encoding provides Decoding implementations.
Decode decodes HTTP responses:
	resp, _ := http.Get("http://api.example.com/")
	...
	var data map[string]interface{}
	err := JSONDecoder(resp.Body, &data)
*/

// A Decoder is a function that reads from the reader and decodes it into an map of interfaces
type Decoder func(r io.Reader, v *map[string]interface{}) error

// DecodeOn struct...
type DecodeOn interface {
	Check(resp *colly.Response) error
	Decode(r io.Reader, v *map[string]interface{}) error
}
