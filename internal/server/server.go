package server

import (
	"net/http"
	"pawpawchat/internal/router"
)

type server struct {
	router router.Router
}

func newServer(router router.Router) *server {
	return &server{router}
}

func (s *server) listenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
