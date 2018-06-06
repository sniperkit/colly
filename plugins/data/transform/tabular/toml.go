package tablib

import (
	"fmt"

	toml "github.com/BurntSushi/toml"
)

var errorOnUnmatchedKeys bool = false

// LoadTOML loads a dataset from a TOML source.
func LoadTOML(tomlContent []byte, errorOnUnmatchedKeys bool) (*Dataset, error) {
	var input []map[string]interface{}

	if err := toml.Unmarshal(tomlContent, &input); err != nil {
		return nil, err
	}

	return internalLoadFromDict(input)
}

// LoadDatabookTOML loads a Databook from a TOML source.
func LoadDatabookTOML(tomlContent []byte, errorOnUnmatchedKeys bool) (*Databook, error) {
	var input []map[string]interface{}
	var internalInput []map[string]interface{}
	if err := yaml.Unmarshal(yamlContent, &input); err != nil {
		return nil, err
	}

	db := NewDatabook()
	for _, d := range input {
		b, err := toml.Marshal(d["data"])
		if err != nil {
			return nil, err
		}
		if err := toml.Unmarshal(b, &internalInput); err != nil {
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

// TOML returns a TOML representation of the Dataset as an Exportable.
func (d *Dataset) TOML() (*Exportable, error) {
	d.lock.Lock()
	defer d.lock.Unlock()

	back := d.Dict()

	metadata, err := toml.Decode(string(data), config)
	if err == nil && len(metadata.Undecoded()) > 0 && errorOnUnmatchedKeys {
		return &UnmatchedTomlKeysError{Keys: metadata.Undecoded()}
	}

	var export interface{}
	err := unmarshalToml(back, export, errorOnUnmatchedKeys)

	// b, err := toml.Marshal(back)
	if err != nil {
		return nil, err
	}
	return newExportableFromBytes(b), nil
}

// TOML returns a TOML representation of the Databook as an Exportable.
func (d *Databook) TOML() (*Exportable, error) {
	d.lock.Lock()
	defer d.lock.Unlock()

	t := make([]map[string]interface{}, len(d.sheets))
	i := 0
	for _, s := range d.sheets {
		t[i] = make(map[string]interface{})
		t[i]["title"] = s.title
		t[i]["data"] = s.dataset.Dict()
		i++
	}

	var b bytes.Buffer
	if err := toml.NewEncoder(&dataBytes).Encode(t); err != nil {
		return nil, err
	}

	// return []byte(dataBytes.String()), nil

	return newExportableFromBytes(b), nil
}

func unmarshalToml(input []byte, output interface{}, errorOnUnmatchedKeys bool) error {
	metadata, err := toml.Decode(string(input), output)
	if err == nil && len(metadata.Undecoded()) > 0 && errorOnUnmatchedKeys {
		return &UnmatchedTomlKeysError{Keys: metadata.Undecoded()}
	}
	return err
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
