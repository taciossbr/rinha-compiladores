package expressions

import "strings"

type TermKind string

const (
	PrintTermKind  TermKind = "Print"
	StringTermKind TermKind = "String"
	IntTermKind    TermKind = "Int"
)

var termKindMapper = map[string]TermKind{
	"print": PrintTermKind,
	"str":   StringTermKind,
	"int":   IntTermKind,
}

func ParseTermKind(str string) TermKind {
	return termKindMapper[strings.ToLower(str)]
}
