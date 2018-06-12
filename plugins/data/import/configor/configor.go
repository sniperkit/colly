package configor

import (
	"fmt"
	"os"
	"regexp"
)

type Configor struct {
	*Config
}

type Config struct {

	// Environment specifies...
	Environment string `json:"env" yaml:"env" toml:"env" xml:"env" ini:"env"`

	// ENVPrefix specifies...
	ENVPrefix string `default:"CONFIGOR_;XCONF_;XCONFIG_" json:"env_prefix" yaml:"env_prefix" toml:"env_prefix" xml:"envPrefix" ini:"envPrefix"`

	// XDGBaseDir specifies...
	XDGBaseDir string `json:"xdg_base_dir" yaml:"xdg_base_dir" toml:"xdg_base_dir" xml:"xdgBaseDir" ini:"xdgBaseDir"`

	// Debug specifies...
	Debug bool `default:"false" json:"debug" yaml:"debug" toml:"debug" xml:"debug" ini:"debug"`

	// Verbose specifies...
	Verbose bool `default:"false" json:"verbose" yaml:"verbose" toml:"verbose" xml:"verbose" ini:"verbose"`

	// Validate specifies...
	Validate bool `default:"false" json:"validate" yaml:"validate" toml:"validate" xml:"validate" ini:"validate"`

	// Dump specifies...
	Dump bool `default:"false" json:"dump" yaml:"dump" toml:"dump" xml:"dump" ini:"dump"`

	// Lookup specifies...
	Lookup bool `default:"false" json:"lookup" yaml:"lookup" toml:"lookup" xml:"lookup" ini:"lookup"`

	// Paths specifies...
	Paths string `default:"./conf;../conf;./shared/conf;../shared/conf;" json:"paths" yaml:"paths" toml:"paths" xml:"paths" ini:"paths"`

	// Basenames specifies...
	Basenames string `default:"config,app,application,global" json:"basenames" yaml:"basenames" toml:"basenames" xml:"basenames" ini:"basenames"`

	// Keywords specifies...
	Keywords string `default:"xgd,xdg_base_dir,xdgBaseDir,auto_config,autoConfig" json:"keywords" yaml:"keywords" toml:"keywords" xml:"keywords" ini:"keywords"`

	// Extensions specifies...
	Extensions string `default:"yaml,yml,json,toml,xml,ini" json:"extensions" yaml:"extensions" toml:"extensions" xml:"extensions" ini:"extensions"`

	// ErrorOnUnmatchedKeys specifies...
	Inspect bool `default:"false" json:"inspect" yaml:"inspect" toml:"inspect" xml:"inspect" ini:"inspect"`

	// ErrorOnUnmatchedKeys specifies...
	// Supported only for toml and yaml files.
	// json does not currently support this: https://github.com/golang/go/issues/15314
	// This setting will be ignored for json files.
	ErrorOnUnmatchedKeys bool `default:"true" json:"strict_mode" yaml:"strict_mode" toml:"strict_mode" xml:"strict_mode" ini:"strict_mode"`

	//-- End
}

// New initialize a Configor
func New(config *Config) *Configor {
	if config == nil {
		config = &Config{}
	}

	if os.Getenv("CONFIGOR_DEBUG_MODE") != "" {
		config.Debug = true
	}

	if os.Getenv("CONFIGOR_VERBOSE_MODE") != "" {
		config.Verbose = true
	}

	if os.Getenv("CONFIGOR_INSPECT_MODE") != "" {
		config.Dump = true
	}

	if os.Getenv("CONFIGOR_DUMP_MODE") != "" {
		config.Inspect = true
	}

	return &Configor{Config: config}
}

var testRegexp = regexp.MustCompile("_test|(\\.test$)")

// GetEnvironment get environment
func (configor *Configor) GetEnvironment() string {
	if configor.Environment == "" {
		if env := os.Getenv("CONFIGOR_ENV"); env != "" {
			return env
		}

		if testRegexp.MatchString(os.Args[0]) {
			return "test"
		}

		return "development"
	}
	return configor.Environment
}

// GetErrorOnUnmatchedKeys returns a boolean indicating if an error should be
// thrown if there are keys in the config file that do not correspond to the
// config struct
func (configor *Configor) GetErrorOnUnmatchedKeys() bool {
	return configor.ErrorOnUnmatchedKeys
}

// Load will unmarshal configurations to struct from files that you provide
func (configor *Configor) Load(config interface{}, files ...string) error {
	defer func() {
		if configor.Config.Debug {
			prettyPrinter("Load", config)
		}
	}()

	for _, file := range configor.getConfigurationFiles(files...) {
		if configor.Config.Debug || configor.Config.Verbose {
			// prettyPrinter("Load", "Configurations from file(s)=", config)
			fmt.Printf("Loading configurations from file '%v'...\n", file)
		}
		if err := processFile(config, file, configor.GetErrorOnUnmatchedKeys()); err != nil {
			// prettyPrinter("Load", "Error=", err, "file:", file, "unmatchedKeys:", configor.GetErrorOnUnmatchedKeys())
			fmt.Println("error: ", err, "file:", file, "unmatchedKeys:", configor.GetErrorOnUnmatchedKeys())
			return err
		}
	}

	prefix := configor.getENVPrefix(config)
	if prefix == "-" {
		return configor.processTags(config)
	}
	return configor.processTags(config, prefix)
}

// ENV return environment
func ENV() string {
	return New(nil).GetEnvironment()
}

// Load will unmarshal configurations to struct from files that you provide
func Load(config interface{}, files ...string) error {
	return New(nil).Load(config, files...)
}
