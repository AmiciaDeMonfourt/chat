package routes

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func NewRoute(path, method string, handler func(http.ResponseWriter, *http.Request)) Route {
	return Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}
