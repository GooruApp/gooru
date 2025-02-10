package env

type env struct {
	appEnv    string
	dbConnStr string
	port      int
}

func Get() (*env, error) {
	env := &env{
		appEnv:    initAppEnv(),
		dbConnStr: initDBConnStr(),
		port:      initPort(),
	}

	return env, nil
}

func (e *env) AppEnv() string {
	return e.appEnv
}

func (e *env) DBConnStr() string {
	return e.dbConnStr
}

func (e *env) Port() int {
	return e.port
}
