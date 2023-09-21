package expressions

import "strings"

type TermKind string

const (
	BinaryTermKind TermKind = "Binary"
	BoolTermKind   TermKind = "Bool"
	IntTermKind    TermKind = "Int"
	LetTermKind    TermKind = "Let"
	PrintTermKind  TermKind = "Print"
	StringTermKind TermKind = "String"
	VarTermKind    TermKind = "Var"
)

var termKindMapper = map[string]TermKind{
	"binary": BinaryTermKind,
	"bool":   BoolTermKind,
	"int":    IntTermKind,
	"let":    LetTermKind,
	"print":  PrintTermKind,
	"str":    StringTermKind,
	"var":    VarTermKind,
}

func ParseTermKind(str string) TermKind {
	return termKindMapper[strings.ToLower(str)]
}
