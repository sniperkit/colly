package PZP26

import "errors"

type TermOperator interface {
	Compute(left interface{}, right interface{}) (bool, []error, []error)
}

type Equals struct{}

func (op Equals) Compute(left interface{}, right interface{}) (bool, []error, []error) {
	return false, []error{}, []error{errors.New("TODO")}
}

type NotEquals struct{}

func (op NotEquals) Compute(left interface{}, right interface{}) (bool, []error, []error) {
	return false, []error{}, []error{errors.New("TODO")}
}

// >
type StrictGreater struct{}

func (op StrictGreater) Compute(left interface{}, right interface{}) (bool, []error, []error) {
	return false, []error{}, []error{errors.New("TODO")}
}

// <
type StrictLower struct{}

func (op StrictLower) Compute(left interface{}, right interface{}) (bool, []error, []error) {
	return false, []error{}, []error{errors.New("TODO")}
}

type Or struct{}

func (op Or) Compute(left interface{}, right interface{}) (bool, []error, []error) {
	return false, []error{}, []error{errors.New("TODO")}
}

type And struct{}

func (op And) Compute(left interface{}, right interface{}) (bool, []error, []error) {
	return false, []error{}, []error{errors.New("TODO")}
}

type In struct{}

func (op In) Compute(left interface{}, right interface{}) (bool, []error, []error) {
	return false, []error{}, []error{errors.New("TODO")}
}
