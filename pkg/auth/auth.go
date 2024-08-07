package auth

import (
	"log"
	"log/slog"
	"net"
	"pawpawchat/config"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/pkg/auth/authdb"
	"pawpawchat/pkg/auth/server"
	"pawpawchat/pkg/profile"

	"google.golang.org/grpc"
)

// Start ...
func Start() {
	config, envConfig := config.GetConfiguration("config.yaml")
	db := openAuthDB(envConfig.AuthEnvCfg().DBURL)

	profileClient := newProfileClient(envConfig.ProfileEnvCfg().ExtAddr)
	runGRPCServer(config, db, profileClient)
}

// openAuthDB ...
func openAuthDB(dsn string) authdb.Database {
	database, err := authdb.OpenPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}
	slog.Debug("connection with ppc_authdb", "addr", dsn)
	return database
}

// newProfileClient ...
func newProfileClient(profileGRPSServerAddr string) profilepb.ProfileClient {
	client, err := profile.NewClient(profileGRPSServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	slog.Debug("initialize profile client", "addr", profileGRPSServerAddr)
	return client
}

// runGRPCServer ...
func runGRPCServer(cfg *config.Config, db authdb.Database, pc profilepb.ProfileClient) {
	gRPSServer := grpc.NewServer()
	authpb.RegisterAuthServer(gRPSServer, server.New(db, pc))

	listener, err := net.Listen("tcp", cfg.Auth.Addr)
	if err != nil {
		log.Fatal(err)
	}

	slog.Debug("grpc server is running", "addr", cfg.Auth.Addr)
	if err := gRPSServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
