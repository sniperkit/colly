package main

import (
	tablib "github.com/sniperkit/colly/plugins/data/transform/tabular"
)

// data structure - databooks and datasets
var (
	sheets   map[string][]interface{} = make(map[string][]interface{}, 0)
	dsExport *tablib.Dataset
	dsURLs   *tablib.Dataset
	dataBook *tablib.Databook
	clm      = cmmap.NewConcurrentMultiMap()
)
