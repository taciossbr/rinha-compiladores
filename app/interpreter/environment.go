package interpreter

type Environment struct {
	environmentMap map[string]any
}

func (e *Environment) Get(name string) any {
	return e.environmentMap[name]
}

func (e *Environment) Set(name string, value any) {
	e.environmentMap[name] = value
}

func makeEnvironment() Environment {
	env := Environment{
		environmentMap: make(map[string]any),
	}
	return env
}
