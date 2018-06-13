package colly

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	// "github.com/sniperkit/structs"

	// iterators
	json "encoding/json"
	xml "encoding/xml"
	toml "github.com/BurntSushi/toml"
	ini "github.com/go-ini/ini"
	yaml "github.com/sniperkit/yaml"
)

func inArray(e string, s []string) bool {
	for _, a := range s {
		fmt.Println("inArray.match=", e, "val", a)
		if a == e {
			return true
		}
	}
	return false
}

func field(t interface{}, key string) reflect.Value {
	strs := strings.Split(key, ".")
	v := reflect.ValueOf(t)
	for _, s := range strs[1:] {
		v = v.FieldByName(s)
	}
	return v
}

// nodes := []string{"contacts", "db", "oauth2"}
// configor.Dump(Config, "yaml", "contacts", "db", "oauth2")
func dump(c *Collector, nodes []string, formats []string, prefixPath string) error {
	err := os.MkdirAll(prefixPath, 0700)
	if err != nil {
		return err
	}
	if c == nil {
		c = &Collector{}
	}

	// fmt.Println("nodes=", nodes)
	// fmt.Println("formats=", formats)
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

			// s := reflect.ValueOf(c).Elem()
			// v := reflect.ValueOf(c)

			// s := structs.New(c)

			for _, n := range nodes {

				// fn := s.FieldByName(n).Interface()
				// fmt.Println(fn)
				// fn := v.FieldByName(n)

				// fn := s.Field(n)

				// Get the value for addr
				// a := addrField.Value().(string)

				// Or get all fields
				// fn := s.Field(c).Fields()

				data, err := encodeFile(c, f, n)
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

func encodeFile(c interface{}, format, node string) ([]byte, error) {

	// fmt.Println("node=", node, ", format=", format)
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
