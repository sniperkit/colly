package configor

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"

	// crypto
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"

	// encoding
	hex "encoding/hex"
	// xml parser/iterator
	xml "encoding/xml"
	// json
	json "encoding/json"
	// env vars iterator
	dotenv "github.com/sniperkit/colly/plugins/system/dotenv"
	// json parser/iterator
	// jsoniter "github.com/json-iterator/go"
	// yaml parser/iterator
	yaml "gopkg.in/yaml.v2"
	// toml parser/iterator
	toml "github.com/BurntSushi/toml"
	// ini parser/iterator
	ini "github.com/go-ini/ini"
	// xdgbasedir
	xdgbasedir "github.com/sniperkit/colly/plugins/system/xdgbasedir"
	// helpers
	pp "github.com/sniperkit/xutil/plugin/debug/pp"
)

// private variables
var (
	xdgBaseDir string = "."
	envKeys    map[string]string
	// json       = jsoniter.ConfigCompatibleWithStandardLibrary
	testRegexp = regexp.MustCompile("_test|(\\.test$)")
)

// public variables
var (
	DefaultEnvFiles     = []string{".env"}
	ErrConfigFileDecode = errors.New("failed to decode config")
	ErrConfigFileEncode = errors.New("failed to encode config")
	ErrConfigFileDump   = errors.New("failed to dump loaded config")
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
	configor.Config.xdgBaseDir = xdgBaseDir
	return xdgBaseDir, nil
}

func XDGBaseDir() (string, error) {
	xdgBaseDir, err := getBaseDirectory()
	if err != nil {
		return DEFAULT_BASE_DIR, err
	}
	return xdgBaseDir, nil
}

func (configor *Configor) WithBaseDir(xdgBaseDir string) *Configor {
	configor.Config.xdgBaseDir = xdgBaseDir
	return configor
}

// func (configor *Configor) WithWorkDir(workdir string) *Configor {}

/*
	Refs:
	- "github.com/zieckey/goini"
	- "github.com/go-ini/ini"
	- "github.com/vaughan0/go-ini"
	- "github.com/knq/ini"
*/

// ENV return environment
func ENV() string {
	return New(nil).GetEnvironment()
}

// Load will unmarshal configurations to struct from files that you provide
func Load(config interface{}, files ...string) error {
	return New(nil).Load(config, files...)
}

// UnmatchedTomlKeysError errors are returned by the Load function when
// ErrorOnUnmatchedKeys is set to true and there are unmatched keys in the input
// toml config file. The string returned by Error() contains the names of the
// missing keys.
type UnmatchedTomlKeysError struct {
	Keys []toml.Key
}

func (e *UnmatchedTomlKeysError) Error() string {
	return fmt.Sprintf("There are keys in the config file that do not match any field in the given struct: %v", e.Keys)
}

func getConfigurationFileWithENVPrefix(file, env string) (string, error) {
	var (
		envFile string
		extname = path.Ext(file)
	)

	if extname == "" {
		envFile = fmt.Sprintf("%v.%v", file, env)
	} else {
		envFile = fmt.Sprintf("%v.%v%v", strings.TrimSuffix(file, extname), env, extname)
	}

	if fileInfo, err := os.Stat(envFile); err == nil && fileInfo.Mode().IsRegular() {
		return envFile, nil
	}
	return "", fmt.Errorf("failed to find file %v", file)
}

func processFile(config interface{}, file string, errorOnUnmatchedKeys bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := dotenv.Load(); err != nil {
		return err
	}

	// replace KeyHolders before loading the YAML/JSON/TOML file
	envKeys, _ = dotenv.Read(DefaultEnvFiles...)
	dataStr := string(data)
	for k, v := range envKeys {
		holderKey := fmt.Sprintf("{ENV.%s}", strings.Replace(k, "\"", "", -1))
		dataStr = strings.Replace(dataStr, holderKey, v, -1)
	}
	data = []byte(dataStr)

	switch {
	case strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml"):
		if errorOnUnmatchedKeys {
			return yaml.UnmarshalStrict(data, config)
		}
		return yaml.Unmarshal(data, config)

	case strings.HasSuffix(file, ".ini"):
		var err error
		config, err = ini.Load(data) // ref. https://github.com/go-ini/ini
		if err != nil {
			return err
		}
		return nil

	case strings.HasSuffix(file, ".toml"):
		return unmarshalToml(data, config, errorOnUnmatchedKeys)

	case strings.HasSuffix(file, ".json"):
		return json.Unmarshal(data, config)

	case strings.HasSuffix(file, ".xml"):
		return xml.Unmarshal(data, config)

	default:

		if err := unmarshalToml(data, config, errorOnUnmatchedKeys); err == nil {
			return nil
		} else if errUnmatchedKeys, ok := err.(*UnmatchedTomlKeysError); ok {
			return errUnmatchedKeys
		}

		if json.Unmarshal(data, config) == nil {
			return nil
		}

		var yamlError error
		if errorOnUnmatchedKeys {
			yamlError = yaml.UnmarshalStrict(data, config)
		} else {
			yamlError = yaml.Unmarshal(data, config)
		}

		if yamlError == nil {
			return nil
		} else if yErr, ok := yamlError.(*yaml.TypeError); ok {
			return yErr
		}

		return errors.New("failed to decode config")
	}
}

func (configor *Configor) processTags(config interface{}, prefixes ...string) error {
	configValue := reflect.Indirect(reflect.ValueOf(config))
	if configValue.Kind() != reflect.Struct {
		return errors.New("invalid config, should be struct")
	}

	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		var (
			envNames    []string
			fieldStruct = configType.Field(i)
			field       = configValue.Field(i)
			envName     = fieldStruct.Tag.Get("env") // read configuration from shell env
		)

		if !field.CanAddr() || !field.CanInterface() {
			continue
		}

		if envName == "" {
			envNames = append(envNames, strings.Join(append(prefixes, fieldStruct.Name), "_"))                  // Configor_DB_Name
			envNames = append(envNames, strings.ToUpper(strings.Join(append(prefixes, fieldStruct.Name), "_"))) // CONFIGOR_DB_NAME
		} else {
			envNames = []string{envName}
		}

		if configor.Config.Verbose {
			fmt.Printf("Trying to load struct `%v`'s field `%v` from env %v\n", configType.Name(), fieldStruct.Name, strings.Join(envNames, ", "))
		}

		// Load From Shell ENV
		for _, env := range envNames {
			if value := os.Getenv(env); value != "" {
				if configor.Config.Debug || configor.Config.Verbose {
					fmt.Printf("Loading configuration for struct `%v`'s field `%v` from env %v...\n", configType.Name(), fieldStruct.Name, env)
				}
				if err := yaml.Unmarshal([]byte(value), field.Addr().Interface()); err != nil {
					return err
				}
				break
			}
		}

		if isBlank := reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()); isBlank {
			// Set default configuration if blank
			if value := fieldStruct.Tag.Get("default"); value != "" {
				if err := yaml.Unmarshal([]byte(value), field.Addr().Interface()); err != nil {
					return err
				}
			} else if fieldStruct.Tag.Get("required") == "true" {
				// return error if it is required but blank
				return errors.New(fieldStruct.Name + " is required, but blank")
			}
		}

		for field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			if err := configor.processTags(field.Addr().Interface(), getPrefixForStruct(prefixes, &fieldStruct)...); err != nil {
				return err
			}
		}

		if field.Kind() == reflect.Slice {
			for i := 0; i < field.Len(); i++ {
				if reflect.Indirect(field.Index(i)).Kind() == reflect.Struct {
					if err := configor.processTags(field.Index(i).Addr().Interface(), append(getPrefixForStruct(prefixes, &fieldStruct), fmt.Sprint(i))...); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// GetStringTomlKeys returns a string array of the names of the keys that are passed in as args
func GetStringTomlKeys(list []toml.Key) []string {
	arr := make([]string, len(list))

	for index, key := range list {
		arr[index] = key.String()
	}
	return arr
}

func unmarshalToml(data []byte, config interface{}, errorOnUnmatchedKeys bool) error {
	metadata, err := toml.Decode(string(data), config)
	if err == nil && len(metadata.Undecoded()) > 0 && errorOnUnmatchedKeys {
		return &UnmatchedTomlKeysError{Keys: metadata.Undecoded()}
	}
	return err
}

func getPrefixForStruct(prefixes []string, fieldStruct *reflect.StructField) []string {
	if fieldStruct.Anonymous && fieldStruct.Tag.Get("anonymous") == "true" {
		return prefixes
	}
	return append(prefixes, fieldStruct.Name)
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

func encodeFile(config interface{}, node string, format string) ([]byte, error) {
	switch format {
	case "ini":

		pp.Println("encodeInterfaceToINI=", config)
		err := ini.MapTo(config, "./colly.ini")
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
			return nil, err
		}

		// cfg := ini.Empty()

		// cfg.WriteToIndent(writer, "\t")
		// err = cfg.SaveToIndent("my.ini", "\t")
		/*
			var dataBytes bytes.Buffer
			// dataBytes := bytes.NewBuffer(nil)
			// configValue := reflect.ValueOf(config)
			outFile := ini.Empty()
			configValue := reflect.Indirect(reflect.ValueOf(*config))
			configValue2 := reflect.ValueOf(config)
			pp.Println(config)
			pp.Println(configValue)
			pp.Println(configValue2)
			err := ini.ReflectFrom(outFile, config)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			output, err := outFile.WriteTo(&dataBytes)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			fmt.Println(output)
			return []byte(dataBytes.String()), nil
		*/
	case "json":
		data, err := json.MarshalIndent(config, "", "\t")
		if err != nil {
			return nil, err
		}
		return data, nil

	case "toml":
		var dataBytes bytes.Buffer
		if err := toml.NewEncoder(&dataBytes).Encode(config); err != nil {
			return nil, err
		}
		// fmt.Println(dataBytes.String())
		return []byte(dataBytes.String()), nil

	case "yaml":
		data, err := yaml.Marshal(config)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New("Unkown format to export")

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

	exportNodesCount := len(nodes)
	for _, f := range formats {
		switch {
		case exportNodesCount == 0:
			nodeName := "config"
			data, err := encodeFile(config, "config", f)
			if err != nil {
				return err
			}
			filePath := getConfigDumpFilePath(prefixPath, nodeName, f)
			// fmt.Printf("filePath: %s \n", filePath)
			if err := writeFile(filePath, data); err != nil {
				return err
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
