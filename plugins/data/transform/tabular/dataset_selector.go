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
	Slug           string                 `json:"slug" yaml:"slug" toml:"slug" xml:"slug" ini:"slug"`
	HasPrefix      string                 `json:"has_prefix" yaml:"has_prefix" toml:"has_prefix" xml:"hasPrefix" ini:"hasPrefix"`
	HasSuffix      string                 `json:"has_suffixslug" yaml:"has_suffix" toml:"has_suffix" xml:"hasSuffix" ini:"hasSuffix"`
	PatternURL     string                 `json:"pattern_url" yaml:"pattern_url" toml:"pattern_url" xml:"patternURL" ini:"patternURL"`
	PatternRegex   *regexp.Regexp         `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
	DynamicColumns map[string]DynamicFunc `json:"dynamic_columns" yaml:"dynamic_columns" toml:"dynamic_columns" xml:"dynamicColumns" ini:"dynamicColumns"`
	Headers        []string               `required:"true" json:"headers" yaml:"headers" toml:"headers" xml:"headers" ini:"headers"`
	Mixed          []interface{}          `json:"mixed" yaml:"mixed" toml:"mixed" xml:"mixed" ini:"mixed"`
	DataQL         *DataQL                `json:"data_ql" yaml:"data_ql" toml:"data_ql" xml:"dataQL" ini:"dataQL"`
	Slicer         *Slicer                `json:"slicer" yaml:"slicer" toml:"slicer" xml:"slicer" ini:"slicer"`
	Writer         *Writer                `json:"writer" yaml:"writer" toml:"writer" xml:"writer" ini:"writer"`
	Printer        *Printer               `json:"printer" yaml:"printer" toml:"printer" xml:"printer" ini:"printer"`

	// not exported
	errs []error `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`

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
	Func   DynamicColumn `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-" csv:"-"`
	Cols   []int         `json:"cols" yaml:"cols" toml:"cols" xml:"cols" ini:"cols"`
}

// Printer
type Printer struct {
	Colorize bool   `default:"true" json:"colorize" yaml:"colorize" toml:"colorize" xml:"colorize" ini:"colorize"`
	Format   string `default:"tabular-grid" json:"format" yaml:"format" toml:"format" xml:"format" ini:"format"`
}

// Format
type Format struct {
	Extension string `json:"extension" yaml:"extension" toml:"extension" xml:"extension" ini:"extension"`
	MimeType  string `json:"mime_type" yaml:"mime_type" toml:"mime_type" xml:"mimeType" ini:"mimeType"`
}

type DataQL struct {
	Query string `json:"query" yaml:"query" toml:"query" xml:"query" ini:"query"`
	DSL   string `json:"dsl" yaml:"dsl" toml:"dsl" xml:"dsl" ini:"dsl"`
}

// Slicer
type Slicer struct {
	Headers []string `json:"headers" yaml:"headers" toml:"headers" xml:"headers" ini:"headers"`
	Cols    *Ranger  `json:"columns" yaml:"columns" toml:"columns" xml:"columns" ini:"columns"`
	Rows    *Ranger  `json:"rows" yaml:"rows" toml:"rows" xml:"rows" ini:"rows"`
	// Cols    string   `json:"columns" yaml:"columns" toml:"columns" xml:"columns" ini:"columns"`
	// Rows    string   `json:"rows" yaml:"rows" toml:"rows" xml:"rows" ini:"rows"`
}

type Ranger struct {
	Expr  string `json:"expr" yaml:"expr" toml:"expr" xml:"expr" ini:"expr"`
	Lower int    `json:"lower" yaml:"lower" toml:"lower" xml:"lower" ini:"lower"`
	Upper int    `json:"upper" yaml:"upper" toml:"upper" xml:"upper" ini:"upper"`
	Cap   int    `json:"cap" yaml:"cap" toml:"cap" xml:"cap" ini:"cap"`
}

// Writer
type Writer struct {
	Split      bool     `json:"split" yaml:"split" toml:"split" xml:"split" ini:"split"`
	SplitAt    int      `json:"split_at" yaml:"split_at" toml:"split_at" xml:"splitAt" ini:"splitAt"`
	Concurrent bool     `json:"concurrent" yaml:"concurrent" toml:"concurrent" xml:"concurrent" ini:"concurrent"`
	PrefixPath string   `json:"prefix_path" yaml:"prefix_path" toml:"prefix_path" xml:"prefixPath" ini:"prefixPath"`
	Basename   string   `json:"basename" yaml:"basename" toml:"basename" xml:"basename" ini:"basename"`
	Formats    []string `json:"formats" yaml:"formats" toml:"formats" xml:"formats" ini:"formats"`
}

// type Plucker struct {}
