package go_web

import(
	"net/http"
)

type ControllerRegister struct {
	routers map[string]*Tree
}

func NewControllerRegister() *ControllerRegister {
	return &ControllerRegister() *ControllerRegister {
		routers: make(map[string]*Tree),
	}
}

func (p *ControllerRegister) ServeHttp(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
} 
