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
	}
	return nil
}
