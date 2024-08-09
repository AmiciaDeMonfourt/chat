package server

import (
	"log"
	"pawpawchat/config"
	"pawpawchat/internal/router"
	"pawpawchat/internal/router/routes"
	"pawpawchat/internal/router/routes/graph"
	"pawpawchat/pkg/auth"

	authroutes "pawpawchat/internal/router/routes/auth"
)

func Run() {
	config, envConfig := config.LoadConfiguration("config.yaml")
	router := newConfiguratedRouter(envConfig)

	runHTPPServer(router, config.App.Addr)
}

func runHTPPServer(router router.Router, addr string) {
	newServer(router).listenAndServe(addr)
}

func newConfiguratedRouter(envConfig config.EnvConfigurationProvider) router.Router {
	// prrepo := factory.NewPostgresRepositoryFactory().OpenProfile(envConfig.ProfileEnvCfg().DBURL, "orm")
	authClient, err := auth.NewClient(envConfig.AuthEnvCfg().ExtAddr)
	if err != nil {
		log.Fatal(err)
	}

	routesmap := []routes.Routes{
		graph.NewRoutes(nil),
		authroutes.NewRoutes(authClient),
	}

	router := router.New()
	routes.RegisterRoutes(router, routesmap)
	return router
}
