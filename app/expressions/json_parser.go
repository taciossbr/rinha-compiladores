package expressions

func ParseJsonToNodes(expression map[string]any) Term {
	kind := ParseTermKind(expression["kind"].(string))
	meta := TermMeta{
		Kind:     kind,
		Location: makeLocation(expression["location"].(map[string]any)),
	}

	switch kind {
	case StringTermKind:
		return &RinString{
			Meta:  meta,
			Value: expression["value"].(string),
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
	}
	return nil
}
