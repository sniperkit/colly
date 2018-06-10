package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestOr(t *testing.T) {
	success := 0

	or := PZP26.Or{}
	for _, test := range OrTestSet {
		result, warn, err := or.Compute(test.Left, test.Right)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: or.Compute(%v, %v) :\n", test.Left, test.Right)

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
			t.Logf("OK: or.Compute(%v, %v) = (%v, %v, %v)\n", test.Left, test.Right, result, warn, err)
			success++
		}
	}

	total := len(OrTestSet)
	fmt.Printf("=== RESULT TestOr : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type OrTest struct {
	Right    interface{}
	Left     interface{}
	Expected BoolExpectation
}

var OrTestSet = []OrTest{
	OrTest{
		Right:    false,
		Left:     true,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right:    false,
		Left:     false,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right:    true,
		Left:     false,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right:    true,
		Left:     true,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right:    nil,
		Left:     true,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right:    true,
		Left:     nil,
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right:    false,
		Left:     nil,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right:    nil,
		Left:     nil,
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	OrTest{
		Right: PZP26.Or{},
		Left:  true,
		Expected: BoolExpectation{
			Result:   false,
			Warnings: []error{},
			Errors: []error{
				PZP26.UnexpectedTermTypeValue{Term: "PZP26.Or{} OR true", Got: "PZP26.Or", Expected: "boolean"},
			},
		},
	},
}
