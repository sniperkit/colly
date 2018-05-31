package config

import (
	"fmt"
	"os"
)

type Configor struct {
	*Setting
}

type Setting struct {
	Environment string
	ENVPrefix   string
	Debug       bool
	Verbose     bool

	// Supported only for toml and yaml files.
	// json does not currently support this: https://github.com/golang/go/issues/15314
	// This setting will be ignored for json files.
	ErrorOnUnmatchedKeys bool
}

// New initialize a Configor
func New(config *Setting) *Configor {
	if config == nil {
		config = &Setting{}
	}

	if os.Getenv("CONFIGOR_DEBUG_MODE") != "" {
		config.Debug = true
	}

	if os.Getenv("CONFIGOR_VERBOSE_MODE") != "" {
		config.Verbose = true
	}

	return &Configor{Setting: config}
}

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
		if configor.Setting.Debug || configor.Setting.Verbose {
			fmt.Printf("Configuration:\n  %#v\n", config)
		}
	}()

	for _, file := range configor.getConfigurationFiles(files...) {
		if configor.Setting.Debug || configor.Setting.Verbose {
			fmt.Printf("Loading configurations from file '%v'...\n", file)
		}
		if err := processFile(config, file, configor.GetErrorOnUnmatchedKeys()); err != nil {
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
