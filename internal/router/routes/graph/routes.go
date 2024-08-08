package graph

import (
	"net/http"
	"pawpawchat/api/graph/resolvers"
	"pawpawchat/generated/graphgen"
	"pawpawchat/internal/router"
	"pawpawchat/internal/router/routes"
	"pawpawchat/pkg/profile/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GraphRoutes struct {
	routes []routes.RouteInfo
	prrepo repository.Profile
}

func NewRoutes(profileRepository repository.Profile) routes.Routes {
	gr := &GraphRoutes{
		routes: make([]routes.RouteInfo, 0),
		prrepo: profileRepository,
	}
	gr.configure()
	return gr
}

func (r *GraphRoutes) Register(router router.Router) {
	for _, route := range r.routes {
		router.Handle(route.Endpoint, route.Methods, route.Handler)
	}
}

func (r *GraphRoutes) GraphHandler() http.Handler {
	return handler.NewDefaultServer(graphgen.NewExecutableSchema(graphgen.Config{Resolvers: &resolvers.Resolver{ProfileRepository: r.prrepo}}))
}

func (r *GraphRoutes) PlaygroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/query")
}

func (r *GraphRoutes) configure() {
	r.routes = append(r.routes,
		routes.RouteInfo{
			Endpoint: "/playground",
			Methods:  []string{"GET", "POST"},
			Handler:  r.PlaygroundHandler(),
		},
		routes.RouteInfo{
			Endpoint: "/query",
			Methods:  []string{"GET", "POST"},
			Handler:  r.GraphHandler(),
		},
	)
}
