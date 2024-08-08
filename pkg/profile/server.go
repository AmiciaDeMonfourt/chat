package profile

import (
	"context"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/pkg/profile/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// profileServer ...
type profileServer struct {
	profileHandler *handler.Profile
	profilepb.UnimplementedProfileServer
}

func newServer(profileHandler *handler.Profile) *profileServer {
	return &profileServer{profileHandler: profileHandler}
}

// Create ...
func (s *profileServer) Create(ctx context.Context, req *profilepb.CreateRequest) (*profilepb.CreateResponse, error) {
	return s.profileHandler.Create(ctx, req)
}

// NewClient ...
func NewClient(profileGRPSServerAddr string) (profilepb.ProfileClient, error) {
	conn, err := grpc.NewClient(profileGRPSServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return profilepb.NewProfileClient(conn), nil
}
