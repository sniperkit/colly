package config

type DatasetConfig struct {
	Disabled     bool             `default:"false" json:"enabled" yaml:"enabled" toml:"enabled" xml:"enabled" ini:"enabled"`
	IsExportable bool             `default:"true" json:"is_exportable" yaml:"is_exportable" toml:"is_exportable" xml:"is_exportable" ini:"is_exportable"`
	Datasets     []*DatasetConfig `json:"datasets" yaml:"datasets" toml:"datasets" xml:"datasets" ini:"datasets"`
	MaxRows      int              `default:"100000" json:"max_rows" yaml:"max_rows" toml:"max_rows" xml:"max_rows" ini:"max_rows"`
	MaxCols      int              `default:"100" json:"max_cols" yaml:"max_cols" toml:"max_cols" xml:"max_cols" ini:"max_cols"`
	Charset      string           `default:"UTF-8" json:"charset" yaml:"charset" toml:"charset" xml:"charset" ini:"charset"`
	Exports      []*ExportConfig  `json:"exports" yaml:"exports" toml:"exports" xml:"exports" ini:"exports"`
	err          error            `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`
}
