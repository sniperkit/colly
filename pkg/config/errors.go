package config

import (
	"errors"
	"fmt"
	"regexp"
)

var testRegexp = regexp.MustCompile("_test|(\\.test$)")

var (
	ErrConfigFileDecode = errors.New("failed to decode config")
	ErrConfigFileEncode = errors.New("failed to encode config")
	ErrConfigFileDump   = errors.New("failed to dump loaded config")
)

func (e *UnmatchedTomlKeysError) Error() string {
	return fmt.Sprintf("There are keys in the config file that do not match any field in the given struct: %v", e.Keys)
}
