package server

import (
	"flag"
	"log"
	"log/slog"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/config"
	"pawpawchat/internal/router"
	"pawpawchat/internal/router/middleware"
	"pawpawchat/internal/router/routes"
	profile "pawpawchat/pkg/profile/database"

	"pawpawchat/internal/router/routes/auth"
	"pawpawchat/internal/router/routes/graph"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Start() {
	env := flag.String("env", "dev", "environment [dev/test]")
	flag.Parse()
	cfg := config.New(*env)

	httpRouter := router.New()

	rmap := []routes.Routes{
		graph.NewRoutes(profile.NewPostgresDB(prdbconn(cfg))),
		auth.NewRoutes(authclient(cfg)),
	}

	httpRouter.Use(middleware.CORS, middleware.Log)
	routes.RegisterRoutes(httpRouter, rmap)

	slog.Debug("server is running", "address", cfg.Addr)
	if err := newServer(httpRouter).listenAndServe(cfg.Addr); err != nil {
		log.Fatal(err)
	}
}

func prdbconn(cfg *config.AppConfig) *gorm.DB {
	var level logger.LogLevel
	if cfg.LogLevel == "info" {
		level = logger.Info
	} else {
		level = logger.Error
	}

	prdb, err := gorm.Open(postgres.Open(cfg.ProfileDBURL), &gorm.Config{Logger: logger.Default.LogMode(level)})
	if err != nil {
		log.Fatal("profile database:", err)
	}

	return prdb
}

func authclient(cfg *config.AppConfig) authpb.AuthClient {
	conn, err := grpc.NewClient(cfg.AuthAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	slog.Debug("new auth client in app:", "addr", cfg.AuthAddr)
	return authpb.NewAuthClient(conn)
}
