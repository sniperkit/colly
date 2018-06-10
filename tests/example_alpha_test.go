package PZP26_test

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Jimskapt/PZP26"
)

func ExampleMain() {
	warnList := []error{}
	errList := []error{}

	// Classic order : ParseQuery > Resolver > Selector > Display (JSON, XML, ...)

	queries := `first_set:(id='1'){id,name,@city{name}}, second_set:(__type='People' AND age>18){id, age, name}`

	fmt.Println("Queries :", queries)

	// ParseQuery : from string to easier-to-use objects

	parsedQuery, warns, errs := PZP26.ParseQuery(queries)
	warnList = append(warnList, warns...)
	errList = append(errList, errs...)

	// Resolver : ask you your data, requested by the query

	resolvedData := []interface{}{}
	if len(errs) == 0 {
		resolver := PZP26.Resolver{
			ResolveFunc: func(term PZP26.TermNode) (interface{}, []error, []error, bool) {
				for _, item := range STORE {
					if PZP26.DataRangeSolver(term, item) {
						return item, []error{}, []error{}, true
					}
				}

				return nil, []error{errors.New("No one item was found with term" + term.String())}, []error{}, false
			},
		}
		resolvedData, warns, errs = resolver.ResolveEach(parsedQuery)
		warnList = append(warnList, warns...)
		errList = append(errList, errs...)
	}

	// Selector : hide potentially not requested data (fields and sub-fields) by the query

	selected := map[string]interface{}{}
	if len(errs) == 0 {
		selected, warns, errs = PZP26.Selector(parsedQuery, resolvedData)
		warnList = append(warnList, warns...)
		errList = append(errList, errs...)
	}

	// Display : display the data as you want (hide errors, warnings, XML instead JSON, ...)

	result := map[string]interface{}{
		"data":     selected,
		"warnings": warnList,
		"errors":   errList,
	}

	parsed, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Result :", string(parsed))

	// Output:
	// Queries : first_set:(id='1'){id,name,@city{name}}, second_set:(__type='People' AND age>18){id, age, name}
	// Result :
	// {
	//   "data": {
	// 		"first_set": {
	// 			"id": "1",
	// 			"name": "Smith",
	// 			"city": "JKB21"
	// 		},
	// 		"JKB21": {
	// 			"name": "New York"
	// 		},
	// 		"second_set": [
	// 			{
	// 				"id": "1",
	// 				"age": 21,
	// 				"name": "Smith"
	// 			},
	// 			{
	// 				"id": "2",
	// 				"age": 52,
	// 				"name": "John"
	// 			}
	// 		]
	//   },
	//   "warnings": [],
	//   "errors": []
	// }
}
