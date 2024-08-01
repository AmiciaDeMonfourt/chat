package server

import (
	"pawpawchat/internal/router"
)

type server struct {
	router *router.Router
}

func newServer() *server {
	return &server{
		router: router.New(),
	}
}

func (s *server) ListenAndServe() error {
	return nil
	// http.ServeHTTP("", s.router)
}
