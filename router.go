package go_web

import (
	"fmt"
	"net/http"
)

func RouterBind(controllerPath string) http.HandlerFunc {
	handle := func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello")
	}
	return handle
}
