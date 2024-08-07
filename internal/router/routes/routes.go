package routes

import (
	"net/http"
	"pawpawchat/internal/router"
)

type Routes interface {
	Register(router.Router)
}

type RouteInfo struct {
	Endpoint   string
	Methods    []string
	Handler    http.Handler
	HandleFunc func(http.ResponseWriter, *http.Request)
}

func RegisterRoutes(router router.Router, routes []Routes) {
	for _, route := range routes {
		route.Register(router)
	}
}
