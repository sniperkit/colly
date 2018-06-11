package tablib

import (
	"reflect"
	"sort"
)

// KeyStyle represents the specific style of the key.
type KeyStyle uint

// Header style
const (
	// "/foo/bar/0/baz"
	JSONPointerStyle KeyStyle = iota

	// "foo/bar/0/baz"
	SlashStyle

	// "foo.bar.0.baz"
	DotNotationStyle

	// "foo.bar[0].baz"
	DotBracketStyle // KeyStyle = iota

	// "foo_bar_0_baz"
	UnderscoreStyle
)

type mapKeys []reflect.Value

func (k mapKeys) Len() int           { return len(k) }
func (k mapKeys) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k mapKeys) Less(i, j int) bool { return k[i].String() < k[j].String() }

func sortedMapKeys(v reflect.Value) []reflect.Value {
	var keys mapKeys = v.MapKeys()
	sort.Sort(keys)
	return keys
}

// KeyValue represents key(path)/value map.
type KeyValue map[string]interface{}

// Keys returns all keys.
func (kv KeyValue) Keys() []string {
	keys := make([]string, 0, len(kv))
	for k := range kv {
		keys = append(keys, k)
	}
	return keys
}
