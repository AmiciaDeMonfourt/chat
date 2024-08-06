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

	if os.Getenv("AUTH_ADDR") == "" {
		log.Fatal("auth -> AUTH_ADDR is missing")

	} else if os.Getenv("AUTH_DB_URL") == "" {
		log.Fatal("auth -> AUTH_DB_URL is missing")

	} else if os.Getenv("PROFILE_EXTERNAL_ADDR") == "" {
		log.Fatal("auth -> PROFILE_EXTERNAL_ADDR is missing")
	}

}

func Start() {
	srv := grpc.NewServer()
	authpb.RegisterAuthServer(srv, server.New())

	listener, err := net.Listen("tcp", os.Getenv("AUTH_ADDR"))
	if err != nil {
		log.Fatalf("auth -> failed to create a listener: " + err.Error())
	}

	slog.Info("auth -> server is running", "addr", os.Getenv("AUTH_ADDR"))

	if err := srv.Serve(listener); err != nil {
		log.Fatal("auth -> server error:", err.Error())
	}
}
