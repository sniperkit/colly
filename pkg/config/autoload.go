package config

/*
import (
	cfgParser "github.com/sniperkit/colly/plugins/cmd/config"
)
*/

func DumpConfig(onlySchema bool, formats []string) map[string]bool {
	exportResults := make(map[string]bool, 0)
	return exportResults
}

func (c *CollectorConfig) SetConfigExports(withSchema bool, withAutoLoad bool) *CollectorConfig {
	c.AllowExportConfigSchema = withSchema
	c.AllowExportConfigAutoload = withAutoLoad
	return c
}

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
	return c
}
