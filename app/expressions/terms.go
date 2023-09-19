package expressions

type Visitor interface {
	VisitString(t RinString) any
	VisitInteger(t RinInteger) any
	VisitPrint(t RinPrint) any
}

type RinString struct {
	Meta TermMeta
	Value string
}

func (t *RinString) Accept(v Visitor) any {
	return v.VisitString(*t)
}

func (t *RinString) GetKind() TermKind {
	return t.Meta.Kind
}

func (t *RinString) GetLocation() Location {
	return t.Meta.Location
}

type RinInteger struct {
	Meta TermMeta
	Value int
}

func (t *RinInteger) Accept(v Visitor) any {
	return v.VisitInteger(*t)
}

func (t *RinInteger) GetKind() TermKind {
	return t.Meta.Kind
}

func (t *RinInteger) GetLocation() Location {
	return t.Meta.Location
}

type RinPrint struct {
	Meta TermMeta
	Value Term
}

func (t *RinPrint) Accept(v Visitor) any {
	return v.VisitPrint(*t)
}

func (t *RinPrint) GetKind() TermKind {
	return t.Meta.Kind
}

func (t *RinPrint) GetLocation() Location {
	return t.Meta.Location
}

