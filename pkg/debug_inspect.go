package colly

import (
	// "github.com/sniperkit/metaflector"
	"github.com/sniperkit/structs"
)

/*
	Dev-Only:
	- Used to debug the collector and config struct
	  - Check if to IsStruct(), IsZero(), HasZero()
	  - Convert to Map(), Values(), Names(), Fields()
*/

// InspectMode specifies if the inspector is enabled by default
var InspectMode bool = false

// Debug represents...
type Debug struct {

	// Inspect
	Inspect struct {

		// Struct
		Config struct {

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
				Fields []*structs.Field `json:"-" yaml:"-" toml:"-" xml:"-" ini:"-"`

				//-- End
			} `json:"details" yaml:"details" toml:"details" xml:"details" ini:"details"`

			//-- End
		} `json:"struct" yaml:"struct" toml:"struct" xml:"struct" ini:"struct"`

		//-- End
	} `json:"inspect" yaml:"inspect" toml:"inspect" xml:"inspect" ini:"inspect"`

	//-- End
}

// inspectStruct function extract from the collector object some infos
func inspectStruct(c *Collector) *Debug {
	debug := &Debug{}
	return debug

	structs.OnlyExported = true

	/*
		metaflector.EachField(c, func(obj interface{}, name string, kind reflect.Kind) {
			fmt.Printf("obj=%v name=%v kind=%v\n", obj, name, kind)
		})
	*/

	// sd := structs.New(c)
	// prettyPrinter("inspect", sd)

	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	debug.Inspect.Config.Details.Map = structs.Map(c)

	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	debug.Inspect.Config.Details.Values = structs.Values(c)

	// Convert the names of a struct to a []string
	// (see "Names methods" for more info about fields)
	debug.Inspect.Config.Details.Names = structs.Names(c)

	// Convert the values of a struct to a []*Field
	// (see "Field methods" for more info about fields)
	debug.Inspect.Config.Details.Fields = structs.Fields(c)

	// Return the struct name => "Server"
	debug.Inspect.Config.Name = structs.Name(c)

	// Check if any field of a struct is initialized or not.
	debug.Inspect.Config.HasZero = structs.HasZero(c)

	// Check if all fields of a struct is initialized or not.
	debug.Inspect.Config.IsZero = structs.IsZero(c)

	// Check if server is a struct or a pointer to struct
	debug.Inspect.Config.IsStruct = structs.IsStruct(c)

	// prettyPrinter("debug", debug)

	return debug
}
