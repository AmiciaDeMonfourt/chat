package auth

import (
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/router"
	"pawpawchat/internal/router/routes"
)

type AuthRoutes struct {
	authClient authpb.AuthClient
	routes     []routes.RouteInfo
}

func NewRoutes(authclient authpb.AuthClient) routes.Routes {
	authRoutes := &AuthRoutes{authClient: authclient, routes: make([]routes.RouteInfo, 0)}
	authRoutes.configure()
	return authRoutes
}

func (r *AuthRoutes) Register(router router.Router) {
	for _, route := range r.routes {
		router.HandleFunc(route.Endpoint, route.Methods, route.HandleFunc)
	}
}

func (r *AuthRoutes) configure() {
	r.routes = append(r.routes,
		routes.RouteInfo{
			Endpoint:   "/signup",
			Methods:    []string{"POST"},
			HandleFunc: r.SignUp,
		},
		routes.RouteInfo{
			Endpoint:   "/signin",
			Methods:    []string{"POST"},
			HandleFunc: r.SignIn,
		},
	)
}
