package tachymeter

import (
	"bytes"

	tabular "github.com/sniperkit/colly/addons/convert/tabular"
	jsoniter "github.com/sniperkit/xutil/plugin/format/json"
)

var (
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	dataSet  *tabular.Dataset
	dataBook *tabular.Databook
)

/*
	Refs:
	- https://github.com/agrison/go-tablib
	- github.com/sniperkit/colly/addons/convert/tabular
	- github.com/sniperkit/xtask/plugin
*/

type Export struct {
	Encoding   string `default:'tsv'`
	PrefixPath string `default:'./shared/exports/stats/tachymeter/'`
	Basename   string `default:'tachymeter_export'`
	SplitLimit int    `default:'2500'`
	BufferSize int    `default:'20000'`
	BackupMode bool   `default:'true'`
	Overwrite  bool   `default:'true'`
	buffer     *bytes.Buffer
	dataset    *tabular.Dataset
	databook   *tabular.Databook
}