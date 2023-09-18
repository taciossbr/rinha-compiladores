package main

import (
	"fmt"
	"os"
)

// TODO separar em pacotes

func main() {
	// TODO ler dos parametros
	codeFile := readFile(os.Args[1])

	fmt.Println("Executing:", codeFile.Name)
	fmt.Println()
	fmt.Println()
	expression := parseJsonToNodes(codeFile.Expression)

	interpreter := Interpreter{}
	interpreter.Interpret(expression)
}
