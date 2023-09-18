package interpreter

import (
	"fmt"
	"rinha/expressions"
)

type Interpreter struct{}

func (i *Interpreter) Interpret(term expressions.Term) any {
	return term.Accept(i)
}

// visitInteger implements Visitor.
func (i *Interpreter) VisitInteger(t expressions.RinInteger) any {
	return t.Value
}

// visitPrint implements Visitor.
func (i *Interpreter) VisitPrint(t expressions.RinPrint) any {
	value := i.Interpret(t.Value)
	switch v := value.(type) {
	default:
		fmt.Printf("unexpected type %T\n", v)
	case int:
		fmt.Printf("%d\n", v)
	case string:
		fmt.Println(value)
	}

	return nil
}

// visitString implements Visitor.
func (i *Interpreter) VisitString(t expressions.RinString) any {
	return t.Value
}
