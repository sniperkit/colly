package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestParseTerm(t *testing.T) {
	success := 0

	for _, test := range TermTestSet {
		result, warn, err := PZP26.ParseTerm(test.Term)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: ParseTerm(\"%v\") :\n", test.Term)

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
			t.Logf("OK: ParseTerm(\"%v\") = (%v, %v, %v)\n", test.Term, result, warn, err)
			success++
		}
	}

	total := len(TermTestSet)
	fmt.Printf("=== RESULT TestParseTerm : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type TermTest struct {
	Term     string
	Expected TermExpectation
}

type TermExpectation struct {
	Result   PZP26.TermNode
	Warnings []error
	Errors   []error
}

var TermTestSet = []TermTest{
	TermTest{
		Term: "id='1'",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "id", Operator: PZP26.Equals{}, Right: "1"},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "id = '1'",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "id", Operator: PZP26.Equals{}, Right: "1"},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "id = 1",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "id", Operator: PZP26.Equals{}, Right: 1},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "male = true",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "male", Operator: PZP26.Equals{}, Right: true},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "male != true",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "male", Operator: PZP26.NotEquals{}, Right: true},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "male <> true",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "male", Operator: PZP26.NotEquals{}, Right: true},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "age>18",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "age", Operator: PZP26.StrictGreater{}, Right: 18},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "birth>time:RFC3339(2000-01-01T01:01:01Z)",
		Expected: TermExpectation{
			// Time2000 is declared in PZP26_test.go
			Result:   PZP26.TermNode{Left: "birth", Operator: PZP26.StrictGreater{}, Right: Time2000},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "age>'18'",
		Expected: TermExpectation{
			Result: PZP26.TermNode{Left: "age", Operator: PZP26.StrictGreater{}, Right: "18"},
			Warnings: []error{
				PZP26.UnexpectedTermTypeValue{Term: "age>'18'", Got: "string", Expected: "number"},
			},
			Errors: []error{},
		},
	},
	TermTest{
		Term: "age>'18'",
		Expected: TermExpectation{
			Result: PZP26.TermNode{Left: "age", Operator: PZP26.StrictGreater{}, Right: "18"},
			Warnings: []error{
				PZP26.UnexpectedTermTypeValue{Term: "age>'18'", Got: "string", Expected: "number"},
			},
			Errors: []error{},
		},
	},
	TermTest{
		Term: "age IN (18, 25, 30, 40)",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "age", Operator: PZP26.In{}, Right: []int{18, 25, 30, 40}},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "age IN ('18', '25', '30', '40')",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{Left: "age", Operator: PZP26.In{}, Right: []string{"18", "25", "30", "40"}},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "id = 1 OR age > 18",
		Expected: TermExpectation{
			Result: PZP26.TermNode{
				Left:     PZP26.TermNode{Left: "id", Operator: PZP26.Equals{}, Right: 1},
				Operator: PZP26.Or{},
				Right:    PZP26.TermNode{Left: "age", Operator: PZP26.StrictGreater{}, Right: 18},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "id = 1 OR (age > 18)",
		Expected: TermExpectation{
			Result: PZP26.TermNode{
				Left:     PZP26.TermNode{Left: "id", Operator: PZP26.Equals{}, Right: 1},
				Operator: PZP26.Or{},
				Right:    PZP26.TermNode{Left: "age", Operator: PZP26.StrictGreater{}, Right: 18},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "(id = 1) OR age > 18",
		Expected: TermExpectation{
			Result: PZP26.TermNode{
				Left:     PZP26.TermNode{Left: "id", Operator: PZP26.Equals{}, Right: 1},
				Operator: PZP26.Or{},
				Right:    PZP26.TermNode{Left: "age", Operator: PZP26.StrictGreater{}, Right: 18},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "id = 1 OR ",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{},
			Warnings: []error{},
			Errors: []error{
				PZP26.UnclosedTerm{Term: "id = 1 OR "},
			},
		},
	},
	TermTest{
		Term: "id = 1 OR ()",
		Expected: TermExpectation{
			Result:   PZP26.TermNode{},
			Warnings: []error{},
			Errors: []error{
				PZP26.UnspecifiedTerm{Term: "()"},
			},
		},
	},
	TermTest{
		Term: "id = 1 AND (age > 18 OR male = true)",
		Expected: TermExpectation{
			Result: PZP26.TermNode{
				Left:     PZP26.TermNode{Left: "id", Operator: PZP26.Equals{}, Right: 1},
				Operator: PZP26.And{},
				Right: PZP26.TermNode{
					Left:     PZP26.TermNode{Left: "age", Operator: PZP26.StrictGreater{}, Right: 18},
					Operator: PZP26.Or{},
					Right:    PZP26.TermNode{Left: "male", Operator: PZP26.Equals{}, Right: true},
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "(id > 1 AND age < 18) OR male = true",
		Expected: TermExpectation{
			Result: PZP26.TermNode{
				Left: PZP26.TermNode{
					Left:     PZP26.TermNode{Left: "id", Operator: PZP26.StrictGreater{}, Right: 1},
					Operator: PZP26.And{},
					Right:    PZP26.TermNode{Left: "age", Operator: PZP26.StrictLower{}, Right: 18},
				},
				Operator: PZP26.Or{},
				Right:    PZP26.TermNode{Left: "male", Operator: PZP26.Equals{}, Right: true},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	TermTest{
		Term: "id > 1 AND age < 18 OR male = true",
		Expected: TermExpectation{
			Result: PZP26.TermNode{
				Left: PZP26.TermNode{
					Left:     PZP26.TermNode{Left: "id", Operator: PZP26.StrictGreater{}, Right: 1},
					Operator: PZP26.And{},
					Right:    PZP26.TermNode{Left: "age", Operator: PZP26.StrictLower{}, Right: 18},
				},
				Operator: PZP26.Or{},
				Right:    PZP26.TermNode{Left: "male", Operator: PZP26.Equals{}, Right: true},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
}
