package environment

type Environment struct {
	appEnv    string
	dbConnStr string
	port      int
}

func Get() (*Environment, error) {
	env := &Environment{
		appEnv:    initAppEnv(),
		dbConnStr: initDBConnStr(),
		port:      initPort(),
	}

	return env, nil
}

func (e *Environment) AppEnv() string {
	return e.appEnv
}

func (e *Environment) DBConnStr() string {
	return e.dbConnStr
}

func (e *Environment) Port() int {
	return e.port
}
