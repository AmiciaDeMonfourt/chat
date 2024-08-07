package router

import (
	"net/http"
	"pawpawchat/internal/router/middleware"
)

type Router interface {
	// ServeHTTP ...
	ServeHTTP(http.ResponseWriter, *http.Request)

	// HandleFunc ...
	HandleFunc(endpoint string, methods []string, handlerFunc func(http.ResponseWriter, *http.Request))

	// Handle..
	Handle(endpoint string, methods []string, handler http.Handler)

	// Use ...
	Use(middlewares ...middleware.Middleware)
}
