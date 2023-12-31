package main

import (
	"rinha/expressions"
	inter "rinha/interpreter"
)

func main() {
	codeFile := expressions.ReadFile("/var/rinha/source.rinha.json")

	expression := expressions.ParseJsonToNodes(codeFile.Expression)

	interpreter := inter.MakeInterpreter()
	interpreter.Interpret(expression)
}
