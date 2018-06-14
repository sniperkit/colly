package json

import (
	"strings"

	// external
	jsonql_v1 "github.com/sniperkit/jsonql/pkg"

	// internal
	jsonquery "github.com/sniperkit/colly/plugins/data/extract/query/json"
)

// JsonParser defines
type JsonParser string

const (

	// MXJ is the key for...
	MXJ JsonParser = "mxj" // https://github.com/clbanning/mxj

	// GABS is the key for...
	GABS JsonParser = "gabs" // https://github.com/Jeffail/gabs

	// GJSON is the key for...
	GJSON JsonParser = "gjson" // https://github.com/tidwall/gjson

	// LAZYJSON is the key for...
	LAZYJSON JsonParser = "lazyjson" // https://github.com/qw4990/lazyjson

	// FASTJSON is the key for...
	FASTJSON JsonParser = "fastjson" // https://github.com/valyala/fastjson

	// FFJSON is the key for...
	FFJSON JsonParser = "ffjson" // https://github.com/pquerna/ffjson

	// EASYJSON is the key for...
	EASYJSON JsonParser = "easyjson" // https://github.com/mailru/easyjson

	// JSONPARSER is the key for...
	JSONPARSER JsonParser = "jsonparser" // https://github.com/buger/jsonparser

	// DJSON is the key for...
	DJSON JsonParser = "djson" // https://github.com/a8m/djson

	// JSNM is the key for...
	JSNM JsonParser = "jsnm" // https://github.com/toukii/jsnm

	// JSONSTREAM is the key for...
	JSONSTREAM JsonParser = "jsonstream" // https://github.com/pb-/jsonstream

	// JSONEZ is the key for...
	JSONEZ JsonParser = "jsonez" // https://github.com/srikanth2212/jsonez

	// JSON_DEFAULT is the key for...
	JSON_DEFAULT JsonParser = "json" // encoding/json

	//-- End
)

// JSONElement is the representation of a JSON tag.
type JSONElement struct {

	////
	////// exported /////////////////////////////////////////////
	////

	// Name is the name of the tag
	Name string
	// Text is the content node
	Text string

	// Request is the request object of the element's HTML document
	Request *Request

	// Response is the Response object of the element's HTML document
	Response *Response

	// Extractor
	// Extractor *Extractor

	// DOM is the DOM object of the page. DOM is relative
	// to the current JSONElement and is either a html.Node or jsonquery.Node
	// based on how the JSONElement was created.
	DOM interface{}

	////
	////// not exported /////////////////////////////////////////////
	////

	// jsonql_v1 --> github.com/elgs/jsonql
	jsonql_v1 *jsonql_v1.JSONQL

	//-- End
}

// JSONCallback is a type alias for OnJSON callback functions
type JSONCallback func(*JSONElement)

// JSONCallbackContainer
type JSONCallbackContainer struct {

	// Query specifies
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Function specifies
	Function JSONCallback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	//-- End
}

// NewJSONElementFromJSONNode creates a JSONElement from a jsonquery.Node.
func NewJSONElementFromJSONNode(resp *Response, s *jsonquery.Node) *JSONElement {
	return &JSONElement{
		Name:     s.Data,
		Request:  resp.Request,
		Response: resp,
		Text:     s.InnerText(),
		DOM:      s,
	}
}

// JSON: Find, FindOne, Extract, Header, Headers, Keys, Values, Map,

// Extract
func (h *JSONElement) Extract(pluckerConfig map[string]interface{}) string {
	return ""
}

// Header
func (h *JSONElement) Header(key string) (value string) {
	value = strings.TrimSpace(h.Response.Headers.Get(key))
	return
}

// Headers
func (h *JSONElement) Headers() map[string]string {
	res := make(map[string]string, 0)
	/*
		for key, val := range h.Response.Headers {
			res[key] = val
		}
	*/
	return res
}

// FindOne
func (h *JSONElement) FindOne(xpathQuery string) string {

	n := jsonquery.FindOne(h.DOM.(*jsonquery.Node), xpathQuery)
	if n == nil {
		return ""
	}
	return strings.TrimSpace(n.InnerText())
}

// Find
func (h *JSONElement) Find(xpathQuery, attrName string) []string {
	var res []string
	child := jsonquery.Find(h.DOM.(*jsonquery.Node), xpathQuery)
	for _, node := range child {
		res = append(res, node.InnerText())
	}
	return res
}
