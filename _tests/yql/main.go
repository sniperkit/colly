package main

import (
	"encoding/json"
	"fmt"

	// external
	"github.com/caibirdme/yql"

	// colly - plugins
	pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

/*
	Refs:
	- github.com/antlr/antlr4
*/

var jsonString = `
  {
    "name": "enny",
    "gender": "f",
    "age": 36,
    "hobby": null,
    "skills": [
      "IC",
      "Electric design",
      "Verification"
    ]
  }
`

func main() {
	fmt.Println("YQL parser start...")

	exampleRuleAdvancedQuery()
	// exampleStream()
	// exampleMatch()
	// exampleRule()

}

var OnTABSelectors map[string]*DatasetSelect = make(map[string]*DatasetSelect, 4)

func exampleRuleAdvancedQuery() {
	// with rows 7,1,9,5,3
	rawYQL := `rows in (7,1,9,5,3)`
	result, _ := yql.Rule(rawYQL)
	pp.Println(result)

	// from rows 3 to 10
	rawYQL = `rows in [3:10]`
	result, _ = yql.Rule(rawYQL)
	pp.Println(result)

	// from rows 2 to end
	rawYQL = `SELECT COLUMNS('id', 'fullname', 'owner_login') WHERE ROWS=[2:]`
	result, _ = yql.Rule(rawYQL)
	pp.Println(result)

	// from rows 3 to 10
	rawYQL = `SELECT COLUMNS('id', 'fullname', 'owner_login') WHERE ROWS=[3:10]`
	result, _ = yql.Rule(rawYQL)
	pp.Println(result)

	// from 1 to end with cap at 15
	rawYQL = `SELECT COLUMNS('id', 'fullname', 'owner_login') WHERE ROWS=[1::15]`
	result, _ = yql.Rule(rawYQL)
	pp.Println(result)

	// with rows 7,1,9,5,3
	rawYQL = `SELECT COLUMNS('id', 'fullname', 'owner_login') WHERE ROWS IN (7,1,9,5,3)`
	result, _ = yql.Rule(rawYQL)
	pp.Println(result)

}

type DatasetSelect struct {
	Select     string
	PatternURL string
	PluckURL   string
}

func exampleRuleQuery() {
	rawYQL := `score in (7,1,9,5,3)`
	result, _ := yql.Rule(rawYQL)
	pp.Println(result)

}

func exampleStream() {
	rawYQL := `name='enny'`
	var temp map[string]interface{}
	json.Unmarshal([]byte(jsonString), &temp)
	result, _ := yql.Match(rawYQL, temp)
	pp.Println(result)
}

func exampleMatch() {

	rawYQL := `name='deen' and age>=23 and (hobby in ('soccer', 'swim') or score>90)`
	result, _ := yql.Match(rawYQL, map[string]interface{}{
		"name":  "deen",
		"age":   23,
		"hobby": "basketball",
		"score": int64(100),
	})
	pp.Println(result)

	rawYQL = `score ∩ (7,1,9,5,3)`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 100, 200},
	})
	pp.Println(result)

	rawYQL = `score in (7,1,9,5,3)`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 5, 2},
	})
	pp.Println(result)

	rawYQL = `score.sum() > 10`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int{1, 2, 3, 4, 5},
	})
	pp.Println(result)
	//Output:
	//true
	//true
	//false
	//true

}

func exampleRule() {

	rawYQL := `name='deen' and age>=23 and (hobby in ('soccer', 'swim') or score>90)`
	ruler, _ := yql.Rule(rawYQL)
	result, _ := ruler.Match(map[string]interface{}{
		"name":  "deen",
		"age":   23,
		"hobby": "basketball",
		"score": int64(100),
	})
	pp.Println(result)

	result, _ = ruler.Match(map[string]interface{}{
		"name":  "deen",
		"age":   23,
		"hobby": "basketball",
		"score": int64(90),
	})
	pp.Println(result)
	//Output:
	//true
	//false

}
