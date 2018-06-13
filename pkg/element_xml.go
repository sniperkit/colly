// Copyright 2018 Adam Tauber
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package colly

import (
	"encoding/xml"
	"strings"

	// external
	"golang.org/x/net/html"

	// internal
	htmlquery "github.com/sniperkit/colly/plugins/data/extract/query/html"
	xmlquery "github.com/sniperkit/colly/plugins/data/extract/query/xml"
)

// XMLElement is the representation of a XML tag.
type XMLElement struct {

	////// exported //////////////////////////////////////////////////
	// Name is the name of the tag
	Name string

	// Text stores...
	Text string

	// Request is the request object of the element's HTML document
	Request *Request

	// Response is the Response object of the element's HTML document
	Response *Response

	// Extractor points to...
	// Extractor *Extractor

	// DOM is the DOM object of the page. DOM is relative
	// to the current XMLElement and is either a html.Node or xmlquery.Node
	// based on how the XMLElement was created.
	DOM interface{}

	////// not exported /////////////////////////////////////////////
	// attributes
	attributes interface{}
	// isHTML specifies...
	isHTML bool

	//--- End
}

// NewXMLElementFromHTMLNode creates a XMLElement from a html.Node.
func NewXMLElementFromHTMLNode(resp *Response, s *html.Node) *XMLElement {
	return &XMLElement{
		Name:       s.Data,
		Request:    resp.Request,
		Response:   resp,
		Text:       htmlquery.InnerText(s),
		DOM:        s,
		attributes: s.Attr,
		isHTML:     true,
	}
}

// NewXMLElementFromXMLNode creates a XMLElement from a xmlquery.Node.
func NewXMLElementFromXMLNode(resp *Response, s *xmlquery.Node) *XMLElement {
	return &XMLElement{
		Name:       s.Data,
		Request:    resp.Request,
		Response:   resp,
		Text:       s.InnerText(),
		DOM:        s,
		attributes: s.Attr,
		isHTML:     false,
	}
}

// XML: Attr, ChildText, ChildAttr, ChildAttrs

// Attr returns the selected attribute of a HTMLElement or empty string
// if no attribute found
func (h *XMLElement) Attr(k string) string {
	if h.isHTML {
		for _, a := range h.attributes.([]html.Attribute) {
			if a.Key == k {
				return a.Val
			}
		}
	} else {
		for _, a := range h.attributes.([]xml.Attr) {
			if a.Name.Local == k {
				return a.Value
			}
		}
	}
	return ""
}

// ChildText returns the concatenated and stripped text content of the matching
// elements.
func (h *XMLElement) ChildText(xpathQuery string) string {
	if h.isHTML {
		return strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(h.DOM.(*html.Node), xpathQuery)))
	}
	n := xmlquery.FindOne(h.DOM.(*xmlquery.Node), xpathQuery)
	if n == nil {
		return ""
	}
	return strings.TrimSpace(n.InnerText())
}

// ChildAttr returns the stripped text content of the first matching
// element's attribute.
func (h *XMLElement) ChildAttr(xpathQuery, attrName string) string {
	if h.isHTML {
		child := htmlquery.FindOne(h.DOM.(*html.Node), xpathQuery)
		if child != nil {
			for _, attr := range child.Attr {
				if attr.Key == attrName {
					return strings.TrimSpace(attr.Val)
				}
			}
		}
	} else {
		child := xmlquery.FindOne(h.DOM.(*xmlquery.Node), xpathQuery)
		if child != nil {
			for _, attr := range child.Attr {
				if attr.Name.Local == attrName {
					return strings.TrimSpace(attr.Value)
				}
			}
		}
	}

	return ""
}

// ChildAttrs returns the stripped text content of all the matching
// element's attributes.
func (h *XMLElement) ChildAttrs(xpathQuery, attrName string) []string {
	var res []string
	if h.isHTML {
		htmlquery.FindEach(h.DOM.(*html.Node), xpathQuery, func(i int, child *html.Node) {
			for _, attr := range child.Attr {
				if attr.Key == attrName {
					res = append(res, strings.TrimSpace(attr.Val))
				}
			}
		})
	} else {
		xmlquery.FindEach(h.DOM.(*xmlquery.Node), xpathQuery, func(i int, child *xmlquery.Node) {
			for _, attr := range child.Attr {
				if attr.Name.Local == attrName {
					res = append(res, strings.TrimSpace(attr.Value))
				}
			}
		})
	}
	return res
}
