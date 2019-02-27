package go_web

import (
	"go-web/context"
)

type Controller struct {
	Ctx *context.Context
	Data map[interface{}]interface{}
	
	// route controller info
	controllerName string	
	actionName string
	methodMapping map[string]func()
	
	// template data
	TplName string
	ViewPath string
}

type ControllerInterface interface {
	Get()
}
