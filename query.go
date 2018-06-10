package PZP26

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type QueryNode struct {
	Name       string
	Rename     string
	Terms      TermNode
	Fields     []QueryNode
	AllFields  bool
	GoesToRoot bool
	Index      string
	/*
		Rename     string   `json:"-"`
		Terms      TermNode `json:"-"`
		Fields     []QueryNode
		AllFields  bool   `json:"-"`
		GoesToRoot bool   `json:"-"`
		Index      string `json:"-"`
	*/
}

func matching(open rune, close rune, text string) (string, error) {
	level := 0

	for i, character := range text {
		// fmt.Println(string(character))
		if character == open {
			level++
		} else if character == close {
			if level > 0 {
				level--
			} else {
				return text[:i], nil
			}
		}
	}

	return "", Unclosed{Text: text}
}

func ParseQuery(query string) ([]QueryNode, []error, []error) {
	result := []QueryNode{}
	warnings := []error{}
	errors := []error{}

	available := query
	rootID := 0

	firstParenthesis := strings.Index(available, "(")
	for firstParenthesis > 0 {
		fmt.Printf("\n%#v\n", available)

		//////////////////////////////////////////////////////////////////

		termsString, err := matching('(', ')', available[firstParenthesis+1:])

		if err != nil {
			errors = append(errors, err)
			panic(errors)
			return result, warnings, errors
		}

		//////////////////////////////////////////////////////////////////

		beforeTerms := available[:firstParenthesis]
		afterTerms := available[firstParenthesis+1+len(termsString)+1:]

		//////////////////////////////////////////////////////////////////

		fieldName := ""
		renameRegex := regexp.MustCompile("([^(,\n )]+):$")
		find := renameRegex.FindString(strings.TrimSpace(beforeTerms))
		if find != "" {
			fieldName = find[:len(find)-1]
		}

		//////////////////////////////////////////////////////////////////

		firstBrace := strings.Index(afterTerms, "{")
		if firstBrace == -1 {
			errors = append(errors, FieldHasNotSubSelection{Field: fieldName})
			panic(errors)
			return result, warnings, errors
		}

		fieldsString, err := matching('{', '}', afterTerms[firstBrace+1:])
		if err != nil {
			errors = append(errors, err)
			panic(errors)
			return result, warnings, errors
		}
		afterTerms = afterTerms[:firstBrace] + afterTerms[firstBrace+1+len(fieldsString)+1:]

		//////////////////////////////////////////////////////////////////

		fmt.Printf("fieldName = <%v>\n", strings.TrimSpace(fieldName))
		fmt.Printf("termsString = <%v>\n", strings.TrimSpace(termsString))
		fmt.Printf("fieldsString = <%v>\n", strings.TrimSpace(fieldsString))

		newItem := QueryNode{}
		newItem.Name = "root[" + strconv.Itoa(rootID) + "]"
		newItem.Rename = strings.TrimSpace(fieldName)

		//////////////////////////////////////////////////////////////////

		terms, warns, errs := ParseTerm(strings.TrimSpace(termsString))
		if errs != nil && len(errs) > 0 {
			errors = append(errors, errs...)
			panic(errors)
			return result, warnings, errors
		}
		if warns != nil && len(errs) > 0 {
			warnings = append(warnings, warns...)
		}
		newItem.Terms = terms

		//////////////////////////////////////////////////////////////////

		fields, warns, errs := ParseFields(strings.TrimSpace(fieldsString))
		if len(errs) > 0 {
			errors = append(errors, errs...)
			panic(errors)
			return result, warnings, errors
		}
		if len(warns) > 0 {
			warnings = append(warnings, warns...)
		}
		newItem.Fields = fields

		//////////////////////////////////////////////////////////////////

		result = append(result, newItem)

		//available = beforeTerms + afterTerms
		available = afterTerms

		firstParenthesis = strings.Index(available, "(")

		rootID++
	}

	fmt.Printf("\n%#v\n", available)

	return result, warnings, errors
}

func ParseFields(data string) ([]QueryNode, []error, []error) {
	result := []QueryNode{}
	warnings := []error{}
	errors := []error{}

	available := data

	commaSplit := strings.Split(available, ",")
	for _, childField := range commaSplit {
		childField = strings.TrimSpace(childField)

		firstBrace := strings.Index(available, "{")
		if firstBrace > 0 {
			braceString, err := matching('{', '}', available[firstBrace+1:])
			if err != nil {
				errors = append(errors, err)
				panic(errors)
				return result, warnings, errors
			}

			afterBrace := available[firstBrace+1+len(braceString)+1:]

			//////////////////////////////////////////////////////////////////

			subs, warns, errs := ParseFields(available[firstBrace+1 : firstBrace+1+len(braceString)])
			warnings = append(warnings, warns...)
			if len(errs) > 0 {
				errors = append(errors, errs...)
				panic(errors)
				return result, warnings, errors
			}

			//////////////////////////////////////////////////////////////////

			parentName := ""
			nameRegex := regexp.MustCompile("([^(,\n)]+)$")
			find := nameRegex.FindString(strings.TrimSpace(available[:firstBrace]))
			if find != "" {
				parentName = find
			}

			//////////////////////////////////////////////////////////////////

			newItem := QueryNode{
				Name:   parentName,
				Fields: subs,
			}

			result = append(result, newItem)

			//////////////////////////////////////////////////////////////////

			available = afterBrace

			firstBrace = strings.Index(available, "{")

			result = append(result, newItem)
		} else {
			newItem := QueryNode{
				Name: childField,
			}

			result = append(result, newItem)

			fmt.Printf("available = %v\n", available)
			fmt.Printf("childField = %v\n", childField)

			if len(available) > len(childField) {
				available = available[len(childField):]
			}
		}
	}

	fmt.Printf("ParseFields(\"%v\") = %#v\n", data, result)

	return result, warnings, errors
}
