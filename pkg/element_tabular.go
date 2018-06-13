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
	Enabled  bool                `default:"true" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`
	Registry map[string]*TABHook `json:"registry" yaml:"registry" toml:"registry" xml:"-" ini:"registry" csv:"registry"`
}

// TABElement is the representation of a TAB tag.
type TABElement struct {

	// Hooks represents...
	Hook *TABHook `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	////// exported //////////////////////////////////////////////////
	// Name is the name of the tag
	Name string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Event is the name of the pre-processing task to trigger
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Dataset represents...
	Dataset *tabular.Dataset `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Extractor
	Extractor *Extractor `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Text
	Text string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Request is the request object of the element's HTML document
	Request *Request `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Response is the Response object of the element's HTML document
	Response *Response `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	////// not exported /////////////////////////////////////////////
	// err stores the loading error
	err error `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
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

// DATA: Slice, Select

// Slice
func (h *TABElement) Slice(pluckerConfig map[string]interface{}) string {
	return ""
}

// Select
func (h *TABElement) Select(extractorConfig map[string]interface{}) string {
	return ""
}
