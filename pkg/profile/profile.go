package profile

import (
	"log"
	"log/slog"
	"net"
	"pawpawchat/config"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/pkg/profile/profiledb"
	"pawpawchat/pkg/profile/server"

	"google.golang.org/grpc"
)

// Start ...
func Start() {
	config, envConfig := config.GetConfiguration("config.yaml")
	profiledb := openProfileDB(envConfig.ProfileEnvCfg().DBURL)

	runGRPCServer(config.Profile, profiledb)
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

// runGRPCServer ...
func runGRPCServer(cfg config.ServiceConfig, db profiledb.Database) {
	gRPSServer := grpc.NewServer()
	profilepb.RegisterProfileServer(gRPSServer, server.New(db))

	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatal(err)
	}

	slog.Debug("grpc server is running", "addr", cfg.Addr)
	if err := gRPSServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
