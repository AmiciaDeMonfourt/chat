package auth

import (
	"log"
	"log/slog"
	"net"
	"os"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/pkg/auth/server"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(wd + "/.env"); err != nil {
		log.Fatal(err)
	}
}

func Start() {
	addr := os.Getenv("AUTH_ADDR")
	if addr == "" {
		log.Fatal("AUTH_ADDR is missing")
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to create a listener: " + err.Error())
	}

	srv := grpc.NewServer()
	authpb.RegisterAuthServer(srv, server.New())

	slog.Info("auth service is started")

	srv.Serve(listener)
}
