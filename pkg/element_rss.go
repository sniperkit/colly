package colly

import (
	rssquery "github.com/sniperkit/gofeed/pkg"
)

// RSSElement is the representation of a RSS tag.
type RSSElement struct {

	////
	////// exported /////////////////////////////////////////////
	////

	// Name is the name of the tag
	Name string

	// Attrs specifies the attributes to extract
	Attrs []string

	// Request is the request object of the element's HTML document
	Request *Request

	// Response is the Response object of the element's HTML document
	Response *Response

	// Extractor points to...
	// Extractor *Extractor

	// DOM is the DOM object of the page. DOM is relative
	// to the current RSSElement and is either a html.Node or rssquery.Node
	// based on how the RSSElement was created.
	DOM interface{}

	////
	////// not exported /////////////////////////////////////////////
	////

	// attributes
	attributes interface{}

	//--- End
}

// RSSCallback is a type alias for OnRSS callback functions
type RSSCallback func(*RSSElement)

// rssCallbackContainer
type rssCallbackContainer struct {

	// Query specifies
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Function specifies
	Function RSSCallback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	//-- End
}

// NewRSSElementFromRSSFeed creates a RSSElement from a rssquery.Feed.
func NewRSSElementFromRSSFeed(resp *Response, fp *rssquery.Feed) *RSSElement {

	return &RSSElement{
		Request:  resp.Request,
		Response: resp,
	}
}
