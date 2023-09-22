package expressions

import "strings"

type TermKind string

const (
	BinaryTermKind   TermKind = "Binary"
	BoolTermKind     TermKind = "Bool"
	CallTermKind     TermKind = "Call"
	FunctionTermKind TermKind = "Function"
	FirstTermKind    TermKind = "First"
	IfTermKind       TermKind = "If"
	IntTermKind      TermKind = "Int"
	LetTermKind      TermKind = "Let"
	PrintTermKind    TermKind = "Print"
	SecondTermKind   TermKind = "Second"
	StringTermKind   TermKind = "String"
	TupleTermKind    TermKind = "Tuple"
	VarTermKind      TermKind = "Var"
)

var termKindMapper = map[string]TermKind{
	"binary":   BinaryTermKind,
	"bool":     BoolTermKind,
	"call":     CallTermKind,
	"first":    FirstTermKind,
	"function": FunctionTermKind,
	"if":       IfTermKind,
	"int":      IntTermKind,
	"let":      LetTermKind,
	"print":    PrintTermKind,
	"second":   SecondTermKind,
	"str":      StringTermKind,
	"tuple":    TupleTermKind,
	"var":      VarTermKind,
}

func ParseTermKind(str string) TermKind {
	return termKindMapper[strings.ToLower(str)]
}
