package routes

import (
	"net/http"
)

type Routes struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func New() *Routes {
	return &Routes{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (r *Routes) GetRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return r.routes
}

func (r *Routes) Configure() {
	r.routes["/"] = TestFoo
}

func TestFoo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("xui"))
}
