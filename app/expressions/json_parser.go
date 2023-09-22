package expressions

func ParseJsonToNodes(expression map[string]any) Term {
	kind := ParseTermKind(expression["kind"].(string))
	meta := TermMeta{
		Kind:     kind,
		Location: makeLocation(expression["location"].(map[string]any)),
	}

	switch kind {
	case BoolTermKind:
		return &RinBool{
			Meta:  meta,
			Value: expression["value"].(bool),
		}
	case BinaryTermKind:
		return &RinBinary{
			Meta:      meta,
			Left:      ParseJsonToNodes(expression["lhs"].(map[string]any)),
			Operation: ParseOp(expression["op"].(string)),
			Right:     ParseJsonToNodes(expression["rhs"].(map[string]any)),
		}
	case IntTermKind:
		return &RinInteger{
			Meta:  meta,
			Value: int(expression["value"].(float64)),
		}
	case PrintTermKind:
		return &RinPrint{
			Meta:  meta,
			Value: ParseJsonToNodes(expression["value"].(map[string]any)),
		}
	case StringTermKind:
		return &RinString{
			Meta:  meta,
			Value: expression["value"].(string),
		}
	case LetTermKind:
		return &RinLet{
			Meta:  meta,
			Name:  makeParameter(expression["name"].(map[string]any)),
			Value: ParseJsonToNodes(expression["value"].(map[string]any)),
			Next:  ParseJsonToNodes(expression["next"].(map[string]any)),
		}
	case VarTermKind:
		return &RinVar{
			Meta: meta,
			Text: expression["text"].(string),
		}
	case FunctionTermKind:
		var parameters []Parameter
		for _, param := range expression["parameters"].([]interface{}) {
			parameters = append(parameters, makeParameter(param.(map[string]any)))
		}
		return &RinFunction{
			Meta:       meta,
			Parameters: parameters,
			Value:      ParseJsonToNodes(expression["value"].(map[string]any)),
		}
	case CallTermKind:
		var arguments []Term
		for _, arg := range expression["arguments"].([]interface{}) {
			arguments = append(arguments, ParseJsonToNodes(arg.(map[string]any)))
		}
		return &RinCall{
			Meta:      meta,
			Callee:    ParseJsonToNodes(expression["callee"].(map[string]any)),
			Arguments: arguments,
		}
	case IfTermKind:
		return &RinIf{
			Meta:      meta,
			Condition: ParseJsonToNodes(expression["condition"].(map[string]any)),
			Then:      ParseJsonToNodes(expression["then"].(map[string]any)),
			Otherwise: ParseJsonToNodes(expression["otherwise"].(map[string]any)),
		}
	case TupleTermKind:
		return &RinTuple{
			Meta:   meta,
			First:  ParseJsonToNodes(expression["first"].(map[string]any)),
			Second: ParseJsonToNodes(expression["second"].(map[string]any)),
		}
	case FirstTermKind:
		return &RinFirst{
			Meta:  meta,
			Value: ParseJsonToNodes(expression["value"].(map[string]any)),
		}
	case SecondTermKind:
		return &RinSecond{
			Meta:  meta,
			Value: ParseJsonToNodes(expression["value"].(map[string]any)),
		}
	}
	return nil
}
