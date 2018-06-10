# FlatStruct [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/ksinica/flatstruct) [![Build Status](https://travis-ci.org/ksinica/flatstruct.svg?branch=master)](https://travis-ci.org/ksinica/flatstruct) [![codecov](https://codecov.io/gh/ksinica/flatstruct/branch/master/graph/badge.svg)](https://codecov.io/gh/ksinica/flatstruct) [![Go Report Card](https://goreportcard.com/badge/github.com/ksinica/flatstruct)](https://goreportcard.com/report/github.com/ksinica/flatstruct)
```flatstruct``` is package for "flattening" structure fields to a map, giving possibility to access structure field values using unique paths.
``` go
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
```
