package colly

import (
	tabular "github.com/sniperkit/colly/plugins/data/transform/tabular"
	// pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

/*
	Notes:

	- Export formats supported:
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

type TABOutput = tabular.Printer

type TABFormat = tabular.Format

type TABSlicer = tabular.Slicer

type TABRanger = tabular.Ranger

type TABHook = tabular.Hook

type TABDynamicColumn = tabular.DynamicColumn

// type TABSelector = tabular.Selector

type TABWriter = tabular.Writer

type TABHooks struct {
	Enabled  bool
	Registry map[string]*TABHook
}

// TABElement is the representation of a TAB tag.
type TABElement struct {

	// Hooks represents...
	Hook *TABHook

	////// exported //////////////////////////////////////////////////
	// Name is the name of the tag
	Name string

	// Event is the name of the pre-processing task to trigger
	Query string

	// Dataset represents...
	Dataset *tabular.Dataset

	// Datasets represents...
	// Datasets []*tabular.Dataset

	// Extractor
	Extractor *Extractor

	// Text
	Text string

	// Request is the request object of the element's HTML document
	Request *Request

	// Response is the Response object of the element's HTML document
	Response *Response

	////// not exported /////////////////////////////////////////////
	// err stores the loading error
	err error
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

func NewTABElementFromTABSelect(resp *Response, hook *TABHook, ds *tabular.Dataset) *TABElement {
	// pp.Println("hook=", hook)
	// new TABElement object
	t := &TABElement{
		Dataset:  ds,
		Hook:     hook,
		Request:  resp.Request,
		Response: resp,
	}
	return t
}

// Slice
func (h *TABElement) Slice(pluckerConfig map[string]interface{}) string {
	return ""
}

// Select
func (h *TABElement) Select(extractorConfig map[string]interface{}) string {
	return ""
}
