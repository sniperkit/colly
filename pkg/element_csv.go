package colly

import (
	"github.com/sniperkit/textql/pkg/inputs"
	"github.com/sniperkit/textql/pkg/outputs"
	// "github.com/sniperkit/textql/pkg/storage"
	// "github.com/sniperkit/trdsql/pkg"
)

// `CSV` is the key for the csv encoding
const CSV = "csv"

// `TSV` is the key for the tsv encoding
const TSV = "tsv"

// `FORMAT_TABLE` specifies the
type FORMAT_TABLE string

const (

	// `TABLE_CSV`
	TABLE_CSV FORMAT_TABLE = "csv"

	// `TABLE_CSV`
	TABLE_TSV FORMAT_TABLE = "tsv"

	//-- End
)

var (

	// inputOpts
	inputOpts *inputs.CSVInputOptions

	// displayOpts
	displayOpts *outputs.CSVOutputOptions

	// storageSqlite (not implemented as we will use either pivot or gorm as abstraction data layer)
	// storageSqlite *storage.SQLite3Storage
)
