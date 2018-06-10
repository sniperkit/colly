package tablib

import (
	"strings"

	// internal
	jsonpointer "github.com/sniperkit/colly/plugins/data/transform/tabular/jsonpointer"
)

type pointers []jsonpointer.JSONPointer

func (pts pointers) Len() int      { return len(pts) }
func (pts pointers) Swap(i, j int) { pts[i], pts[j] = pts[j], pts[i] }
func (pts pointers) Less(i, j int) bool {
	// shallow path first
	if pts[i].Len() != pts[j].Len() {
		return pts[i].Len() < pts[j].Len()
	}

	// compare each part
	for n := 0; n < pts[i].Len(); n++ {
		if pts[i][n] != pts[j][n] {
			return pts[i][n] < pts[j][n]
		}
	}
	return false
}

func (pts pointers) Strings() []string {
	keys := make([]string, 0, pts.Len())
	for _, p := range pts {
		keys = append(keys, p.String())
	}
	return keys
}

func (pts pointers) Slashes() []string {
	keys := make([]string, 0, pts.Len())
	for _, p := range pts {
		keys = append(keys, strings.Join(p.Strings(), "/"))
	}
	return keys
}

func (pts pointers) DotNotations(bracketIndex bool) []string {
	keys := make([]string, 0, pts.Len())
	for _, p := range pts {
		keys = append(keys, p.DotNotation(bracketIndex))
	}
	return keys
}

func (pts pointers) Underscores() []string {
	keys := make([]string, 0, pts.Len())
	for _, p := range pts {
		keys = append(keys, strings.Join(p.Strings(), "_"))
	}
	return keys
}

func allPointers(results []KeyValue) (pointers pointers, err error) {
	set := make(map[string]bool, 0)
	for _, result := range results {
		for _, key := range result.Keys() {
			if !set[key] {
				set[key] = true
				pointer, err := jsonpointer.New(key)
				if err != nil {
					return nil, err
				}
				pointers = append(pointers, pointer)
			}
		}
	}
	return
}