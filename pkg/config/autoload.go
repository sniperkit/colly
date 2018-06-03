package config

import (
	colly "github.com/sniperkit/colly"

	cfgParser "github.com/sniperkit/colly/plugins/cmd/config"
)

func NewCollector(async bool, randomUserAgent bool, isDebugMode bool) *colly.Collector {
	c := &colly.Collector{}
= randomUserAgent
	c.DebugMode = isDebugMode
	return c
}

func NewCollectorWithConfig(cfg *CollectorConfig) *colly.Collector {
	c := &colly.Collector{}
	return c
}

func ExportConfig(formats []string) map[string]bool {
	exportResults = make(map[string]bool, 0)
	return exportResults
}

func (c *CollectorConfig) SetConfigExports(withSchema bool, withAutoLoad bool) *CollectorConfig {
	c.AllowExportConfigSchema = withSchema
	c.AllowExportConfigAutoload = withAutoLoad
	return c
}

// userAgent "",

func (c *CollectorConfig) ToFile(formats []string) (ok bool, err error) {
	return
}

func (c *CollectorConfig) ToBytes(formats []string) (output []byte) {
	return
}

func (c *CollectorConfig) ToString(formats []string) (output string) {
	return
}

func (c *CollectorConfig) Reset() (ok bool) {
	return
}

func (c *CollectorConfig) Clone() *CollectorConfig {
	return
}
