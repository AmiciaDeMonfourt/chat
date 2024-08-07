package auth

import (
	"log"
	"log/slog"
	"net"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/auth/config"
	"pawpawchat/pkg/auth/database"
	"pawpawchat/pkg/auth/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	cfg := config.New()

	srv := grpc.NewServer()
	authpb.RegisterAuthServer(srv, server.New(authdb(cfg), profileclient(cfg)))

	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatalf("auth -> failed to create a listener: " + err.Error())
	}

	slog.Info("auth -> server is running", "addr", cfg.Addr)

	if err := srv.Serve(listener); err != nil {
		log.Fatal("auth -> server error:", err.Error())
	}
}

func authdb(cfg *config.AuthConfig) database.Database {
	db, err := gorm.Open(postgres.Open(cfg.DBURL))
	if err != nil {
		log.Fatal("failed to connect to db:", err.Error())
	}

	if err := db.AutoMigrate(domain.UserCredentials{}); err != nil {
		log.Fatal("failed to run migrations:", err.Error())
	}

	return database.NewPosgresDB(db)
}

func profileclient(cfg *config.AuthConfig) profilepb.ProfileClient {
	conn, err := grpc.NewClient(cfg.ProfileAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("auth -> failed to connection to profile grpc server:", err.Error())
	}

	slog.Info("profile client ina auth service", "addr", cfg.ProfileAddr)
	return profilepb.NewProfileClient(conn)
}
