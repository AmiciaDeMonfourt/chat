package auth

import (
	"pawpawchat/internal/router"
	"pawpawchat/internal/router/routes"
	"pawpawchat/pkg/auth/client"
)

type AuthRoutes struct {
	authClient client.AuthServiceClient
	routes     []routes.RouteInfo
}

func NewRoutes(authclient client.AuthServiceClient) routes.Routes {
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
