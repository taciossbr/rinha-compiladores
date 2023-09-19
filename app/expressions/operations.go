package expressions

type RinOp int

const (
	AddOp RinOp = iota
	SubOp
	MulOp
	DivOp
	RemOp
	EqOp
	NeqOp
	LtOp
	GtOp
	LteOp
	GteOp
	AndOp
	OrOp
)

var operationsMapper = map[string]RinOp{
	"Add": AddOp,
	"Sub": SubOp,
	"Mul": MulOp,
	"Div": DivOp,
	"Rem": RemOp,
	"Eq":  EqOp,
	"Neq": NeqOp,
	"Lt":  LtOp,
	"Gt":  GtOp,
	"Lte": LteOp,
	"Gte": GteOp,
	"And": AndOp,
	"Or":  OrOp,
}

func ParseOp(str string) RinOp {
	return operationsMapper[str]
}
