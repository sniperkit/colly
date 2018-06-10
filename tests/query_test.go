package PZP26_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jimskapt/PZP26"
)

func TestParseQuery(t *testing.T) {
	success := 0

	for _, test := range QueryTestSet {
		result, warn, err := PZP26.ParseQuery(test.Query)

		if !reflect.DeepEqual(result, test.Expected.Result) || !reflect.DeepEqual(warn, test.Expected.Warnings) || !reflect.DeepEqual(err, test.Expected.Errors) {
			t.Error()
			t.Errorf("ERROR: ParseQuery(\"%v\") :\n", test.Query)

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
			t.Logf("OK: ParseQuery(\"%v\") = (%v, %v, %v)\n", test.Query, result, warn, err)
			success++
		}
	}

	total := len(QueryTestSet)
	fmt.Printf("=== RESULT TestParseQuery : %v%% (%v/%v)\n", int((float32(success)/float32(total))*100), success, total)
}

type QueryTest struct {
	Query    string
	Expected QueryExpectation
}

type QueryExpectation struct {
	Result   []PZP26.QueryNode
	Warnings []error
	Errors   []error
}

var idTerm1 = PZP26.TermNode{
	Left:     "id",
	Operator: PZP26.Equals{},
	Right:    "1",
}
var idTerm2 = PZP26.TermNode{
	Left:     "id",
	Operator: PZP26.Equals{},
	Right:    "2",
}
var populationTerm = PZP26.TermNode{
	Left:     "population",
	Operator: PZP26.StrictGreater{},
	Right:    150000,
}
var ageTerm = PZP26.TermNode{
	Left:     "age",
	Operator: PZP26.StrictGreater{},
	Right:    18,
}

var QueryTestSet = []QueryTest{
	QueryTest{
		Query: "(id='1'){__all_fields}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:       "root[0]",
					Rename:     "",
					Terms:      idTerm1,
					Fields:     []PZP26.QueryNode{},
					AllFields:  true,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){name}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1') {name}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1') { name }",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{PZP26.QueryNode{
				Name:  "root[0]",
				Terms: idTerm1,
				Fields: []PZP26.QueryNode{
					PZP26.QueryNode{
						Name:       "name",
						Terms:      PZP26.TermNode{},
						Fields:     []PZP26.QueryNode{},
						AllFields:  false,
						GoesToRoot: false,
					},
				},
				AllFields:  false,
				GoesToRoot: true,
			},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){name},(age>18){name}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Rename:     "",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:  "root[1]",
					Terms: ageTerm,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Rename:     "",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){name}, (age>18){name}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:  "root[1]",
					Terms: ageTerm,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){name} ,(age>18){name}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:  "root[1]",
					Terms: ageTerm,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: `(id='1'){name},
(age>18){name}`,
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:  "root[1]",
					Terms: ageTerm,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: `(id='1'){name},
				(age>18){name}`,
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:  "root[1]",
					Terms: ageTerm,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: `(id='1'){name}
(age>18){name}`,
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:  "root[1]",
					Terms: ageTerm,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "alpha:(id='1'){name},beta:(age>18){name}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:   "root[0]",
					Rename: "alpha",
					Terms:  idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:   "root[0]",
					Rename: "beta",
					Terms:  ageTerm,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "alpha:(id='1'){beta:@city{name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:   "root[0]",
					Rename: "alpha",
					Terms:  idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:   "city",
							Rename: "",
							Terms:  PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Rename:     "",
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: true,
							Index:      "id",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.CanNotRenameField{FieldName: "city", Rename: "beta", Query: "{beta:@city{name}}"},
				PZP26.NoIndexSelected{FieldName: "city"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "alpha:(id='1'){beta:name}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:   "root[0]",
					Rename: "alpha",
					Terms:  idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:   "city",
							Rename: "",
							Terms:  PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Rename:     "",
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: true,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.CanNotRenameField{FieldName: "name", Rename: "beta", Query: "{beta:name}"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "alpha:(id='1'){beta:city{name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:   "root[0]",
					Rename: "alpha",
					Terms:  idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:   "city",
							Rename: "beta",
							Terms:  PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Rename:     "",
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.CanNotRenameField{FieldName: "city", Rename: "beta", Query: "{beta=city{name}}"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){city{$reproduced_schema={id, name}}}, (id='2'){city{$reproduced_schema}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "id",
									AllFields:  false,
									GoesToRoot: false,
								},
								PZP26.QueryNode{
									Name:       "name",
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:  "root[1]",
					Terms: idTerm2,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "id",
									AllFields:  false,
									GoesToRoot: false,
								},
								PZP26.QueryNode{
									Name:       "name",
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "alpha:(id='1'){city{$reproduced_schema={id, name}}}, beta:(id='2'){city{$reproduced_schema}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:   "root[0]",
					Rename: "alpha",
					Terms:  idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "id",
									AllFields:  false,
									GoesToRoot: false,
								},
								PZP26.QueryNode{
									Name:       "name",
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
				PZP26.QueryNode{
					Name:   "root[1]",
					Rename: "beta",
					Terms:  idTerm2,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "id",
									AllFields:  false,
									GoesToRoot: false,
								},
								PZP26.QueryNode{
									Name:       "name",
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){city{$not_existing_schema}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:       "root[0]",
					Terms:      idTerm1,
					Fields:     []PZP26.QueryNode{},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.NotExistingSchema{Name: "not_existing_schema", Query: "city{$not_existing_schema}"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){age,city}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Rename:     "",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
						PZP26.QueryNode{
							Name:       "city",
							Rename:     "",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){age, city}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
						PZP26.QueryNode{
							Name:       "city",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){age ,city}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "name",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
						PZP26.QueryNode{
							Name:       "city",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){city{__all_fields}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "city",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  true,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){city{name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){__all_fields, city{name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: false,
						},
					},
					AllFields:  true,
					GoesToRoot: true,
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){@city}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "city",
							Rename:     "",
							Terms:      PZP26.TermNode{},
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
							Index:      "id",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.FieldHasNotSubSelection{Field: "city"},
				PZP26.NoIndexSelected{FieldName: "city"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){@city{name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:   "city",
							Rename: "",
							Terms:  PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Rename:     "",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
									Index:      "",
								},
							},
							AllFields:  false,
							GoesToRoot: true,
							Index:      "id",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
					Index:      "",
				},
			},
			Warnings: []error{
				PZP26.NoIndexSelected{FieldName: "city"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){@city#name{name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:   "city",
							Rename: "",
							Terms:  PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Rename:     "",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
									Index:      "",
								},
							},
							AllFields:  false,
							GoesToRoot: true,
							Index:      "`name`",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
					Index:      "",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){@city#(`name`){name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:   "city",
							Rename: "",
							Terms:  PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Rename:     "",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
									Index:      "",
								},
							},
							AllFields:  false,
							GoesToRoot: true,
							Index:      "`name`",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
					Index:      "",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){@city#(Cities#`name`){name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:   "city",
							Rename: "",
							Terms:  PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Rename:     "",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
									Index:      "",
								},
							},
							AllFields:  false,
							GoesToRoot: true,
							Index:      "Cities#`name`",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
					Index:      "",
				},
			},
			Warnings: []error{},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){__all_fields, @city{name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: PZP26.TermNode{},
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: true,
							Index:      "id",
						},
					},
					AllFields:  true,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.NoIndexSelected{FieldName: "city"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){@city(population>150000)}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:       "city",
							Terms:      populationTerm,
							Fields:     []PZP26.QueryNode{},
							AllFields:  false,
							GoesToRoot: false,
							Index:      "id",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.FieldHasNotSubSelection{Field: "city"},
				PZP26.NoIndexSelected{FieldName: "city"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){@city(population>150000){name}}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:  "root[0]",
					Terms: idTerm1,
					Fields: []PZP26.QueryNode{
						PZP26.QueryNode{
							Name:  "city",
							Terms: populationTerm,
							Fields: []PZP26.QueryNode{
								PZP26.QueryNode{
									Name:       "name",
									Terms:      PZP26.TermNode{},
									Fields:     []PZP26.QueryNode{},
									AllFields:  false,
									GoesToRoot: false,
								},
							},
							AllFields:  false,
							GoesToRoot: true,
							Index:      "id",
						},
					},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{
				PZP26.NoIndexSelected{FieldName: "city"},
			},
			Errors: []error{},
		},
	},
	QueryTest{
		Query: "(id='1'){}",
		Expected: QueryExpectation{
			Result: []PZP26.QueryNode{
				PZP26.QueryNode{
					Name:       "root[0]",
					Terms:      idTerm1,
					Fields:     []PZP26.QueryNode{},
					AllFields:  false,
					GoesToRoot: true,
				},
			},
			Warnings: []error{PZP26.FieldHasNotSubSelection{Field: "root[0]"}},
			Errors:   []error{},
		},
	},
	QueryTest{
		Query: "{name}",
		Expected: QueryExpectation{
			Result:   []PZP26.QueryNode{},
			Warnings: []error{},
			Errors: []error{
				PZP26.MissingRootSelection{Query: "{name}"},
			},
		},
	},
	QueryTest{
		Query: "(){name}",
		Expected: QueryExpectation{
			Result:   []PZP26.QueryNode{},
			Warnings: []error{},
			Errors: []error{
				PZP26.MissingRootSelection{Query: "(){name}"},
			},
		},
	},
}
