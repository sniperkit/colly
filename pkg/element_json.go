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
	"strings"

	"github.com/antchfx/jsonquery"
)

// JSONElement is the representation of a JSON tag.
type JSONElement struct {
	// Name is the name of the tag
	Name string
	Text string
	// Request is the request object of the element's HTML document
	Request *Request
	// Response is the Response object of the element's HTML document
	Response *Response
	// DOM is the DOM object of the page. DOM is relative
	// to the current JSONElement and is either a html.Node or jsonquery.Node
	// based on how the JSONElement was created.
	DOM interface{}
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

// Attr returns the selected attribute of a JSONElement or empty string
// if no attribute found
func (h *JSONElement) Node(xpathQuery string) string {
	n := jsonquery.FindOne(h.DOM.(*jsonquery.Node), xpathQuery)
	if n == nil {
		return ""
	}
	return strings.TrimSpace(n.InnerText())
}

// ChildText returns the concatenated and stripped text content of the matching
// elements.
func (h *JSONElement) ChildText(xpathQuery string) string {

	n := jsonquery.FindOne(h.DOM.(*jsonquery.Node), xpathQuery)
	if n == nil {
		return ""
	}
	return strings.TrimSpace(n.InnerText())
}

// ChildAttr returns the stripped text content of the first matching
// element's attribute.
func (h *JSONElement) ChildAttr(xpathQuery, attrName string) string {
	child := jsonquery.FindOne(h.DOM.(*jsonquery.Node), xpathQuery)
	if child != nil {
		return strings.TrimSpace(child.InnerText())
	}
	return ""
}

// ChildAttrs returns the stripped text content of all the matching
// element's attributes.
func (h *JSONElement) ChildAttrs(xpathQuery, attrName string) []string {
	var res []string
	child := jsonquery.Find(h.DOM.(*jsonquery.Node), xpathQuery)
	for _, node := range child {
		res = append(res, node.InnerText())
	}
	return res
}
