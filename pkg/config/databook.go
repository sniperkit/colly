package config

type DatabookConfig struct {
	Disabled     bool             `default:'false' json:'disabled' yaml:'disabled' toml:'disabled' xml:'disabled' ini:'disabled'`
	IsExportable bool             `default:'true' json:'is_exportable' yaml:'is_exportable' toml:'is_exportable' xml:'is_exportable' ini:'is_exportable'`
	Datasets     []*DatasetConfig `json:'datasets' yaml:'datasets' toml:'datasets' xml:'datasets' ini:'datasets'`
	MaxDatasets  int              `default:'5' json:'max_datasets' yaml:'max_datasets' toml:'max_datasets' xml:'maxDatasets' ini:'maxDatasets'`
	Charset      string           `default:'UTF-8' json:'charset' yaml:'charset' toml:'charset' xml:'charset' ini:'charset'`
	Exports      []*ExportConfig  `json:'exports' yaml:'exports' toml:'exports' xml:'exports' ini:'exports'`
	err          error            `json:'-' yaml:'-' toml:'-' xml:'-' ini:'-'`
}
