package PZP26_test

import (
	"time"
)

type People struct {
	ID            string `PZP26:"id"`
	Name          string `PZP26:"name"`
	Age           int    `PZP26:"age"`
	CityID        string `PZP26:"city"`
	UselessField  string `PZP26:"-"`
	UselessField2 string `PZP26:""`
}

type City struct {
	ID   string `PZP26:"id"`
	Name string `PZP26:"name"`
}

var DataSetSmith = People{
	ID:            "1",
	Name:          "Smith",
	Age:           21,
	CityID:        "JKB21",
	UselessField:  "none",
	UselessField2: "null",
}

var DataSetJohn = People{
	ID:            "2",
	Name:          "John",
	Age:           52,
	CityID:        "MKJ86",
	UselessField:  "null",
	UselessField2: "none",
}

var DataSetNewYork = City{
	ID:   "JKB21",
	Name: "New York",
}

var DataSetToronto = City{
	ID:   "MKJ86",
	Name: "Toronto",
}

var STORE = []interface{}{DataSetSmith, DataSetNewYork, DataSetJohn, DataSetToronto}

var Time2000, _ = time.Parse(time.RFC3339, "2000-01-01T01:01:01Z")
var Time2001, _ = time.Parse(time.RFC3339, "2001-01-01T01:01:01Z")

type BoolExpectation struct {
	Result   bool
	Warnings []error
	Errors   []error
}
