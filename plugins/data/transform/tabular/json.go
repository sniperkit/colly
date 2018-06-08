package tablib

import (
	"encoding/json"
	"fmt"
	"strings"

	gjson "github.com/sniperkit/colly/plugins/data/extract/json/gson"
	mxj "github.com/sniperkit/colly/plugins/data/transform/mxj/v2/pkg"
	// "github.com/sniperkit/colly/plugins/data/transform/mxj/v1"

	pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

// LoadMXJ loads a dataset from a XML/JSON source.
// Forked from `github.com/clbanning/mxj`
func LoadMXJ(jsonContent []byte) (*Dataset, error) {

	// var input []map[string]interface{}
	mxj.JsonUseNumber = true
	mv, err := mxj.NewMapJson(jsonContent)
	if err != nil {
		fmt.Println("NewMapJson, error: ", err)
		return nil, err
	}

	var paths []string
	// pp.Println("jsonContent=", string(jsonContent))
	// pp.Println("mv=", mv)
	mxj.LeafUseDotNotation()
	paths = mv.LeafPaths()
	pp.Println(paths)

	return nil, ErrUnmarshallingJsonWithMxj
	// return internalLoadFromDict(input)
}

// LoadGSON loads a dataset from a JSON source.
// Forked from `github.com/tidwall/gjson`
func LoadGJSON(jsonContent []byte) (*Dataset, error) {

	// var input []map[string]interface{}
	// results := gjson.GetMany(json, "name.first", "name.last", "age")
	// input = gjson.GetBytes(jsonContent, "").Value().([]map[string]interface{})
	m, ok := gjson.GetBytes(jsonContent, "").Value().(map[string]interface{})
	if !ok {
		fmt.Println("Error")
		return nil, ErrUnmarshallingJsonWithGson
	}
	pp.Println("map[string]interface{}=", m)

	return nil, ErrUnmarshallingJsonWithGson
	// return internalLoadFromDict(input)
}

// LoadJSON loads a dataset from a JSON source.
func LoadJSON(jsonContent []byte) (*Dataset, error) {

	var input []map[string]interface{}
	d := json.NewDecoder(strings.NewReader(string(jsonContent)))
	d.UseNumber()
	if err := d.Decode(&input); err != nil {
		return nil, err
	}

	// if err := json.Unmarshal(jsonContent, &input); err != nil {
	// 	return nil, err
	// }

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
