package server

import (
	"context"
	"log/slog"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/auth/authdb"
	"pawpawchat/utils/encrypt"
	"pawpawchat/utils/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Server struct {
	authpb.UnimplementedAuthServer
	profileClient profilepb.ProfileClient
	database      authdb.Database
}

func New(db authdb.Database, pc profilepb.ProfileClient) *Server {
	return &Server{database: db, profileClient: pc}
}

func NewClient(AuthGRPCServerAddr string) (authpb.AuthClient, error) {
	connection, err := grpc.NewClient(AuthGRPCServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return authpb.NewAuthClient(connection), nil
}

// SignUp receives credentials for the new user and returns the created user
func (s *Server) SignUp(ctx context.Context, req *authpb.SignUpRequest) (*authpb.SignUpResponse, error) {
	wd := "auth -> s.SignUp()"
	hashPass, err := encrypt.EncryptString(req.GetCredentials().GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s -> encryptString() -> %v", wd, err)
	}

	credentials := &domain.UserCredentials{
		Email:    req.GetCredentials().GetEmail(),
		HashPass: hashPass,
	}

	tx := s.database.Begin()
	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "%s -> database.Begin() -> %v", wd, tx.Error)
	}

	// insert credentials in database
	if err := s.database.InsertUserCredentials(ctx, tx, credentials); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.InvalidArgument, "%s -> database.InsertUserCredentials() -> %v", wd, err)
	}

	createRequest := &profilepb.CreateRequest{
		Userbio: &profilepb.UserBiography{
			Firstname:  req.GetUserbio().GetFirstname(),
			Secondname: req.GetUserbio().GetSecondname(),
		},
	}

	createResponse, err := s.profileClient.Create(ctx, createRequest)
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.InvalidArgument, "%s -> profileCLient.Create() -> %v", wd, err)
	}

	// Commit the transaction if everything succeeds
	if err := tx.Commit().Error; err != nil {
		return nil, status.Errorf(codes.Internal, "%s -> tx.Commit() -> %v", wd, err)
	}

	// insert user personal info in page service
	user := &authpb.User{
		Id: createResponse.GetUser().GetUserid(),
		Userinfo: &authpb.Biography{
			Firstname:  createResponse.GetUser().GetUserbio().GetFirstname(),
			Secondname: createResponse.GetUser().GetUserbio().GetSecondname(),
		},
		Credentials: &authpb.Credentials{
			Email: credentials.Email,
		},
	}

	slog.Info("SignUp:", "user", user)

	// generate jwt token
	tokenStr, err := jwt.GenerateToken(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s -> jwt.GenerateToken() -> %v", wd, err)
	}

	return &authpb.SignUpResponse{User: user, TokenString: tokenStr}, nil
}

// SignIn receives credentials for the authorization and return the authorized user
func (s *Server) SignIn(ctx context.Context, req *authpb.SignInRequest) (*authpb.SignInResponse, error) {

	return nil, nil
}
