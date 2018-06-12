package tablib

import (
	"errors"
	"regexp"
	// "sync"
)

type Selector = Hook

type Hook struct {

	// exported
	Id             string                 `json:"identifier" yaml:"identifier" toml:"identifier" xml:"identifier" ini:"identifier"`
	Slug           string                 `json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty" xml:"slug,omitempty" ini:"slug,omitempty"`
	HasPrefix      string                 `json:"has_prefix,omitempty" yaml:"has_prefix,omitempty" toml:"has_prefix,omitempty" xml:"hasPrefix,omitempty" ini:"hasPrefix,omitempty"`
	HasSuffix      string                 `json:"has_suffixslug,omitempty" yaml:"has_suffix,omitempty" toml:"has_suffix,omitempty" xml:"hasSuffix,omitempty" ini:"hasSuffix,omitempty"`
	PatternURL     string                 `json:"pattern_url" yaml:"pattern_url" toml:"pattern_url" xml:"patternURL" ini:"patternURL"`
	PatternRegex   *regexp.Regexp         `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
	DynamicColumns map[string]DynamicFunc `json:"dynamic_columns,omitempty" yaml:"dynamic_columns,omitempty" toml:"dynamic_columns,omitempty" xml:"dynamicColumns,omitempty" ini:"dynamicColumns,omitempty"`
	Headers        []string               `required:"true" json:"headers" yaml:"headers" toml:"headers" xml:"headers" ini:"headers"`
	Mixed          []interface{}          `json:"mixed,omitempty" yaml:"mixed,omitempty" toml:"mixed,omitempty" xml:"mixed,omitempty" ini:"mixed,omitempty"`
	DataQL         *DataQL                `json:"data_ql" yaml:"data_ql" toml:"data_ql" xml:"dataQL" ini:"dataQL"`
	Slicer         *Slicer                `json:"slicer" yaml:"slicer" toml:"slicer" xml:"slicer" ini:"slicer"`
	Writer         *Writer                `json:"writer,omitempty" yaml:"writer,omitempty" toml:"writer,omitempty" xml:"writer,omitempty" ini:"writer,omitempty"`
	Printer        *Printer               `json:"printer,omitempty" yaml:"printer,omitempty" toml:"printer,omitempty" xml:"printer,omitempty" ini:"printer,omitempty"`

	// not exported
	errs []error

	//-- End
}

// PatternURL
func (s *Selector) PatternRegexp(rule string) *Selector {
	if rule == "" {
		s.errs = append(s.errs, errors.New("Empty pattern regex, skipping..."))
		return s
	}
	s.PatternURL = rule
	s.PatternRegex = regexp.MustCompile(rule)
	return s
}

// type DynamicColumn func([]interface{}) interface{}
func (s *Selector) AddDynamicColumn(fn DynamicColumn, header string, cols ...int) *Selector {
	if s.DynamicColumns == nil {
		s.DynamicColumns = make(map[string]DynamicFunc, 0)
	}
	s.DynamicColumns[header] = DynamicFunc{
		Header: header,
		Func:   fn,
		Cols:   cols,
	}
	return s
}

func (s *Selector) ID() string {
	if s.Id == "" {
		s.Id = "default"
	}
	return s.Id
}

// DynamicFunc
type DynamicFunc struct {
	Header string        `json:"header" yaml:"header" toml:"header" xml:"header" ini:"header"`
	Name   string        `json:"func_name" yaml:"func_name" toml:"func_name" xml:"funcName" ini:"funcName"`
	Func   DynamicColumn `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
	Cols   []int         `json:"cols" yaml:"cols" toml:"cols" xml:"cols" ini:"cols"`
}

// Printer
type Printer struct {
	Colorize bool   `default:"true" json:"colorize" yaml:"colorize" toml:"colorize" xml:"colorize" ini:"colorize"`
	Format   string `default:"tabular-grid" json:"format,omitempty" yaml:"format,omitempty" toml:"format,omitempty" xml:"format,omitempty" ini:"format,omitempty"`
}

// Format
type Format struct {
	Extension string
	MimeType  string
}

type DataQL struct {
	Query string `json:"query,omitempty" yaml:"query,omitempty" toml:"query,omitempty" xml:"query,omitempty" ini:"query,omitempty"`
}

// Slicer
type Slicer struct {
	Headers []string `json:"headers,omitempty" yaml:"headers,omitempty" toml:"headers,omitempty" xml:"headers,omitempty" ini:"headers,omitempty"`
	Cols    *Ranger  `json:"columns,omitempty" yaml:"columns,omitempty" toml:"columns,omitempty" xml:"columns,omitempty" ini:"columns,omitempty"`
	Rows    *Ranger  `json:"rows,omitempty" yaml:"rows,omitempty" toml:"rows,omitempty" xml:"rows,omitempty" ini:"rows,omitempty"`
	// Cols    string   `json:"columns,omitempty" yaml:"columns,omitempty" toml:"columns,omitempty" xml:"columns,omitempty" ini:"columns,omitempty"`
	// Rows    string   `json:"rows,omitempty" yaml:"rows,omitempty" toml:"rows,omitempty" xml:"rows,omitempty" ini:"rows,omitempty"`
}

type Ranger struct {
	Expr  string `json:"expr,omitempty" yaml:"expr,omitempty" toml:"expr,omitempty" xml:"expr,omitempty" ini:"expr,omitempty"`
	Lower int    `json:"lower,omitempty" yaml:"lower,omitempty" toml:"lower,omitempty" xml:"lower,omitempty" ini:"lower,omitempty"`
	Upper int    `json:"upper,omitempty" yaml:"upper,omitempty" toml:"upper,omitempty" xml:"upper,omitempty" ini:"upper,omitempty"`
	Cap   int    `json:"cap,omitempty" yaml:"cap,omitempty" toml:"cap,omitempty" xml:"cap,omitempty" ini:"cap,omitempty"`
}

// Writer
type Writer struct {
	Split      bool
	SplitAt    int
	Concurrent bool
	PrefixPath string
	Basename   string
	Formats    []string
}

// type Plucker struct {}
