package configor

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	// crypto
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"

	// internal
	// xdgbasedir
	xdgbasedir "github.com/sniperkit/colly/plugins/system/xdgbasedir"
)

/*
	TODO:
	- test/implement these packages
	  - github.com/sniperkit/gographer
	  - github.com/sniperkit/struct2schema
*/

// public variables
var (
	DefaultEnvFiles     = []string{".env"}
	ErrConfigFileDecode = errors.New("failed to decode config")
	ErrConfigFileEncode = errors.New("failed to encode config")
	ErrConfigFileDump   = errors.New("failed to dump loaded config")
)

// private variables
var (
	xdgBaseDir string = "."
	envKeys    map[string]string
)

// constant
const (
	DEFAULT_BASE_DIR string = "."
)

func getBaseDirectory() (string, error) {
	xdgBaseDir, err := xdgbasedir.ConfigHomeDirectory()
	if err != nil {
		return DEFAULT_BASE_DIR, err
	}
	return xdgBaseDir, nil
}

func (configor *Configor) ExportToFile(filePath string) error {
	xdgBaseDir, err := configor.XDGBaseDir()
	if err != nil {
		return err
	}
	fmt.Printf("Default XDG Base Dir '%v'...\n", xdgBaseDir)
	fmt.Printf("Export to file '%s'...\n", filePath)
	return nil
}

func (configor *Configor) ExportTo(prefixPath string, formats ...string) error {
	xdgBaseDir, err := configor.XDGBaseDir()
	if err != nil {
		return err
	}
	fmt.Printf("Default XDG Base Dir '%v'...\n", xdgBaseDir)
	fmt.Printf("Export formats '%s'...\n", strings.Join(formats, ","))
	fmt.Printf("Export to path '%s'...\n", prefixPath)
	return nil
}

func (configor *Configor) XDGBaseDir() (string, error) {
	xdgBaseDir, err := getBaseDirectory()
	if err != nil {
		return DEFAULT_BASE_DIR, err
	}
	if configor.Config.Debug || configor.Config.Verbose {
		fmt.Printf("Default XDG Base Dir '%v'...\n", xdgBaseDir)
	}
	configor.Config.XDGBaseDir = xdgBaseDir
	return xdgBaseDir, nil
}

func XDGBaseDir() (string, error) {
	xdgBaseDir, err := getBaseDirectory()
	if err != nil {
		return DEFAULT_BASE_DIR, err
	}
	return xdgBaseDir, nil
}

func isEmptyStruct(object interface{}) bool {
	//First check normal definitions of empty
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}
	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}

func (configor *Configor) Dump(config interface{}, nodes []string, formats []string, prefixPath string) error {
	err := os.MkdirAll(prefixPath, 0700)
	if err != nil {
		return err
	}
	if config == nil {
		config = &Config{}
	}

	// var inspectConfig *Debug
	if InspectMode {
		// configor.debug = &Debug{}
		configor.debug = inspectStruct(config)
	}

	if nodes[0] == "all" {
		nodes = []string{}
	}

	exportNodesCount := len(nodes)
	for _, f := range formats {
		switch {
		case exportNodesCount == 0:
			nodeName := "global"
			data, err := encodeFile(config, nodeName, f)
			if err != nil {
				return err
			}
			filePath := getConfigDumpFilePath(prefixPath, nodeName, f)
			if err := writeFile(filePath, data); err != nil {
				return err
			}

			if configor.Inspect {
				nodeName := "inspect"
				data, err := encodeFile(configor.debug, nodeName, f)
				if err != nil {
					return err
				}
				filePath := getConfigDumpFilePath(prefixPath, nodeName, f)
				if err := writeFile(filePath, data); err != nil {
					return err
				}
			}

		case exportNodesCount > 0:
			for _, n := range nodes {
				data, err := encodeFile(config, n, f)
				if err != nil {
					return err
				}
				filePath := getConfigDumpFilePath(prefixPath, n, f)
				if err := writeFile(filePath, data); err != nil {
					return err
				}
			}
		}
	}
	return nil
	//return errors.New("error occured while selecting the node to export")
}

// nodes := []string{"contacts", "db", "oauth2"}
// configor.Dump(Config, "yaml", "contacts", "db", "oauth2")
func Dump(config interface{}, nodes []string, formats []string, prefixPath string) error {
	err := os.MkdirAll(prefixPath, 0700)
	if err != nil {
		return err
	}
	if config == nil {
		config = &Config{}
	}

	fmt.Printf("InspectMode: %t, prefixPath=%s \n", InspectMode, prefixPath)

	// TODO: create a method for dumping the config struct inspection
	var inspectConfig *Debug
	if InspectMode {
		inspectConfig = &Debug{}

		// Load(config, []string{})
		inspectConfig = inspectStruct(config)
		// prettyPrinter("inspect", inspectConfig)
		for _, f := range formats {
			nodeName := "inspect"
			fmt.Println("format=", f, "nodeName=", nodeName)
			data, err := encodeFile(inspectConfig, nodeName, f)
			if err != nil {
				fmt.Println("error, msg=", err)
				os.Exit(1)
				return err
			}
			filePath := getConfigDumpFilePath(prefixPath, nodeName, f)
			if err := writeFile(filePath, data); err != nil {
				fmt.Println("error, msg=", err)
				os.Exit(1)
				return err
			}
		}

	}

	if nodes[0] == "all" {
		nodes = []string{}
	}

	exportNodesCount := len(nodes)
	for _, f := range formats {
		switch {
		case exportNodesCount == 0:
			nodeName := "config"
			data, err := encodeFile(config, nodeName, f)
			if err != nil {
				fmt.Println("error, msg=", err)
				os.Exit(1)
				return err
			}
			filePath := getConfigDumpFilePath(prefixPath, nodeName, f)
			if err := writeFile(filePath, data); err != nil {
				fmt.Println("error, msg=", err)
				os.Exit(1)
				return err
			}

		case exportNodesCount > 0:
			for _, n := range nodes {
				data, err := encodeFile(config, n, f)
				if err != nil {
					fmt.Println("error, msg=", err)
					os.Exit(1)
					return err
				}
				filePath := getConfigDumpFilePath(prefixPath, n, f)
				if err := writeFile(filePath, data); err != nil {
					fmt.Println("error, msg=", err)
					os.Exit(1)
					return err
				}
			}
		}
	}
	return nil
	//return errors.New("error occured while selecting the node to export")
}

func getConfigDumpFilePath(prefixPath string, nodeName string, format string) string {
	return fmt.Sprintf("%s/%s.%s", prefixPath, nodeName, format)
}

func getAttributesListToExport(attrs string) []string {
	return strings.Split(attrs, ",")
}

func writeFile(filePath string, data []byte) error {
	if err := ioutil.WriteFile(filePath, data, 0600); err != nil {
		return err
	}
	return nil
}

func Md5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func Hmac(k, s string) string {
	h := hmac.New(sha256.New, []byte(k))
	h.Write([]byte(s))
	return string(hex.EncodeToString(h.Sum(nil)))
}
