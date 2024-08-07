package server

import (
	"log"
	"log/slog"
	"pawpawchat/config"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/router"
	"pawpawchat/internal/router/routes"
	"pawpawchat/internal/router/routes/graph"
	"pawpawchat/pkg/profile/profiledb"

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
	profiledb := openProfileDB(envConfig.ProfileEnvCfg().DBURL)
	authClient := newAuthClient(envConfig.AuthEnvCfg().ExtAddr)

	routesmap := []routes.Routes{
		graph.NewRoutes(profiledb),
		authroutes.NewRoutes(authClient),
	}

	router := router.New()
	routes.RegisterRoutes(router, routesmap)
	return router
}

// openAuthDB ...
func openProfileDB(dsn string) profiledb.Database {
	database, err := profiledb.OpenPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}
	slog.Debug("connection with ppc_profile", "dsn", dsn)
	return database
}

func newAuthClient(authGRPCServerAddr string) authpb.AuthClient {
	client, err := authsrv.NewClient(authGRPCServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
