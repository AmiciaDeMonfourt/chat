package routes

import "net/http"

// request handler func
type HandleFunc func(http.ResponseWriter, *http.Request)

type Routes struct {
	routes map[string]HandleFunc
}

func New() *Routes {
	return &Routes{
		routes: make(map[string]HandleFunc),
	}
}

func (r *Routes) GetRoutes() map[string]HandleFunc {
	return r.routes
}

func (r *Routes) Configure() {

}
