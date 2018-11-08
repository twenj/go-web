package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
type AppRouter = mux.Router

var router *mux.Router

func init() {
	router = mux.NewRouter().StrictSlash(true)
}

func AddRoutes(addRoutes Routes) {
	for _, route := range addRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

func GetRouter() *AppRouter {
	return router
}
