package main

import (
	"fmt"
	"os"
	"rinha/expressions"
	"rinha/interpreter"
)

// TODO separar em pacotes

func main() {
	// TODO ler dos parametros
	codeFile := expressions.ReadFile(os.Args[1])

	fmt.Println("Executing:", codeFile.Name)
	fmt.Println()
	fmt.Println()
	expression := expressions.ParseJsonToNodes(codeFile.Expression)

	interpreter := interpreter.Interpreter{}
	interpreter.Interpret(expression)
}
