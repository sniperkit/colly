package PZP26

type TermNode struct {
	Right    interface{}
	Operator TermOperator
	Left     interface{}
}

func (t TermNode) String() string {
	return "TODO"
}

func ParseTerm(term string) (TermNode, []error, []error) {
	//return TermNode{}, []error{}, []error{errors.New("TODO")}
	return TermNode{}, []error{}, []error{}
}
