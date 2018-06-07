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
	// "path"
	// "strings"

	// tabular "github.com/agrison/go-tablib"
	tabular "github.com/sniperkit/colly/plugins/data/transform/tabular"
)

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
*/

// TABElement is the representation of a TAB tag.
type TABElement struct {
	// Name is the name of the tag
	Name string

	// Dataset
	Dataset *tabular.Dataset

	// Text
	Text string

	// Request is the request object of the element's HTML document
	Request *Request

	// Response is the Response object of the element's HTML document
	Response *Response

	// err stores the loading error
	err error

	// DOM is the DOM object of the page. DOM is relative
	// to the current TABElement and is either a html.Node or jsonquery.Node
	// based on how the TABElement was created.
	// DOM interface{}
}

// NewTABElementFromTABNode creates a TABElement from a jsonquery.Node.
func NewTABElementFromTABNode(resp *Response, query string, ds *tabular.Dataset) *TABElement {

	// new TABElement object
	t := &TABElement{
		Dataset:  ds,
		Name:     query,
		Request:  resp.Request,
		Response: resp,
	}

	return t
}

// Extract
func (h *TABElement) Extract(pluckerConfig map[string]interface{}) string {
	return ""
}

// Extract
func (h *TABElement) Query(extractorConfig map[string]interface{}) string {
	return ""
}

// Header
func (h *TABElement) Header(key string) string {
	return ""
}

// Headers
func (h *TABElement) Headers() map[string]string {
	res := make(map[string]string, 0)
	/*
		for key, val := range h.Response.Headers {
			res[key] = val
		}
	*/
	return res
}
