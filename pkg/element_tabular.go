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

/*
	Export formats supported:
	- JSON (Sets + Books)
	- YAML (Sets + Books)
	- XLSX (Sets + Books)
	- XML (Sets + Books)
	- TSV (Sets)
	- CSV (Sets)
	- ASCII + Markdown (Sets)
	- MySQL (Sets)
	- Postgres (Sets)

	Loading formats supported:
	- JSON (Sets + Books)
	- YAML (Sets + Books)
	- XML (Sets)
	- CSV (Sets)
	- TSV (Sets)
*/

package colly

/*
import (
	"github.com/sniperkit/colly/plugins/convert/agnostic-tablib"
)
*/

// TABElement is the representation of a XML tag.
type TABElement struct {
	// Name is the name of the tag
	Name       string
	Text       string
	attributes interface{}
	// Request is the request object of the element's HTML document
	Request *Request
	// Response is the Response object of the element's HTML document
	Response *Response
	// DOM is the DOM object of the page. DOM is relative
	// to the current TabElement and is either a html.Node or xmlquery.Node
	// based on how the TabElement was created.
	DOM    interface{}
	isHTML bool
}

/*

// NewTabElementFromHTMLNode creates a TabElement from a html.Node.
func NewTabElementFromPlainNode(resp *Response, s *html.Node) *TabElement {
	return &TabElement{
		Name:       s.Data,
		Request:    resp.Request,
		Response:   resp,
		Text:       htmlquery.InnerText(s),
		DOM:        s,
		attributes: s.Attr,
		isHTML:     false,
	}
}

// Attr returns the selected attribute of a HTMLElement or empty string
// if no attribute found
func (h *TabElement) Attr(k string) string {
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
*/
