package router

import (
	"log/slog"
	"net/http"
	"pawpawchat/internal/router/middleware"

	"github.com/gorilla/mux"
)

// Router struct holds the gorilla mux router and the routes struct
type MuxRouter struct {
	router      *mux.Router
	middlewares []middleware.Middleware
}

func New() Router {
	return &MuxRouter{router: mux.NewRouter()}
}

// ServeHTTP delegating to the mux router handles HTTP requests
func (r *MuxRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.router

	for idx := range r.middlewares {
		handler = r.middlewares[idx](handler)
	}

	handler.ServeHTTP(w, req)
}

func (r *MuxRouter) Handle(path string, methods []string, handler http.Handler) {
	slog.Debug("rigester handler:", "path", path, "handler", handler)
	r.router.Handle(path, handler) //.Methods(methods...)
}

func (r *MuxRouter) HandleFunc(path string, methods []string, handleFunc func(w http.ResponseWriter, r *http.Request)) {
	slog.Debug("rigester handle func:", "path", path, "func", handleFunc)
	r.router.HandleFunc(path, handleFunc) //.Methods(methods...)
}

func (r *MuxRouter) Use(middlware ...middleware.Middleware) {
	r.middlewares = append(r.middlewares, middlware...)
}
