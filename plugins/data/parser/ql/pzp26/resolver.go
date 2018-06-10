package PZP26

import "errors"

type Resolver struct {
	ResolveFunc func(term TermNode) (result interface{}, warnings []error, errors []error, continueRecursion bool)
}

func (r Resolver) ResolveEach(queries []QueryNode) ([]interface{}, []error, []error) {
	resInterface := []interface{}{}
	resWarnings := []error{}
	resErrors := []error{errors.New("TODO")}

	for _, query := range queries {
		// TODO : do recursive calls from query sub-fields
		res, warns, errs, _ := r.ResolveFunc(query.Terms)

		resInterface = append(resInterface, res)
		resWarnings = append(resWarnings, warns...)
		resErrors = append(resErrors, errs...)
	}

	// TODO : do recursive calls to ResolveFunc on "res" for nested fields (@city{name}) if it is an "id value".
	// Ignore it if it is a struct because data is already there.

	return resInterface, resWarnings, resErrors
}

func DataRangeSolver(term TermNode, item interface{}) bool {
	// TODO
	return false
}
