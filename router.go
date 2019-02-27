package go_web

import (
	"net/http"
)

type ControllerRegister struct {
	routers string
}

func NewControllerRegister() *ControllerRegister {
	return &ControllerRegister {
		routers: "",
	}
}

func (p *ControllerRegister) ServeHttp(rw http.ResponseWriter, r *http.Request) {

}

func (p *ControllerRegister) Add(pattern string, c ControllerInterface) {
	p.addWithMethodParams(pattern, c)
}

func (p *ControllerRegister) addWithMethodParams(patter string, c ControllerInterface) {
	
}
