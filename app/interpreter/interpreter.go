package interpreter

import (
	"fmt"
	"os"
	"reflect"
	exp "rinha/expressions"
	"strconv"
)

type Interpreter struct {
	env *Environment
}

func MakeInterpreter() Interpreter {
	i := Interpreter{
		env: makeEnvironment(nil),
	}

	return i
}

func (i *Interpreter) Interpret(t exp.Term) any {
	return t.Accept(i)
}

func (*Interpreter) VisitBool(t exp.RinBool) any {
	return t.Value
}

// VisitBinary implements exp.Visitor.
func (i *Interpreter) VisitBinary(t exp.RinBinary) any {
	left := i.Interpret(t.Left)
	right := i.Interpret(t.Right)

	switch t.Operation {
	case exp.AddOp:
		leftType := reflect.TypeOf(left).Kind()
		rightType := reflect.TypeOf(right).Kind()

		if leftType == reflect.String || rightType == reflect.String {
			return ensureString(left) + ensureString(right)
		}

		return left.(int) + right.(int)

	case exp.SubOp:
		return left.(int) - right.(int)
	case exp.MulOp:
		return left.(int) * right.(int)
	case exp.DivOp:
		return left.(int) / right.(int)
	case exp.RemOp:
		return left.(int) % right.(int)
	case exp.EqOp:
		leftType := reflect.TypeOf(left)
		rightType := reflect.TypeOf(right)
		if leftType != rightType {
			return false
		}
		return left == right
	case exp.NeqOp:
		leftType := reflect.TypeOf(left)
		rightType := reflect.TypeOf(right)
		if leftType != rightType {
			return true
		}
		return left != right
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
	fmt.Println(format(value))

	return value
}

func format(value any) string {
	switch v := value.(type) {
	default:
		return fmt.Sprintf("unexpected type %T\n", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case int:
		return fmt.Sprintf("%d", v)
	case string:
		return fmt.Sprint(value)
	case RinTupleDS:
		tuple := value.(RinTupleDS)
		first := format(tuple.First)
		second := format(tuple.Second)

		return fmt.Sprintf("(%s, %s)\n", first, second)
	case FunctionEx:
		return "<#closure>"
	}
}

// visitString implements Visitor.
func (i *Interpreter) VisitString(t exp.RinString) any {
	return t.Value
}

// VisitLet implements expressions.Visitor.
func (i *Interpreter) VisitLet(t exp.RinLet) any {
	value := i.Interpret(t.Value)
	i.env.Set(t.Name.Text, value)
	return i.Interpret(t.Next)
}

// VisitLet implements expressions.Visitor.
func (i *Interpreter) VisitVar(t exp.RinVar) any {
	return i.env.Get(t.Text)
}

// VisitLet implements expressions.Visitor.
func (i *Interpreter) VisitFunction(t exp.RinFunction) any {
	return makeFunction(t, i.env)
}

// VisitLet implements expressions.Visitor.
func (i *Interpreter) VisitCall(t exp.RinCall) any {
	function := i.Interpret(t.Callee).(FunctionEx)
	if len(t.Arguments) != function.Arity() {
		i.exit("Invalid number of arguments")
	}
	var argsValues []interface{}
	for _, arg := range t.Arguments {
		value := i.Interpret(arg)
		argsValues = append(argsValues, value)
	}

	return function.Call(i, argsValues)
}

// VisitLet implements expressions.Visitor.
func (i *Interpreter) VisitIf(t exp.RinIf) any {
	condition := i.Interpret(t.Condition).(bool)
	if condition {
		return i.Interpret(t.Then)
	} else {
		return i.Interpret(t.Otherwise)
	}
}

// VisitLet implements expressions.Visitor.
func (i *Interpreter) VisitTuple(t exp.RinTuple) any {
	first := i.Interpret(t.First)
	second := i.Interpret(t.Second)
	tuple := RinTupleDS{
		First:  first,
		Second: second,
	}
	return tuple
}

func (i *Interpreter) VisitFirst(t exp.RinFirst) any {
	tuple := i.Interpret(t.Value).(RinTupleDS)
	return tuple.First
}

func (i *Interpreter) VisitSecond(t exp.RinSecond) any {
	tuple := i.Interpret(t.Value).(RinTupleDS)
	return tuple.Second
}

func (i *Interpreter) ExecuteBlock(t exp.Term, enc *Environment) any {
	oldEnv := i.env
	i.env = enc
	ret := i.Interpret(t)
	i.env = oldEnv
	return ret
}

func (i *Interpreter) exit(message string) {
	fmt.Print("Error: ")
	fmt.Println(message)
	os.Exit(2)
}

func ensureString(value any) string {
	// TODO talvez estourar um erro quando nao for algo valido
	if reflect.TypeOf(value).Kind() == reflect.Int {
		return strconv.Itoa(value.(int))
	}
	return value.(string)
}
