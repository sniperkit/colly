package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestEquals(t *testing.T) {
	success := 0

	equals := PZP26.Equals{}
	for _, test := range EqualsTestSet {
		result, warn, err := equals.Compute(test.Left, test.Right)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: equals.Compute(%v, %v) :\n", test.Left, test.Right)

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
			t.Logf("OK: equals.Compute(%v, %v) = (%v, %v, %v)\n", test.Left, test.Right, result, warn, err)
			success++
		}
	}

	total := len(EqualsTestSet)
	fmt.Printf("=== RESULT TestEquals : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type EqualsTest struct {
	Right    interface{}
	Left     interface{}
	Expected BoolExpectation
}

var EqualsTestSet = []EqualsTest{
	EqualsTest{
		Right:    1,
		Left:     1,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    "474",
		Left:     "163",
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    "163",
		Left:     "474",
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    "4u74",
		Left:     "16u3",
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    "4u74",
		Left:     "4u74",
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    true,
		Left:     true,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    true,
		Left:     false,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    1.6,
		Left:     6.1,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    1.6,
		Left:     1.6,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    'c',
		Left:     'g',
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    'c',
		Left:     'c',
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    'c',
		Left:     nil,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    nil,
		Left:     'c',
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    nil,
		Left:     nil,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    PZP26.Or{},
		Left:     PZP26.Or{},
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	EqualsTest{
		Right:    PZP26.Or{},
		Left:     nil,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
}
