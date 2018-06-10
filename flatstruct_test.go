package flatstruct

import (
	"fmt"
	"testing"
)

func ExampleFlatStruct_Flatten() {
	s := struct {
		Foo int `json:"foo"`
		Bar struct {
			Baz string `json:"baz"`
		} `json:"bar"`
	}{}

	s.Foo = 123
	s.Bar.Baz = "one,two,three"

	// By default, "json" key in the struct's tag field will be used to override field name.
	// It can be easily changed by changing FlatStruct NameTag variable.
	fs := NewFlatStruct()
	// Let's define our custom path separator (dot is a default one).
	fs.PathSeparator = "->"
	f, err := fs.Flatten(&s)
	if err != nil {
		panic(err)
	}

	fmt.Println(f["foo"].Value)
	fmt.Println(f["bar->baz"].Value)

	// Output:
	// 123
	// one,two,three
}

func TestBasicUsage(t *testing.T) {
	s := struct {
		Foo int
		Bar struct {
			Baz string `json:"baz"`
		}
		Qux   map[string]interface{}
		Corge string `json:"-"`
	}{}

	s.Foo = 123
	s.Bar.Baz = "hello"
	s.Qux = make(map[string]interface{})
	s.Qux["Quux"] = "world"

	f, err := NewFlatStruct().Flatten(&s)
	if err != nil {
		t.Fatal(err)
	}

	if f["Foo"].Value.Int() != 123 || len(f["Foo"].Tags) != 0 {
		t.Fail()
	}

	if f["Bar.baz"].Value.String() != "hello" || f["Bar.baz"].Tags["json"][0] != "baz" {
		t.Fail()
	}

	if f["Qux.Quux"].Value.String() != "world" {
		t.Fail()
	}

	if _, ok := f["Corge"]; ok {
		t.Fail()
	}

	m := make(map[string]interface{})
	m["Foo"] = 123
	m["Bar"] = make(map[string]interface{})
	m["Bar"].(map[string]interface{})["baz"] = "baz"

	g, err := NewFlatStruct().Flatten(m)
	if err != nil {
		t.Fatal(err)
	}

	if g["Foo"].Value.Int() != 123 {
		t.Fail()
	}

	if g["Bar.baz"].Value.String() != "baz" {
		t.Fail()
	}
}

func TestBasicUsageErrors(t *testing.T) {
	if _, err := NewFlatStruct().Flatten("error!"); err == nil {
		t.Fail()
	}

	if _, err := NewFlatStruct().Flatten(struct {
		Foo string `json:"foo...bar"`
	}{}); err == nil {
		t.Fail()
	}

	if _, err := NewFlatStruct().Flatten(struct {
		Foo int    `json:"foo"`
		Bar string `json:"foo"`
	}{}); err == nil {
		t.Fail()
	}
}
