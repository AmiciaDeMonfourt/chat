package authroutes

import (
	"log"
	"os"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/producer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthRoutes struct {
	producer   *producer.Producer
	authClient authpb.AuthClient
}

func NewAuthRoutes() *AuthRoutes {
	producer := producer.New("test-topic")
	go producer.StartProduce()

	return &AuthRoutes{
		producer:   producer,
		authClient: authpb.NewAuthClient(newAuthConn()),
	}
}

func newAuthConn() grpc.ClientConnInterface {
	addr := os.Getenv("APP_AUTH_ADDR")
	if addr == "" {
		log.Fatal("APP_AUTH_ADDR is missing")
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
