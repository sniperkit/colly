package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

// TODO test fields name priorities : PZP26, json, native

func TestResolver(t *testing.T) {
	success := 0

	resolver := PZP26.Resolver{
		ResolveFunc: func(terms PZP26.TermNode) (result interface{}, warnings []error, errors []error, continueRecursion bool) {
			return nil, []error{}, []error{}, false
		},
	}

	for _, test := range ResolveTestSet {
		result, warn, err := resolver.ResolveEach(test.Query)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: resolver.ResolveEach(%v) :\n", test.Query)

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
			t.Logf("OK: resolver.ResolveEach(%v) = (%v, %v, %v)\n", test.Query, result, warn, err)
			success++
		}
	}

	total := len(ResolveTestSet)
	fmt.Printf("=== RESULT TestResolver : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type ResolveTest struct {
	Query    []PZP26.QueryNode
	Expected ResolveExpectation
}

type ResolveExpectation struct {
	Result   interface{}
	Warnings []error
	Errors   []error
}

var ResolveTestSet = []ResolveTest{
	ResolveTest{
		Query: []PZP26.QueryNode{
			PZP26.QueryNode{
				Name: "root[0]",
				Terms: PZP26.TermNode{
					Left:     "id",
					Operator: PZP26.Equals{},
					Right:    "1",
				},
				AllFields:  true,
				GoesToRoot: true,
			},
		},
		Expected: ResolveExpectation{
			Result:   DataSetSmith,
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	ResolveTest{
		Query: []PZP26.QueryNode{
			PZP26.QueryNode{
				Name: "root[0]",
				Terms: PZP26.TermNode{
					Left:     "id",
					Operator: PZP26.Equals{},
					Right:    "1",
				},
				Fields: []PZP26.QueryNode{
					PZP26.QueryNode{
						Name:       "city",
						AllFields:  true,
						GoesToRoot: true,
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Expected: ResolveExpectation{
			Result:   []interface{}{DataSetSmith, DataSetNewYork},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	ResolveTest{
		Query: []PZP26.QueryNode{
			PZP26.QueryNode{
				Name: "root[0]",
				Terms: PZP26.TermNode{
					Left:     "id",
					Operator: PZP26.Equals{},
					Right:    "1",
				},
				Fields: []PZP26.QueryNode{
					PZP26.QueryNode{
						Name: "city",
						Terms: PZP26.TermNode{
							Left:     "id",
							Operator: PZP26.Equals{},
							Right:    "1",
						},
						AllFields:  true,
						GoesToRoot: true,
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Expected: ResolveExpectation{
			Result:   []interface{}{DataSetSmith},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	ResolveTest{
		Query: []PZP26.QueryNode{
			PZP26.QueryNode{
				Name: "root[0]",
				Terms: PZP26.TermNode{
					Left:     "id",
					Operator: PZP26.Equals{},
					Right:    "1",
				},
				Fields: []PZP26.QueryNode{
					PZP26.QueryNode{
						Name:       "city",
						AllFields:  true,
						GoesToRoot: false,
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Expected: ResolveExpectation{
			Result:   []interface{}{DataSetSmith},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	ResolveTest{
		Query: []PZP26.QueryNode{
			PZP26.QueryNode{
				Name: "root[0]",
				Terms: PZP26.TermNode{
					Left:     "id",
					Operator: PZP26.Equals{},
					Right:    "nonexistent",
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Expected: ResolveExpectation{
			Result:   []interface{}{},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
}
