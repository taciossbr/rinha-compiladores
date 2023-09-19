package expressions

import "strings"

type TermKind string

const (
	BinaryTermKind TermKind = "Binary"
	BoolTermKind   TermKind = "Bool"
	IntTermKind    TermKind = "Int"
	PrintTermKind  TermKind = "Print"
	StringTermKind TermKind = "String"
)

var termKindMapper = map[string]TermKind{
	"binary": BinaryTermKind,
	"bool":   BoolTermKind,
	"int":    IntTermKind,
	"print":  PrintTermKind,
	"str":    StringTermKind,
}

func ParseTermKind(str string) TermKind {
	return termKindMapper[strings.ToLower(str)]
}
