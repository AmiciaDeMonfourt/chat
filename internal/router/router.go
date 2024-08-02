package router

import (
	"net/http"
	"pawpawchat/internal/router/routes"

	"github.com/gorilla/mux"
)

// Router struct holds the gorilla mux router and the routes struct
type Router struct {
	router *mux.Router
	routes *routes.Routes
}

func New() *Router {
	routes := routes.New()
	routes.Configure()

	return &Router{
		router: mux.NewRouter(),
		routes: routes,
	}
}

// ServeHTTP delegating to the mux router handles HTTP requests
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

// Configure sets up the router by registering all routes and their handlers
func (r *Router) Configure() {
	// iterate over all routes and handlers
	for path, handleFunc := range r.routes.GetRoutes() {
		// register each route and its handler
		r.router.HandleFunc(path, handleFunc)
	}
}
