package main

import (
	"fmt"
	"regexp"

	regex "github.com/dakerfp/re"
	cregex "github.com/mingrammer/commonregex"
	// pluck "github.com/schollz/pluck/plucker"

	// colly - plugins
	pp "github.com/sniperkit/colly/plugins/app/debug/pp"
)

/*
	Refs:
	- https://regex-golang.appspot.com/assets/html/index.html
	-
*/

const (
	BLOCK = `([A-Za-z]+[0-9]+)\:([A-Za-z]+[0-9]+)`
	CELL  = `([A-Za-z]+)([0-9]+)`
	COLS  = `([A-Za-z]+)\:([A-Za-z]+)`
	ROWS  = `([0-9]+)\:([0-9]+)`
	COL   = `([A-Za-z]+)`
	ROW   = `([0-9]+)`
)

const selectQueryRegex = `((col|cols|rows|row))\[((:\d+)|(\d+\:)|(\d+\:\d+)|(\:)|(\d+(,\d+)*))(\])`

var ass *regexp.Regexp

func main() {

	ass = regexp.MustCompile(selectQueryRegex)
	getSelection(`cols[0:5],rows[1:7]`)

	getSelection(`cols[:],rows[:]`)

	getSelection(`cols[1,2],rows[:]`)

}

func getSelection(query string) {
	res := ass.FindAllStringSubmatch(query, 2)
	fmt.Println("\n------------------------------------------------------------")
	fmt.Println("extract for `cols[0:1],rows[1:7]`")

	// pp.Println("res=", res)
	/*
		var isRowSlice, isRowRange, isColSlice, isColRange bool
		var cLowerInt, cUpperInt, rLower, rUpper int
		var cLowerStr, cUpperStr string // column names
		var cSliceInt, rSliceInt []int
		var cSliceList, rSliceList []string
	*/

	for _, r := range res {
		// pp.Println("match=", r)
		fmt.Println("action=", r[2], "range=", r[3])
	}

}

func re() {
	reTest := regex.Regex(
		regex.Group("dividend", regex.Digits),
		regex.Then("/"),
		regex.Group("divisor", regex.Digits),
	)

	m := reTest.FindStringSubmatch("4/3")
	fmt.Println(m[1]) // > 4
	fmt.Println(m[2]) // > 3
}

func test_cregex() {

	dateList := cregex.Date(TEXT)
	pp.Println("regex.Date=", dateList)
	// ['Jan 9th 2012']

	timeList := cregex.Time(TEXT)
	pp.Println("regex.Time=", timeList)
	// ['5:00PM', '4:00']

	linkList := cregex.Links(TEXT)
	pp.Println("regex.Links=", linkList)
	// ['www.linkedin.com', 'harold.smith@gmail.com']

	phoneList := cregex.PhonesWithExts(TEXT)
	pp.Println("regex.PhonesWithExts=", phoneList)
	// ['(519)-236-2723x341']

	emailList := cregex.Emails(TEXT)
	pp.Println("regex.Emails=", emailList)
	// ['harold.smith@gmail.com']
}

func test_regexp() {
	/*
		fmt.Println("\nextract for `cols[:],rows[1:7]`")
		res = ass.FindAllStringSubmatch(`cols[:],rows[1:7]`, -1)
		for _, r := range res {
			// pp.Println("match=", r)
			fmt.Println("action=", r[2], "range=", r[3])
		}

		fmt.Println("\nextract for `cols[:10],row[1,5,7]`")
		res = ass.FindAllStringSubmatch(`cols[:10],row[1,5,7]`, -1)
		for _, r := range res {
			// pp.Println("match=", r)
			fmt.Println("action=", r[2], "range=", r[3])
		}

		fmt.Println("\nextract for `cols[:1],rows[1:7]`")
		res = ass.FindAllStringSubmatch(`cols[:1],rows[1:7]`, -1)
		for _, r := range res {
			fmt.Println("action=", r[2], "range=", r[3])
		}

		fmt.Println("\nextract for `cols[:],rows[1:7]`")
		res = ass.FindAllStringSubmatch(`cols[:],rows[1:7]`, -1)
		for _, r := range res {
			fmt.Println("action=", r[2], "range=", r[3])
		}
	*/
}
