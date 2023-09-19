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
