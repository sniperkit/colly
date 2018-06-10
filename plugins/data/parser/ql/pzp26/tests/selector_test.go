package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestSelector(t *testing.T) {
	success := 0

	for _, test := range SelectorTestSet {
		result, warn, err := PZP26.Selector(test.Query, test.Resolved)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: Selector(%v, %v) :\n", test.Query, test.Resolved)

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
			t.Logf("OK: Selector(%v, %v) = (%v, %v, %v)\n", test.Query, test.Resolved, result, warn, err)
			success++
		}
	}

	total := len(SelectorTestSet)
	fmt.Printf("=== RESULT TestSelector : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type SelectorTest struct {
	Query    []PZP26.QueryNode
	Resolved []interface{}
	Expected SelectorExpectation
}

type SelectorExpectation struct {
	Result   map[string]interface{}
	Warnings []error
	Errors   []error
}

var SelectorTestSet = []SelectorTest{
	SelectorTest{
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
						Name: "name",
					},
				},
				GoesToRoot: true,
			},
		},
		Resolved: []interface{}{DataSetSmith},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"root[0]": map[string]interface{}{
					"name": "Smith",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	// TODO : test selectors on nested-fields, like city.population>150000
	SelectorTest{
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
		Resolved: []interface{}{DataSetSmith},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"root[0]": map[string]interface{}{
					"__type": []string{"People"},
					"id":     "1",
					"name":   "Smith",
					"city":   "JKB21",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	SelectorTest{
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
						Fields: []PZP26.QueryNode{
							PZP26.QueryNode{
								Name: "name",
							},
						},
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Resolved: []interface{}{DataSetSmith},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"root[0]": map[string]interface{}{
					"__type": []string{"People"},
					"id":     "1",
					"name":   "Smith",
					"city": map[string]interface{}{
						"name": "New York",
					},
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	SelectorTest{
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
				AllFields:  true,
				GoesToRoot: true,
			},
		},
		Resolved: []interface{}{DataSetSmith, DataSetNewYork},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"root[0]": map[string]interface{}{
					"__type": []string{"People"},
					"id":     "1",
					"name":   "Smith",
					"city":   "JKB21",
				},
				"JKB21": map[string]interface{}{
					"__type": []string{"City"},
					"id":     "JKB21",
					"name":   "New York",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	SelectorTest{
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
						Name: "id",
					},
					PZP26.QueryNode{
						Name: "nonexisting",
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Resolved: []interface{}{DataSetSmith},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"root[0]": map[string]interface{}{
					"__type": []string{"People"},
					"id":     "1",
				},
			},
			Warnings: []error{
				PZP26.FieldNonExistingOnData{FieldName: "nonexisting", FieldPath: "root[0].nonexisting", Data: DataSetSmith},
			},
			Errors: []error{},
		},
	},
	SelectorTest{
		Query: []PZP26.QueryNode{
			PZP26.QueryNode{
				Name:   "root[0]",
				Rename: "alpha",
				Terms: PZP26.TermNode{
					Left:     "id",
					Operator: PZP26.Equals{},
					Right:    "1",
				},
				Fields: []PZP26.QueryNode{
					PZP26.QueryNode{
						Name: "id",
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Resolved: []interface{}{DataSetSmith},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"alpha": map[string]interface{}{
					"__type": []string{"People"},
					"id":     "1",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	SelectorTest{
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
						Index:      "`name`",
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Resolved: []interface{}{DataSetSmith, DataSetNewYork},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"root[0]": map[string]interface{}{
					"city": "@New York",
				},
				"New York": map[string]interface{}{
					"__type": []string{"City"},
					"id":     "JKB21",
					"name":   "New York",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	SelectorTest{
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
						Index:      "Cities#`id`",
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
		},
		Resolved: []interface{}{DataSetSmith, DataSetNewYork},
		Expected: SelectorExpectation{
			Result: map[string]interface{}{
				"root[0]": map[string]interface{}{
					"city": "@Cities#JKB21",
				},
				"Cities#JKB21": map[string]interface{}{
					"__type": []string{"City"},
					"id":     "JKB21",
					"name":   "New York",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
}
