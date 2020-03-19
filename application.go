package go_application

import "context"

type ApplicationClass struct {
	Debug  bool
	Env    string
	ctx    context.Context
	cancel context.CancelFunc
}

func (app *ApplicationClass) SetEnv(env string) {
	app.Env = env
	app.Debug = isDebug(env)
}

// 退出应用
func (app *ApplicationClass) Exit() {
	app.cancel()
}

// 应用结束通道
func (app *ApplicationClass) OnFinished() <-chan struct{} {
	return app.ctx.Done()
}

var ctx, cancel = context.WithCancel(context.Background())
var Application = ApplicationClass{
	Env:    `dev`,
	Debug:  isDebug(`dev`),
	ctx:    ctx,
	cancel: cancel,
}

func isDebug(env string) bool {
	return env != `prod`
}
