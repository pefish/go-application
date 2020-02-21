package go_application

type ApplicationClass struct {
	Debug bool
	Env   string
}

func (app *ApplicationClass) SetEnv(env string) {
	app.Env = env
	app.Debug = IsDebug(env)
}

var Application = ApplicationClass{
	Env: `dev`,
	Debug: IsDebug(`dev`),
}

var OnTerminated = make(chan interface{})

func IsDebug(env string) bool {
	return env != `prod`
}
