package tablib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	// plugins - json parsers
	/*
		gjson "github.com/sniperkit/colly/plugins/data/parser/json/gson"
		djson "github.com/sniperkit/colly/plugins/data/parser/json/djson"
		easyjson "github.com/sniperkit/colly/plugins/data/parser/json/easyjson"
		fastjson "github.com/sniperkit/colly/plugins/data/parser/json/fastjson"
		ffjson "github.com/sniperkit/colly/plugins/data/parser/json/ffjson"
		gabs "github.com/sniperkit/colly/plugins/data/parser/json/gabs"
		jsnm "github.com/sniperkit/colly/plugins/data/parser/json/jsnm"
		jsonez "github.com/sniperkit/colly/plugins/data/parser/json/jsonez"
		jsonparser "github.com/sniperkit/colly/plugins/data/parser/json/jsonparser"
		jsonstream "github.com/sniperkit/colly/plugins/data/parser/json/jsonstream"
		lazyjson "github.com/sniperkit/colly/plugins/data/parser/json/lazyjson"
	*/
	// gjson "github.com/sniperkit/colly/plugins/data/parser/json/gjson"    // fork
	mxj "github.com/sniperkit/colly/plugins/data/parser/json/mxj/v2/pkg" // fork v2
	// mxj "github.com/sniperkit/colly/plugins/data/transform/mxj/master" // latest commit
	// mxj "github.com/sniperkit/colly/plugins/data/transform/mxj/v1" // fork v1
	// dev helpers
	// pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

var headerStyleTable = map[string]KeyStyle{
	"jsonpointer": JSONPointerStyle,
	"slash":       SlashStyle,
	"underscore":  UnderscoreStyle,
	"dot":         DotNotationStyle,
	"dot-bracket": DotBracketStyle,
}

// LoadJSON loads a dataset from a JSON source.
// - Default pkg "encoding/json"
func LoadJSON(jsonContent []byte) (*Dataset, error) {

	var input []map[string]interface{}

	// quick hack to create json array for tablib ^^
	if jsonContent[0] == '{' {
		jsonContent = []byte(`[` + string(jsonContent) + `]`)
	}

	d := json.NewDecoder(strings.NewReader(string(jsonContent)))
	d.UseNumber()
	if err := d.Decode(&input); err != nil {
		return nil, err
	}

	// if _, ok := headerStyleTable[c.String("header-style")]; !ok {
	//	return fmt.Errorf("Invalid --header-style value %q", c.String("header-style"))
	// }

	var err error
	input, err = ToMapSlices(input)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return internalLoadFromDict(input)
}

func LoadGJSON(jsonContent []byte) (*Dataset, error) {
	return nil, nil
}

// LoadMXJ loads a dataset from a XML/JSON source.
// - MXJ allows to decode / encode XML or JSON to/from map[string]interface{};
//   extract values with dot-notation paths and wildcards.
// - Forked from `github.com/clbanning/mxj`
func LoadMXJ(jsonContent []byte) (*Dataset, error) {

	mxj.JsonUseNumber = true
	mv, err := mxj.NewMapJsonArray(jsonContent)
	if err != nil {
		return nil, err
	}

	// mxj.LeafUseDotNotation()
	// fmt.Println("TypeOf=", reflect.TypeOf(mv))

	var input []map[string]interface{}
	var errConvert error
	input, errConvert = ToMapSlices(mv)
	if errConvert != nil {
		fmt.Println("error:", errConvert)
		return nil, errConvert
	}

	return internalLoadFromDict(input)
}

// LoadDatabookJSON loads a Databook from a JSON source.
func LoadDatabookJSON(jsonContent []byte) (*Databook, error) {
	var input []map[string]interface{}
	var internalInput []map[string]interface{}
	if err := json.Unmarshal(jsonContent, &input); err != nil {
		return nil, err
	}

	db := NewDatabook()
	for _, d := range input {
		b, err := json.MarshalIndent(d["data"], "", "\t")
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, &internalInput); err != nil {
			return nil, err
		}

		if ds, err := internalLoadFromDict(internalInput); err == nil {
			db.AddSheet(d["title"].(string), ds)
		} else {
			return nil, err
		}
	}

	return db, nil
}

// JSON returns a JSON representation of the Dataset as an Export.
func (d *Dataset) JSON() (*Export, error) {
	back := d.Dict()

	b, err := json.MarshalIndent(back, "", "\t")
	if err != nil {
		return nil, err
	}
	return newExportFromBytes(b), nil
}

// JSON returns a JSON representation of the Databook as an Export.
func (d *Databook) JSON() (*Export, error) {
	b := newBuffer()
	b.WriteString("[")
	for _, s := range d.sheets {
		b.WriteString("{\"title\": \"" + s.title + "\", \"data\": ")
		js, err := s.dataset.JSON()
		if err != nil {
			return nil, err
		}
		b.Write(js.Bytes())
		b.WriteString("},")
	}
	by := b.Bytes()
	by[len(by)-1] = ']'
	return newExportFromBytes(by), nil
}

func JsonMarshal(t interface{}) ([]byte, error) {
	buff := &bytes.Buffer{}
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(t)
	return buff.Bytes(), err
}

func JsonMarshalIndent(t interface{}, prefix, indent string) ([]byte, error) {
	b, err := JsonMarshal(t)
	if err != nil {
		return b, err
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
