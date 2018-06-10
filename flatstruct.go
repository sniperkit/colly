// Package flatstruct "flattenss" structure fields to a map, giving possibility to access structure field
// values using unique paths.
package flatstruct

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	DefaultSeparator = "."    // Default path separator.
	DefaultNameTag   = "json" // Default field tag key.
)

var (
	whitespacesRe = regexp.MustCompile("[^\\s]+")
)

// InvalidKindError is returned in case when non-compatible kind is encountered.
type InvalidKindError struct {
	Kind reflect.Kind
}

func (e InvalidKindError) Error() string {
	return fmt.Sprintf("Invalid kind: %s", e.Kind)
}

// AmbiguousPathError is returned in case of non-unique path.
type AmbiguousPathError struct {
	Path string
}

func (e AmbiguousPathError) Error() string {
	return fmt.Sprintf("Path is ambiguous: %s", e.Path)
}

// InvalidNameError is returned in case when invalid name is encountered.
type InvalidNameError struct {
	Name string
}

func (e InvalidNameError) Error() string {
	return fmt.Sprintf("Name is invalid: %s", e.Name)
}

// Value is a reflect.Value along with parsed field's tag.
type Value struct {
	Value reflect.Value
	Tags  map[string][]string
}

type FlatStruct struct {
	PathSeparator string // Path separator used to separate path elements.
	NameTag       string // Struct's field tag key used to override name.
}

// PathJoin joins any number of path elements into a single path using specified separator.
// All empty path strings are ignored.
func PathJoin(sep string, elem ...string) string {
	p := elem[:0]
	for _, v := range elem {
		if len(v) != 0 {
			p = append(p, v)
		}
	}
	return strings.Join(p, sep)
}

func trimLeftRight(s string, set string) string {
	return strings.TrimRight(strings.TrimLeft(s, set), set)
}

func tagsForStructField(f reflect.StructField) map[string][]string {
	r := make(map[string][]string)
	for _, s := range whitespacesRe.FindAllString(string(f.Tag), -1) {
		p := strings.Split(s, ":")
		if len(p) == 2 {
			r[p[0]] = strings.Split(trimLeftRight(p[1], "\""), ",")
		}
	}
	return r
}

func (fs *FlatStruct) flattenDispatch(v reflect.Value, p string, r map[string]Value, t map[string][]string) error {
	switch v.Kind() {
	case reflect.Interface:
		if err := fs.flattenDispatch(reflect.ValueOf(v.Interface()), p, r, t); err != nil {
			return err
		}
	case reflect.Map:
		if err := fs.flattenMap(v, p, r); err != nil {
			return err
		}
	case reflect.Struct:
		if err := fs.flattenStruct(v, p, r); err != nil {
			return err
		}
	default:
		r[p] = Value{
			Value: v,
			Tags:  t,
		}
	}
	return nil
}

func (fs *FlatStruct) flattenMap(v reflect.Value, p string, r map[string]Value) error {
	for _, k := range v.MapKeys() {
		if k.Kind() == reflect.String {
			n := k.String()
			if strings.Contains(n, fs.PathSeparator) {
				return InvalidNameError{n}
			}
			pp := PathJoin(fs.PathSeparator, p, n)
			if _, ok := r[pp]; ok {
				return AmbiguousPathError{pp}
			}
			if err := fs.flattenDispatch(v.MapIndex(k), pp, r, map[string][]string{}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (fs *FlatStruct) flattenStruct(v reflect.Value, p string, r map[string]Value) error {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := tagsForStructField(v.Type().Field(i))
		name := v.Type().Field(i).Name
		if n, ok := t[fs.NameTag]; ok {
			if len(n) > 0 {
				if n[0] == "-" {
					continue
				} else if strings.Contains(n[0], fs.PathSeparator) {
					return InvalidNameError{n[0]}
				}
			}
			if len(n[0]) > 0 {
				name = n[0]
			}
		}
		pp := PathJoin(fs.PathSeparator, p, name)
		if _, ok := r[pp]; ok {
			return AmbiguousPathError{pp}
		}
		if err := fs.flattenDispatch(f, pp, r, t); err != nil {
			return err
		}
	}
	return nil
}

// Creates map of Values from struct, map or mix of both. The Values can be accessed by an unique path constructed
// from nested field names.
// Field names can be overridden by field tags. Tag key is defined in FlatStruct's NameTag variable.
// As a special case, "-" omits field in the "flattening" process.
// Keep in mind, that overriding field names can form ambiguous paths, leading to error.
func (fs *FlatStruct) Flatten(i interface{}) (map[string]Value, error) {
	r := make(map[string]Value)
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Map:
		if err := fs.flattenMap(v, "", r); err != nil {
			return nil, err
		}
	case reflect.Struct:
		if err := fs.flattenStruct(v, "", r); err != nil {
			return nil, err
		}
	case reflect.Ptr:
		return fs.Flatten(v.Elem().Interface())
	default:
		return nil, InvalidKindError{v.Kind()}
	}
	return r, nil
}

// Creates FlatStruct object instance, with a default configuration.
func NewFlatStruct() *FlatStruct {
	return &FlatStruct{
		PathSeparator: DefaultSeparator,
		NameTag:       DefaultNameTag,
	}
}
