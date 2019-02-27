package go_web

import (
	"mime"
	"net/http"
)

func registerMime() error {
	for k, v := range mimemaps {
		mime.AddExtensionType(k, v)
	}
	return nil
}

func registerDefaultErrorHandler() error {
	m := map[string]func(http.ResponseWriter, *http.Request) {
		"401": unauthorized,
	}
	for e, h := range m {
		if _, ok := ErrorMaps[e]; !ok {
			ErrorHandler(e, h);
		}
	}
	return nil
}
