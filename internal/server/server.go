package server

import (
	"net/http"
	"pawpawchat/internal/router"
)

type server struct {
	router *router.Router
}

func newServer() *server {
	router := router.New()
	router.Configure()

	return &server{
		router: router,
	}
}

func (s *server) listenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
