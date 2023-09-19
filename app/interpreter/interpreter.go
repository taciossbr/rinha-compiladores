package interpreter

import (
	"fmt"
	exp "rinha/expressions"
)

type Interpreter struct{}

func (i *Interpreter) Interpret(term exp.Term) any {
	return term.Accept(i)
}

func (*Interpreter) VisitBool(t exp.RinBool) any {
	return t.Value
}

// VisitBinary implements exp.Visitor.
func (i *Interpreter) VisitBinary(t exp.RinBinary) any {
	left := i.Interpret(t.Left)
	right := i.Interpret(t.Right)

	switch t.Operation {
	// TODO
	// case exp.AddOp: return left. + right
	case exp.SubOp:
		return left.(int) - right.(int)
	case exp.MulOp:
		return left.(int) * right.(int)
	case exp.DivOp:
		return left.(int) / right.(int)
	case exp.RemOp:
		return left.(int) % right.(int)
	// case exp.EqOp: return left.(int) == right.(int)
	// case exp.NeqOp: return left.(int) != right.(int)
	case exp.LtOp:
		return left.(int) < right.(int)
	case exp.GtOp:
		return left.(int) > right.(int)
	case exp.LteOp:
		return left.(int) <= right.(int)
	case exp.GteOp:
		return left.(int) >= right.(int)
	case exp.AndOp:
		return left.(bool) && right.(bool)
	case exp.OrOp:
		return left.(bool) || right.(bool)
	}
	return nil
}

// visitInteger implements Visitor.
func (i *Interpreter) VisitInteger(t exp.RinInteger) any {
	return t.Value
}

// visitPrint implements Visitor.
func (i *Interpreter) VisitPrint(t exp.RinPrint) any {
	value := i.Interpret(t.Value)
	switch v := value.(type) {
	default:
		fmt.Printf("unexpected type %T\n", v)
	case bool:
		fmt.Printf("%t\n", v)
	case int:
		fmt.Printf("%d\n", v)
	case string:
		fmt.Println(value)
	}

	return nil
}

// visitString implements Visitor.
func (i *Interpreter) VisitString(t exp.RinString) any {
	return t.Value
}
