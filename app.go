package go_web

import(
	"net/http"
	"fmt"
	"os"
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
	fmt.Println(app.Handlers)
	os.Exit(1)
	// app.Server.Handler = app.Handlers
}

func NewApp() *App {
	cr := NewControllerRegister()
	app := &App{Handlers: cr, Server: &http.Server{}}
	return app
}
