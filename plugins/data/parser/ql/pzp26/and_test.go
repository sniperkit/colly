package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestAnd(t *testing.T) {
	success := 0

	and := PZP26.And{}
	for _, test := range AndTestSet {
		result, warn, err := and.Compute(test.Left, test.Right)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			text := ""

			if !reflect.DeepEqual(result, test.Expected.Result) {
				text += fmt.Sprintf("\tGot %#v, expected %#v\n", result, test.Expected.Result)
			} else {
				text += fmt.Sprintf("\tGot %v, it is OK\n", result)
			}

			if !reflect.DeepEqual(warn, test.Expected.Warnings) {
				text += fmt.Sprintf("\tGot %#v, expected %#v\n", warn, test.Expected.Warnings)
			} else {
				text += fmt.Sprintf("\tGot %v, it is OK\n", warn)
			}

			if !reflect.DeepEqual(err, test.Expected.Errors) {
				text += fmt.Sprintf("\tGot %#v, expected %#v\n", err, test.Expected.Errors)
			} else {
				text += fmt.Sprintf("\tGot %v, it is OK\n", err)
			}

			t.Errorf("ERROR: and.Compute(%v, %v) :\n%v", test.Left, test.Right, text)
		} else {
			t.Logf("OK: and.Compute(%v, %v) = (%v, %v, %v)\n", test.Left, test.Right, result, warn, err)
			success++
		}
	}

	total := len(AndTestSet)
	fmt.Printf("=== RESULT TestAnd : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type AndTest struct {
	Right    interface{}
	Left     interface{}
	Expected BoolExpectation
}

var AndTestSet = []AndTest{
	AndTest{
		Right:    false,
		Left:     true,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right:    false,
		Left:     false,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right:    true,
		Left:     false,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right:    true,
		Left:     true,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right:    nil,
		Left:     true,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right:    true,
		Left:     nil,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right:    false,
		Left:     nil,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right:    nil,
		Left:     nil,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	AndTest{
		Right: PZP26.Or{},
		Left:  true,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UnexpectedTermTypeValue{Term: "PZP26.0r{} OR true", Got: "PZP26.0r", Expected: "boolean"},
			},
		},
	},
}
