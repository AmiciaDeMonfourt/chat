package server

import (
	"context"
	"log/slog"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/pkg/auth/controller"
	"pawpawchat/pkg/auth/service"
	"pawpawchat/pkg/auth/workflow"
	"pawpawchat/pkg/auth/workflow/dbcontext"
)

type authServiceGRPCServer struct {
	userCredentialsController controller.UserController
	authpb.UnimplementedAuthServer
}

func MustNewAuthServiceGRPSServer(authdb any) *authServiceGRPCServer {
	dbctx := dbcontext.MustNew(authdb)
	unit := workflow.NewUnitOfWorkflow(dbctx)
	service := service.NewUserService(unit)
	controller := controller.NewUserController(service)
	return &authServiceGRPCServer{userCredentialsController: controller}
}

func (s *authServiceGRPCServer) SignUp(ctx context.Context, req *authpb.SignUpRequest) (*authpb.SignUpResponse, error) {
	slog.Info("new sign up request:", "req", req)
	return s.userCredentialsController.SignUp(ctx, req)
}

func (s *authServiceGRPCServer) SignIn(ctx context.Context, req *authpb.SignInRequest) (*authpb.SignInResponse, error) {
	return s.userCredentialsController.SignIn(ctx, req)
}
