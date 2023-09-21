package expressions

type Visitor interface {
	VisitBinary(t RinBinary) any
	VisitBool(t RinBool) any
	VisitInteger(t RinInteger) any
	VisitLet(t RinLet) any
	VisitPrint(t RinPrint) any
	VisitString(t RinString) any
	VisitVar(t RinVar) any
}

type RinBinary struct {
	Meta TermMeta
	Left Term
	Operation RinOp
	Right Term
}

func (t *RinBinary) Accept(v Visitor) any {
	return v.VisitBinary(*t)
}

func (t *RinBinary) GetKind() TermKind {
	return t.Meta.Kind
}

func (t *RinBinary) GetLocation() Location {
	return t.Meta.Location
}

type RinBool struct {
	Meta TermMeta
	Value bool
}

func (t *RinBool) Accept(v Visitor) any {
	return v.VisitBool(*t)
}

func (t *RinBool) GetKind() TermKind {
	return t.Meta.Kind
}

func (t *RinBool) GetLocation() Location {
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

type RinLet struct {
	Meta TermMeta
	Name Parameter
	Value Term
	Next Term
}

func (t *RinLet) Accept(v Visitor) any {
	return v.VisitLet(*t)
}

func (t *RinLet) GetKind() TermKind {
	return t.Meta.Kind
}

func (t *RinLet) GetLocation() Location {
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

type RinVar struct {
	Meta TermMeta
	Text string
}

func (t *RinVar) Accept(v Visitor) any {
	return v.VisitVar(*t)
}

func (t *RinVar) GetKind() TermKind {
	return t.Meta.Kind
}

func (t *RinVar) GetLocation() Location {
	return t.Meta.Location
}

