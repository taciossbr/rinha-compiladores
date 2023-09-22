package interpreter

type Environment struct {
	environmentMap map[string]any
	closure        *Environment
}

func (e *Environment) Get(name string) any {
	value := e.environmentMap[name]
	if value == nil && e.closure != nil {
		return e.closure.Get(name)
	}
	return value
}

func (e *Environment) Set(name string, value any) {
	e.environmentMap[name] = value
}

func makeEnvironment(enc *Environment) *Environment {
	env := Environment{
		environmentMap: make(map[string]any),
		closure:        enc,
	}
	return &env
}
