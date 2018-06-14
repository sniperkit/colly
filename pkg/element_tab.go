package colly

import (
	tabular "github.com/sniperkit/colly/plugins/data/transform/tabular"
)

// FORMAT_TAB defines
type FORMAT_TAB string

const (

	// TAB_JSON is the key for `JSON` encoding
	TAB_JSON FORMAT_TAB = "json"

	// TAB_YAML is the key for `YAML` encoding
	TAB_YAML FORMAT_TAB = "yaml"

	// TAB_XML is the key for `XML` encoding
	TAB_XML FORMAT_TAB = "xml"

	// TAB_CSV is the key for `CSV` encoding (columns sperated by comma or semi-colon).
	TAB_CSV FORMAT_TAB = "csv"

	// TAB_TSV is the key for `TSV` encoding (tab separated columns).
	TAB_TSV FORMAT_TAB = "tsv"

	// TAB_XLSX is the key for `XLSX` encoding. (Microsoft Excel)
	TAB_XLSX FORMAT_TAB = "xlsx"

	// TAB_ASCII is the key for `ASCII` encoding
	TAB_ASCII FORMAT_TAB = "ascii"

	// TAB_MD is the key for `MARKDOWN` encoding
	TAB_MD FORMAT_TAB = "md"

	// TAB_MD is the key for `MARKDOWN` encoding
	TAB_MYSQL FORMAT_TAB = "mysql"

	// TAB_POSTGRES is the key for `POSTGRES` encoding
	TAB_POSTGRES FORMAT_TAB = "postgres"

	//-- End
)

var (

// TAB_FORMATS_IMPORT is the slice strings indexing all formats supported for tabular data import.
// TAB_FORMATS_IMPORT = []string{TAB_JSON, TAB_YAML, TAB_XML, TAB_CSV, TAB_TSV}

// TAB_FORMATS_EXPORT is the slice strings indexing all export formats supported for tabular data.
// TAB_FORMATS_EXPORT = []string{TAB_JSON, TAB_YAML, TAB_XML, TAB_CSV, TAB_TSV, TAB_XLSX, TAB_MD, TAB_MYSQL, TAB_POSTGRES}

//-- End
)

// Create aliases with struct defined in tabular pkg
type (

	// `TABDynamicColumn` is an alias of...
	TABDynamicColumn = tabular.DynamicColumn

	// `TABFormat` is an alias of...
	TABFormat = tabular.Format

	// `TABHook` is an alias of...
	TABHook = tabular.Hook

	// `TABOutput` is an alias of...
	TABOutput = tabular.Printer

	// `TABRanger` is an alias of...
	TABRanger = tabular.Ranger

	// `TABSlicer` is an alias of...
	TABSlicer = tabular.Slicer

	// `TABWriter` is an alias of...
	TABWriter = tabular.Writer

	// `TABDynamicColumn` is an alias of...
	// TABSelector = tabular.Selector

	//-- End
)

// TABCallback is a type alias for OnTAB callback functions
type TABCallback func(*TABElement)

// tabCallbackContainer specifies the struct
type tabCallbackContainer struct {

	// Query specifies...
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Hook specifies...
	Hook *TABHook `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Function specifies...
	Function TABCallback `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}

// TABHooks specifies
type TABHooks struct {

	// Enabled
	Enabled bool `default:"true" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled" csv:"enabled"`

	// Registry
	Registry map[string]*TABHook `json:"registry" yaml:"registry" toml:"registry" xml:"-" ini:"registry" csv:"registry"`
}

// TABElement is the representation of a TAB tag.
type TABElement struct {

	////
	////// exported //////////////////////////////////////////////////
	////

	// Name is the name of the tag
	Name string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Event is the name of the pre-processing task to trigger
	Query string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Dataset represents...
	Dataset *tabular.Dataset `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Hooks represents...
	Hook *TABHook `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Extractor
	// Extractor *Extractor `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Text
	Text string `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Request is the request object of the element's HTML document
	Request *Request `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	// Response is the Response object of the element's HTML document
	Response *Response `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

	////
	////// not exported /////////////////////////////////////////////
	////

	// err stores the loading error
	errs []error `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
}

// NewTABElementFromTABNode instanciates a new TABElement extracted with a tabular data query
func NewTABElementFromTABNode(resp *Response, query string, ds *tabular.Dataset) *TABElement {
	t := &TABElement{
		Dataset:  ds,
		Name:     query,
		Request:  resp.Request,
		Response: resp,
	}
	return t
}

// NewTABElementFromTABSelect instanciates a new TABElement extracted with a global processing hook
func NewTABElementFromTABSelect(resp *Response, hook *TABHook, ds *tabular.Dataset) *TABElement {
	t := &TABElement{
		Dataset:  ds,
		Hook:     hook,
		Request:  resp.Request,
		Response: resp,
	}
	return t
}
