package profile

import (
	"fmt"
	"pawpawchat/generated/proto/profilepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(profileGRPSServerAddr string) (profilepb.ProfileClient, error) {
	conn, err := grpc.NewClient(profileGRPSServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to profile grpc server: %v", err)
	}

	return profilepb.NewProfileClient(conn), nil
}
