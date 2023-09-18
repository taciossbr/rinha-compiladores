package main

import "fmt"

type Interpreter struct{}

func (i *Interpreter) Interpret(term Term) any {
	return term.Accept(i)
}

// visitInteger implements Visitor.
func (i *Interpreter) visitInteger(t RinInteger) any {
	return t.Value
}

// visitPrint implements Visitor.
func (i *Interpreter) visitPrint(t RinPrint) any {
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
func (i *Interpreter) visitString(t RinString) any {
	return t.Value
}
