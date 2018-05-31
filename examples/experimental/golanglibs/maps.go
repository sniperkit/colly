package main

import (
	// "github.com/fanliao/go-concurrentMap"

	tablib "github.com/sniperkit/xutil/plugin/format/convert/tabular"
	cmap "github.com/sniperkit/xutil/plugin/map/multi"
)

// concurrent maps, datasets and databooks defaults
var (
	sheets   map[string][]interface{}   = make(map[string][]interface{}, 0)
	datasets map[string]*tablib.Dataset = make(map[string]*tablib.Dataset, 0) // := NewDataset([]string{"firstName", "lastName"})
	cds                                 = cmap.NewConcurrentMultiMap()
)
