package PZP26

type UnexpectedTermTypeValue struct {
	Term     string
	Got      string
	Expected string
}

func (e UnexpectedTermTypeValue) Error() string {
	return "TODO"
}

type UncomparableValues struct {
	Test      string
	TypeLeft  string
	TypeRight string
}

func (e UncomparableValues) Error() string {
	return "TODO"
}

type UnclosedTerm struct {
	Term string
}

func (e UnclosedTerm) Error() string {
	return "TODO"
}

type Unclosed struct {
	Text string
}

func (e Unclosed) Error() string {
	return "TODO"
}

type UnspecifiedTerm struct {
	Term string
}

func (e UnspecifiedTerm) Error() string {
	return "TODO"
}

type FieldHasNotSubSelection struct {
	Field string
}

func (e FieldHasNotSubSelection) Error() string {
	return "TODO"
}

type MissingRootSelection struct {
	Query string
}

func (e MissingRootSelection) Error() string {
	return "TODO"
}

type NoItemWasFoundWithTerm struct {
	Term string
}

func (e NoItemWasFoundWithTerm) Error() string {
	return "TODO"
}

type FieldNonExistingOnData struct {
	FieldName string
	FieldPath string
	Data      interface{}
}

func (e FieldNonExistingOnData) Error() string {
	return "TODO"
}

type CanNotRenameField struct {
	FieldName string
	Rename    string
	Query     string
}

func (e CanNotRenameField) Error() string {
	return "TODO"
}

type NotExistingSchema struct {
	Name  string
	Query string
}

func (e NotExistingSchema) Error() string {
	return "TODO"
}

type NoIndexSelected struct {
	FieldName string
}

func (e NoIndexSelected) Error() string {
	return "TODO"
}

type NotAnArray struct {
	Value interface{}
}

func (e NotAnArray) Error() string {
	return "TODO"
}
