package configor

import (
	"github.com/sniperkit/structs"
)

type Debug struct {

	// Inspect
	Inspect struct {

		// Struct
		Struct struct {

			// Return the Config struct name
			Name string `json:"name" yaml:"name" toml:"name" xml:"name" ini:"name"`

			// Check if any field of the `Config` struct is initialized or not.
			HasZero bool `json:"has_zero" yaml:"has_zero" toml:"has_zero" xml:"hasZero" ini:"hasZero"`

			// Check if all fields of the `Config` struct is initialized or not.
			IsZero bool `json:"is_zero" yaml:"is_zero" toml:"is_zero" xml:"isZero" ini:"isZero"`

			// Check if `Config` is a struct or a pointer to struct
			IsStruct bool `json:"is_struct" yaml:"is_struct" toml:"is_struct" xml:"isStruct" ini:"isStruct"`

			Details struct {

				// Convert `Config` struct to a map[string]interface{}
				Map map[string]interface{} `json:"map" yaml:"map" toml:"map" xml:"map" ini:"map"`

				// Convert `Config` struct to a []interface{}
				Values []interface{} `json:"values" yaml:"values" toml:"values" xml:"values" ini:"values"`

				// Convert `Config` struct to a []string
				Names []string `json:"names" yaml:"names" toml:"names" xml:"names" ini:"names"`

				// Convert `Config` struct to a []*Field
				Fields []*structs.Field `json:"fields" yaml:"fields" toml:"fields" xml:"fields" ini:"fields"`

				//-- End
			} `json:"details" yaml:"details" toml:"details" xml:"details" ini:"details"`

			//-- End
		} `json:"struct" yaml:"struct" toml:"struct" xml:"struct" ini:"struct"`

		//-- End
	} `json:"inspect" yaml:"inspect" toml:"inspect" xml:"inspect" ini:"inspect"`

	//-- End
}

func inspectStruct(obj interface{}) *Debug {
	debug := &Debug{}

	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	debug.Inspect.Struct.Details.Map = structs.Map(obj)

	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	debug.Inspect.Struct.Details.Values = structs.Values(obj)

	// Convert the names of a struct to a []string
	// (see "Names methods" for more info about fields)
	debug.Inspect.Struct.Details.Names = structs.Names(obj)

	// Convert the values of a struct to a []*Field
	// (see "Field methods" for more info about fields)
	debug.Inspect.Struct.Details.Fields = structs.Fields(obj)

	// Return the struct name => "Server"
	debug.Inspect.Struct.Name = structs.Name(obj)

	// Check if any field of a struct is initialized or not.
	debug.Inspect.Struct.HasZero = structs.HasZero(obj)

	// Check if all fields of a struct is initialized or not.
	debug.Inspect.Struct.IsZero = structs.IsZero(obj)

	// Check if server is a struct or a pointer to struct
	debug.Inspect.Struct.IsStruct = structs.IsStruct(obj)

	return debug
}
