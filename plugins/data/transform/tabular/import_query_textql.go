package tablib

import (
	"github.com/sniperkit/textql/inputs"
	"github.com/sniperkit/textql/outputs"
	"github.com/sniperkit/textql/storage"
)

var (
	inputOpts     *inputs.CSVInputOptions
	displayOpts   *outputs.CSVOutputOptions
	storageSqlite *storage.SQLite3Storage
)
