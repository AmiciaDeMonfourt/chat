package client

import (
	"context"
	"log/slog"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/dto"
	"pawpawchat/internal/model/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AuthServiceClient ...
type AuthServiceClient interface {
	// SignUp ...
	SignUp(context.Context, *domain.User) (*domain.User, error)
	// SignIn ...
	SignIn(context.Context, *domain.UserCredentials) (*domain.User, error)
}

type authServiceClient struct {
	connection authpb.AuthClient
}

func NewAuthServiceClient(addr string) (AuthServiceClient, error) {
	connection, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &authServiceClient{authpb.NewAuthClient(connection)}, nil
}

func (c *authServiceClient) SignUp(ctx context.Context, user *domain.User) (*domain.User, error) {
	newuser := &authpb.NewUser{}

	slog.Info("auth service client:", "user", user)

	if err := dto.EncodeUser(user, newuser); err != nil {
		return nil, err
	}

	response, err := c.connection.SignUp(ctx, &authpb.SignUpRequest{User: newuser})
	if err != nil {
		return nil, err
	}

	user, err = dto.ExtractUser(response)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *authServiceClient) SignIn(ctx context.Context, credentials *domain.UserCredentials) (*domain.User, error) {
	return nil, nil
}
