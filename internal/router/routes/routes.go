package routes

import "pawpawchat/internal/router/routes/auth"

type Routes struct {
	routes     []Route
	authRoutes auth.AuthRoutes
}

func New() *Routes {
	return &Routes{
		routes:     make([]Route, 0),
		authRoutes: auth.NewAuthRoutes(),
	}
}

func (r *Routes) GetRoutes() []Route {
	return r.routes
}

func (r *Routes) Configure() {
	r.routes = append(r.routes, NewRoute("/signup", "POST", r.authRoutes.SignUp))
	r.routes = append(r.routes, NewRoute("/signin", "POST", r.authRoutes.SignUp))
}
