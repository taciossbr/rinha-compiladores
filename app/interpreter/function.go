package interpreter

import (
	exp "rinha/expressions"
)

type FunctionEx struct {
	declaration exp.RinFunction
	// closure Environment
}

func makeFunction(declaration exp.RinFunction) FunctionEx {
	return FunctionEx{
		declaration: declaration,
	}
}

func (f *FunctionEx) Call(i *Interpreter, args []interface{}) any {
	enc := makeEnvironment(i.env)
	for i, param := range f.declaration.Parameters {
		enc.Set(param.Text, args[i])
	}
	return i.ExecuteBlock(f.declaration.Value, enc)
}

func (f *FunctionEx) Arity() int {
	return len(f.declaration.Parameters)
}
