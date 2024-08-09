package controller

import (
	"context"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/dto"
	"pawpawchat/pkg/auth/service"
)

// User ...
type UserController interface {
	// SignUp ...
	SignUp(context.Context, *authpb.SignUpRequest) (*authpb.SignUpResponse, error)
	// SignIn
	SignIn(context.Context, *authpb.SignInRequest) (*authpb.SignInResponse, error)
}

func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

// userCredentialsController ...
type userController struct {
	service service.UserService
}

func (uc *userController) SignUp(ctx context.Context, req *authpb.SignUpRequest) (*authpb.SignUpResponse, error) {
	user, err := dto.ExtractUser(req)
	if err != nil {
		return nil, err
	}

	if err = uc.service.SignUp(ctx, &user.Credentials); err != nil {
		return nil, err
	}

	userResponse := &authpb.User{}
	dto.EncodeUser(user, userResponse)

	return &authpb.SignUpResponse{User: userResponse}, nil
}

func (uc *userController) SignIn(context.Context, *authpb.SignInRequest) (*authpb.SignInResponse, error) {
	return nil, nil
}
