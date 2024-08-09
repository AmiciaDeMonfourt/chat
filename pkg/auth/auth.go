package auth

import (
	"log"
	"log/slog"
	"net"
	"pawpawchat/config"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/pkg/auth/client"
	"pawpawchat/pkg/auth/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunServer() {
	cfg, envcfg := config.LoadConfiguration("config.yaml")

	db, err := gorm.Open(postgres.Open(envcfg.AuthEnvCfg().DBURL))
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	authServer := server.MustNewAuthServiceGRPSServer(db)
	authpb.RegisterAuthServer(gRPCServer, authServer)

	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", cfg.Auth.Addr)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("auth service server listening", "addr", cfg.Auth.Addr)
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func NewClient(addr string) (client.AuthServiceClient, error) {
	return client.NewAuthServiceClient(addr)
}
