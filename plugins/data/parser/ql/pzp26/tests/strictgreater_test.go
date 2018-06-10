package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestStrictGreater(t *testing.T) {
	success := 0

	strictgreater := PZP26.StrictGreater{}
	for _, test := range StrictGreaterTestSet {
		result, warn, err := strictgreater.Compute(test.Left, test.Right)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: strictgreater.Compute(%v, %v) :\n", test.Left, test.Right)

			if !reflect.DeepEqual(result, test.Expected.Result) {
				t.Errorf("\tGot %#v, expected %#v\n", result, test.Expected.Result)
			} else {
				t.Errorf("\tGot %v, it is OK\n", result)
			}

			if !reflect.DeepEqual(warn, test.Expected.Warnings) {
				t.Errorf("\tGot %#v, expected %#v\n", warn, test.Expected.Warnings)
			} else {
				t.Errorf("\tGot %v, it is OK\n", warn)
			}

			if !reflect.DeepEqual(err, test.Expected.Errors) {
				t.Errorf("\tGot %#v, expected %#v\n", err, test.Expected.Errors)
			} else {
				t.Errorf("\tGot %v, it is OK\n", err)
			}
		} else {
			t.Logf("OK: strictgreater.Compute(%v, %v) = (%v, %v, %v)\n", test.Left, test.Right, result, warn, err)
			success++
		}
	}

	total := len(StrictGreaterTestSet)
	fmt.Printf("=== RESULT TestStrictGreater : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type StrictGreaterTest struct {
	Right    interface{}
	Left     interface{}
	Expected BoolExpectation
}

var StrictGreaterTestSet = []StrictGreaterTest{
	StrictGreaterTest{
		Right:    1,
		Left:     1,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    6,
		Left:     1,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    1,
		Left:     6,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    "AA",
		Left:     "ZZ",
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    "ZZ",
		Left:     "AA",
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    'A',
		Left:     'Z',
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    'Z',
		Left:     'A',
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    'A',
		Left:     'A',
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    '1',
		Left:     'A',
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    'A',
		Left:     '1',
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    1.6,
		Left:     6.1,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	// Time2000 and Time2001 are declared in PZP26_test.go
	StrictGreaterTest{
		Right:    Time2000,
		Left:     Time2001,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right:    Time2001,
		Left:     Time2000,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	StrictGreaterTest{
		Right: true,
		Left:  true,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UncomparableValues{Test: "true > true", TypeRight: "boolean", TypeLeft: "boolean"},
			},
		},
	},
	StrictGreaterTest{
		Right: true,
		Left:  false,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UncomparableValues{Test: "true > false", TypeRight: "boolean", TypeLeft: "boolean"},
			},
		},
	},
	StrictGreaterTest{
		Right: true,
		Left:  1.6,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UncomparableValues{Test: "true > 1.6", TypeRight: "boolean", TypeLeft: "number"},
			},
		},
	},
	StrictGreaterTest{
		Right: 'c',
		Left:  1.6,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UncomparableValues{Test: "'c' > 1.6", TypeRight: "char", TypeLeft: "number"},
			},
		},
	},
	StrictGreaterTest{
		Right: 'c',
		Left:  nil,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UncomparableValues{Test: "'c' > null", TypeRight: "char", TypeLeft: "null"},
			},
		},
	},
	StrictGreaterTest{
		Right: nil,
		Left:  'c',
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UncomparableValues{Test: "null > 'c'", TypeRight: "null", TypeLeft: "char"},
			},
		},
	},
	StrictGreaterTest{
		Right: nil,
		Left:  nil,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UncomparableValues{Test: "null > null", TypeRight: "null", TypeLeft: "null"},
			},
		},
	},
}
