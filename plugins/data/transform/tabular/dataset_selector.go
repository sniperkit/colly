package tablib

import (
	"errors"
	"regexp"
	// "sync"
)

type Selector struct {
	// exported
	Id         string        `json:"identifier" yaml:"identifier" toml:"identifier" xml:"identifier" ini:"identifier"`
	PatternURL string        `json:"pattern_url" yaml:"pattern_url" toml:"pattern_url" xml:"patternURL" ini:"patternURL"`
	Headers    []string      `required:"true" json:"headers" yaml:"headers" toml:"headers" xml:"headers" ini:"headers"`
	Mixed      []interface{} `json:"mixed,omitempty" yaml:"mixed,omitempty" toml:"mixed,omitempty" xml:"mixed,omitempty" ini:"mixed,omitempty"`
	Slicer     *Slicer       `json:"slicer" yaml:"slicer" toml:"slicer" xml:"slicer" ini:"slicer"`
	Writer     *Writer       `json:"writer,omitempty" yaml:"writer,omitempty" toml:"writer,omitempty" xml:"writer,omitempty" ini:"writer,omitempty"`
	// not exported
	patternRegex *regexp.Regexp
	errs         []error
}

// PatternURL
func (s *Selector) PatternRegexp(rule string) *Selector {
	if rule == "" {
		s.errs = append(s.errs, errors.New("Empty pattern regex, skipping..."))
		return s
	}
	s.PatternURL = rule
	s.patternRegex = regexp.MustCompile(rule)
	return s
}

func (s *Selector) ID() string {
	if s.Id == "" {
		s.Id = "default"
	}
	return s.Id
}

// PluckURL
// func (s *Selector) PluckURL(rule string) *Selector {
// 	s.patternRegex = regexp.MustCompile(rule)
//	return s
//}

// Slicer
type Slicer struct {
	Cols string `json:"columns,omitempty" yaml:"columns,omitempty" toml:"columns,omitempty" xml:"columns,omitempty" ini:"columns,omitempty"`
	Rows string `json:"rows,omitempty" yaml:"rows,omitempty" toml:"rows,omitempty" xml:"rows,omitempty" ini:"rows,omitempty"`
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
