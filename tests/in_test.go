package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestIn(t *testing.T) {
	success := 0

	in := PZP26.In{}
	for _, test := range InTestSet {
		result, warn, err := in.Compute(test.Left, test.Right)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: in.Compute(%v, %v) :\n", test.Left, test.Right)

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
			t.Logf("OK: in.Compute(%v, %v) = (%v, %v, %v)\n", test.Left, test.Right, result, warn, err)
			success++
		}
	}

	total := len(InTestSet)
	fmt.Printf("=== RESULT TestIn : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type InTest struct {
	Right    interface{}
	Left     interface{}
	Expected BoolExpectation
}

var InTestSet = []InTest{
	InTest{
		Right:    12,
		Left:     []int{1, 2, 3, 7, 12, 67, 321},
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	InTest{
		Right:    "Test#54",
		Left:     []string{"alpha", "beta", "72", "Test#33", "Test#54"},
		Expected: BoolExpectation{Result: true, Warnings: []error{}, Errors: []error{}},
	},
	InTest{
		Right:    nil,
		Left:     []string{"alpha", "beta", "72", "Test#33", "Test#54"},
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{}},
	},
	InTest{
		Right: 31,
		Left:  "test",
		Expected: BoolExpectation{Result: false, Warnings: []error{
			PZP26.NotAnArray{Value: "test"},
		}, Errors: []error{}},
	},
	InTest{
		Right: 31,
		Left:  []string{"22", "31"},
		Expected: BoolExpectation{Result: false, Warnings: []error{}, Errors: []error{
			PZP26.UncomparableValues{Test: "In", TypeLeft: "int", TypeRight: "string"},
		}},
	},
	InTest{
		Right: 31,
		Left:  31,
		Expected: BoolExpectation{Result: true, Warnings: []error{
			PZP26.NotAnArray{Value: 31},
		}, Errors: []error{}},
	},
}
