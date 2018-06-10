package main

import (
	"encoding/json"
	"fmt"

	"github.com/Jimskapt/PZP26"
)

func main() {
	fmt.Println("=== PZP26 standalone mode ===")

	query := `first_set:(id IN ('1', '2','3')) {$schema={id, name,age,@city#name{id,name}}}`
	query += `
	`
	query += `(id='1'){
		id
		name
	}`
	query += `, `
	query += `second_set:( __type='People' AND age>18 ){ $schema }`
	query += `,(id='2' ){ id, name}`

	a, b, c := PZP26.ParseQuery(query)

	A, _ := json.MarshalIndent(a, "", "\t")

	fmt.Printf("%v\n%v\n%v\n", string(A), b, c)
}
