package expressions

type Term interface {
	Accept(v Visitor) any
	GetKind() TermKind
	GetLocation() Location
}

type Location struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Filename string `json:"filename"`
}

type Parameter struct {
	Text string
	Location
}

type TermMeta struct {
	Kind     TermKind `json:"kind"`
	Location Location `json:"location"`
}

func makeLocation(data map[string]any) Location {
	location := Location{
		Start:    int(data["start"].(float64)),
		End:      int(data["end"].(float64)),
		Filename: data["filename"].(string),
	}
	return location
}
func makeParameter(data map[string]any) Parameter {
	parameter := Parameter{
		Text:     data["text"].(string),
		Location: makeLocation(data["location"].(map[string]any)),
	}
	return parameter
}
