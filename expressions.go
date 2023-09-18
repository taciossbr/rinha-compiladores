package main

import "strings"

type Location struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Filename string `json:"filename"`
}

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

type TermMeta struct {
	Kind     TermKind `json:"kind"`
	Location Location `json:"location"`
	// Value    *Term    `json:"value"`
	// Literal  interface{}
}

type Term interface {
	Accept(v Visitor) any
	GetKind() TermKind
	GetLocation() Location
}

type Visitor interface {
	// TODO ver algo melhor que o any, tem templates?
	visitPrint(t RinPrint) any
	visitString(t RinString) any
	visitInteger(t RinInteger) any
}

type RinString struct {
	Meta  TermMeta
	Value string
}

func (t *RinString) Accept(v Visitor) any {
	return v.visitString(*t)
}
func (t *RinString) GetKind() TermKind {
	return t.Meta.Kind
}
func (t *RinString) GetLocation() Location {
	return t.Meta.Location
}

type RinInteger struct {
	Meta  TermMeta
	Value int
}

func (t *RinInteger) Accept(v Visitor) any {
	return v.visitInteger(*t)
}
func (t *RinInteger) GetKind() TermKind {
	return t.Meta.Kind
}
func (t *RinInteger) GetLocation() Location {
	return t.Meta.Location
}

type RinPrint struct {
	Meta  TermMeta
	Value Term
}

func (t *RinPrint) Accept(v Visitor) any {
	return v.visitPrint(*t)
}
func (t *RinPrint) GetKind() TermKind {
	return t.Meta.Kind
}
func (t *RinPrint) GetLocation() Location {
	return t.Meta.Location
}

func ParseTermKind(str string) TermKind {
	return termKindMapper[strings.ToLower(str)]
}

func parseJsonToNodes(expression map[string]any) Term {
	kind := ParseTermKind(expression["kind"].(string))
	meta := TermMeta{
		Kind:     kind,
		Location: makeLocation(expression["location"].(map[string]any)),
	}

	switch kind {
	case StringTermKind:
		return &RinString{
			Meta:  meta,
			Value: expression["value"].(string),
		}
	case IntTermKind:
		return &RinInteger{
			Meta:  meta,
			Value: int(expression["value"].(float64)),
		}
	case PrintTermKind:
		return &RinPrint{
			Meta:  meta,
			Value: parseJsonToNodes(expression["value"].(map[string]any)),
		}
	}
	return nil
}

func makeLocation(data map[string]any) Location {
	location := Location{
		Start:    int(data["start"].(float64)),
		End:      int(data["end"].(float64)),
		Filename: data["filename"].(string),
	}
	return location
}
