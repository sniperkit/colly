package colly

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"unicode"

	"github.com/oleiade/reflections"

	// debug - inspect
	// pp "github.com/sniperkit/colly/plugins/app/debug/pp"

	// iterators
	json "encoding/json"
	xml "encoding/xml"
	toml "github.com/BurntSushi/toml"
	ini "github.com/go-ini/ini"
	yaml "github.com/sniperkit/yaml"
)

/*
	Dev-Only:
	- Used to dump the collector loaded configuration
	- Export by section/attributes
	- Export to formats: json, yaml, toml, xml and ini
*/
func inArray(e string, s []string) bool {
	for _, a := range s {
		fmt.Println("inArray.match=", e, "val", a)
		if a == e {
			return true
		}
	}
	return false
}

// ucFirst function
func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// properTitle function
func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "
	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

// field function
func field(t interface{}, key string) reflect.Value {
	strs := strings.Split(key, ".")
	v := reflect.ValueOf(t)
	for _, s := range strs[1:] {
		v = v.FieldByName(s)
	}
	return v
}

// dump function
func dump(c *Collector, nodes []string, formats []string, prefixPath string) error {
	err := os.MkdirAll(prefixPath, 0700)
	if err != nil {
		return err
	}
	if c == nil {
		c = &Collector{}
	}

	fmt.Printf("dump.InspectMode: %t \n", inArray("inspect", nodes))

	// TODO: create a method for dumping the config struct inspection
	if inArray("inspect", nodes) {
		c.Inspect = inspectStruct(c)
	}

	if nodes[0] == "all" {
		nodes = []string{}
	}

	exportNodesCount := len(nodes)
	for _, f := range formats {
		switch {
		case exportNodesCount == 0:
			nodeName := "config"
			data, err := encodeFile(c, f, nodeName)
			if err != nil {
				fmt.Println("error, msg=", err)
				os.Exit(1)
				return err
			}
			filePath := getConfigDumpFilePath(prefixPath, f, nodeName)
			if err := writeFile(filePath, data); err != nil {
				fmt.Println("error, msg=", err)
				os.Exit(1)
				return err
			}

		case exportNodesCount > 0:
			for _, n := range nodes {

				// encodeFile
				// fieldsToExtract := []string{"FirstField", "ThirdField"}
				value, err := reflections.GetField(c, ucFirst(n))
				if err != nil {
					fmt.Println("error, msg=", err)
					os.Exit(1)
				}

				// pp.Println(value)
				configNode := make(map[string]interface{}, 1)
				configNode[strings.ToLower(n)] = value

				data, err := encodeFile(configNode, f, n)
				if err != nil {
					fmt.Println("error, msg=", err)
					os.Exit(1)
					return err
				}
				filePath := getConfigDumpFilePath(prefixPath, f, n)
				if err := writeFile(filePath, data); err != nil {
					fmt.Println("error, msg=", err)
					os.Exit(1)
					return err
				}
			}
		}
	}
	return nil
}

// encodeFile function
func encodeFile(c interface{}, format, node string) ([]byte, error) {
	switch format {
	case "ini":
		err := ini.MapTo(c, "./colly.ini")
		if err != nil {
			// fmt.Println("error: ", err)
			// os.Exit(1)
			return nil, err
		}

	case "json":
		data, err := json.MarshalIndent(c, " ", "\t")
		if err != nil {
			return nil, err
		}
		return data, nil

	case "toml":
		var dataBytes bytes.Buffer
		if err := toml.NewEncoder(&dataBytes).Encode(c); err != nil {
			return nil, err
		}
		return []byte(dataBytes.String()), nil

	case "xml":
		data, err := xml.MarshalIndent(c, "", "\t")
		if err != nil {
			return nil, err
		}
		return data, nil

	case "yaml", "yml":
		data, err := yaml.Marshal(c)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("Unkown format to export")
}

func getConfigDumpFilePath(prefixPath, format, nodeName string) string {
	return fmt.Sprintf("%s/%s.%s", prefixPath, nodeName, format)
}

func getAttributesListToExport(attrs string) []string {
	return strings.Split(attrs, ",")
}

func writeFile(filePath string, data []byte) error {
	if err := ioutil.WriteFile(filePath, data, 0755); err != nil {
		return err
	}
	return nil
}
