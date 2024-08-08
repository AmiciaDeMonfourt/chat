package server

import (
	"log"
	"pawpawchat/config"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/router"
	"pawpawchat/internal/router/routes"
	"pawpawchat/internal/router/routes/graph"
	"pawpawchat/pkg/profile/repository/factory"

	authroutes "pawpawchat/internal/router/routes/auth"
	authsrv "pawpawchat/pkg/auth/server"
)

func Start() {
	config, envConfig := config.GetConfiguration("config.yaml")
	router := newConfiguratedRouter(envConfig)

	runHTPPServer(router, config.App.Addr)
}

func runHTPPServer(router router.Router, addr string) {
	newServer(router).listenAndServe(addr)
}

func newConfiguratedRouter(envConfig config.EnvConfigurationProvider) router.Router {
	prrepo := factory.NewPostgresRepositoryFactory().OpenProfile(envConfig.ProfileEnvCfg().DBURL, "orm")
	authClient := newAuthClient(envConfig.AuthEnvCfg().ExtAddr)

	routesmap := []routes.Routes{
		graph.NewRoutes(prrepo),
		authroutes.NewRoutes(authClient),
	}

	router := router.New()
	routes.RegisterRoutes(router, routesmap)
	return router
}

func newAuthClient(authGRPCServerAddr string) authpb.AuthClient {
	client, err := authsrv.NewClient(authGRPCServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
