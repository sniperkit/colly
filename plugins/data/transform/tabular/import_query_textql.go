package tablib

import (
	"github.com/sniperkit/textql/pkg/inputs"
	"github.com/sniperkit/textql/pkg/outputs"
	"github.com/sniperkit/textql/pkg/storage"
)

var (
	inputOpts     *inputs.CSVInputOptions
	displayOpts   *outputs.CSVOutputOptions
	storageSqlite *storage.SQLite3Storage
)
