package go_web

import(
	"net/http"
	"fmt"
)

var (
	WebApp *App
)

func init() {
	WebApp = NewApp()
}

type App struct {
	Handlers *ControllerRegister
	Server *http.Server
}

func (app *App) Run() {
	app.Server.Handler = app.Handlers
}

func NewApp() *App {
	cr := NewControllerRegister()
	app := &App{Handlers: cr, Server: &http.Server{}}
	return app
}