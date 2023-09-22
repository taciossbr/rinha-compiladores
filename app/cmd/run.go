package main

import (
	"fmt"
	"os"
	"rinha/expressions"
	inter "rinha/interpreter"
)

func main() {
	codeFile := expressions.ReadFile(os.Args[1])

	fmt.Println("Executing:", codeFile.Name)
	fmt.Println()
	fmt.Println()
	expression := expressions.ParseJsonToNodes(codeFile.Expression)

	interpreter := inter.MakeInterpreter()
	interpreter.Interpret(expression)
}
