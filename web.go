package go_web

const (
	VERSION = "0.0.1"
)

type M map[string]interface{}

type hookfunc func() error

var (
	hooks = make([]hookfunc, 0)
)

func AddAPPStartHook(hf ...hookfunc) {
	hooks = append(hooks, hf...)
}

func Run(params ...string) {
	initBeforeHTTPRun()
	
	// todo
	if len(params) > 0 && params[0] != "" {
	
	}
	
	WebApp.Run()
}

func initBeforeHTTPRun() {
	// init hooks
	AddAPPStartHook(
		registerMime,
		registerDefaultErrorHandler,
	)
	
	for _, hk := range hooks {
		if err := hk(); err != nil {
			panic(err)
		}
	}
}
