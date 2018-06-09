package main

import (
	"fmt"
	"regexp"
	"strings"

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

const (
	selectQueryRegexIdx   = `((col|cols|rows|row))\[((:\d+)|(\d+\:)|(\d+\:\d+)|(\:)|(\d+(,\d+)*))(\])`
	selectQueryRegexAlpha = `((col|cols|rows|row))\[((:[a-zA-Z0-9-_\"\']+)|([a-zA-Z0-9-_\"\']+\:)|([a-zA-Z0-9-_\"\']+\:[a-zA-Z0-9-_\"\']+)|([a-zA-Z0-9-_\"\']+(,[a-zA-Z0-9-_\"\']+)*)|(\:)|([a-zA-Z0-9-_\"\']+(,[a-zA-Z0-9-_\"\']+)*))(\])`
	// Ref: https://stackoverflow.com/questions/5988228/how-to-create-a-regex-for-accepting-only-alphanumeric-characters
	selectHeadersUnicode = `^[\\p{L}0-9]*$`
	// ^[a-zA-Z0-9_\"\']*$
)

var ass *regexp.Regexp

func main() {

	ass = regexp.MustCompile(selectQueryRegexAlpha)
	getSelection(`cols[0:5],rows[1:7]`)

	getSelection(`cols[:],rows[:]`)

	getSelection(`cols[1,2],rows[:]`)

	getSelection(`col[name, full_name],row[4]`)

	getSelection(`col["name", "full_name"],rows[1,5,7,8]`)

	getSelection(`col["name", "full_name"],rows[:10]`)

	getSelection(`rows[:10],col["name", "full_name"]`)

	getSelection(`rows[1,10],col["name", "fullname"]`)

	getSelection(`rows[1,10],col["id", "name", "full_name", "description", "language", "stargazers_count", "forks_count"]`)

}

type Selection struct {
	ColumnNames   []string `json:"column_names"`
	ColumnIndices []int    `json:"column_keys"`
	RowIndices    []int    `json:"column_keys"`
}

func getSelection(query string) {
	query = strings.Replace(query, " ", "", -1)
	query = strings.Replace(query, "\"", "", -1)

	selectors := ass.FindAllStringSubmatch(query, 2)
	fmt.Println("\n------------------------------------------------------------")
	fmt.Printf("extract for `%s\n", query)

	// Check if results are < 1
	if len(res) < 1 {
		fmt.Println("An error occured while parsing the selection query syntax.")
	}

	// Loop over query selection parameters
	for _, selectionString := range selectors {

		// extract selection keys from
		var keys []string

		pp.Printf("action=\"%s\", selectionString=`%s`", selectStr[2], selectStr[3])

		// Check if list of ROW/COL keys
		var isList bool
		selectList := strings.Split(selectionString[3], ",")
		if selectList > 0 {
			isList = true

		}

		/*
			// Check if slice of ROW/COL keys
			var isSlice bool
			selectSlice := strings.Split(selectStr[3], ":")
			if selectSlice > 0 {
				isSlice = true
			}
		*/

		// Check if range of ROW/COL keys
		var isRange bool
		selectSlice := strings.Split(selectionString[3], ":")
		if selectSlice > 0 {
			isSlice = true
		}

		// unique key
		if !isSlice && !isList {
			key
		}

		pp.Println("selectParts", selectParts)
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

var headers = []string{
	"html_url",
	"keys_url",
	"pulls_url",
	"milestones_url",
	"subscription_url",
	"compare_url",
	"has_downloads",
	"id",
	"git_refs_url",
	"statuses_url",
	"stargazers_url",
	"git_url",
	"default_branch",
	"branches_url",
	"pushed_at",
	"watchers_count",
	"teams_url",
	"notifications_url",
	"labels_url",
	"name",
	"fork",
	"commits_url",
	"comments_url",
	"full_name",
	"issue_comment_url",
	"merges_url",
	"size",
	"license",
	"clone_url",
	"language",
	"owner",
	"private",
	"events_url",
	"languages_url",
	"contributors_url",
	"contents_url",
	"homepage",
	"has_pages",
	"description",
	"tags_url",
	"blobs_url",
	"git_tags_url",
	"issues_url",
	"stargazers_count",
	"has_wiki",
	"forks_count",
	"forks",
	"url",
	"releases_url",
	"created_at",
	"collaborators_url",
	"has_issues",
	"node_id",
	"forks_url",
	"subscribers_url",
	"downloads_url",
	"deployments_url",
	"svn_url",
	"mirror_url",
	"updated_at",
	"ssh_url",
	"has_projects",
	"hooks_url",
	"archived",
	"open_issues_count",
	"watchers",
	"issue_events_url",
	"assignees_url",
	"trees_url",
	"git_commits_url",
	"archive_url",
	"open_issues",
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
