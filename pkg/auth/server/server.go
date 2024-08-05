package server

import (
	"context"
	"log"
	"os"
	pb "pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/producer"
	"pawpawchat/pkg/auth/database"
	"pawpawchat/pkg/auth/model"
	"pawpawchat/utils/errors"
	"pawpawchat/utils/jwt"

	"google.golang.org/grpc/codes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedAuthServer
	producer *producer.Producer
	database database.Database
}

func New() *Server {
	dbURL := os.Getenv("AUTH_DB_URL")
	if dbURL == "" {
		log.Fatal("env AUTH_DB_URL is missing")
	}

	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		log.Fatal("failed to connect to auth db: " + err.Error())
	}

	if err := db.AutoMigrate(model.UserCredentials{}); err != nil {
		log.Fatal(err)
	}

	return &Server{
		producer: nil, //producer.New("new-users"),
		database: database.NewPosgresDB(db),
	}
}

// SignUp receives credentials for the new user and returns the created user
func (s *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	cr := &model.UserCredentials{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	// insert credentials in database
	if err := s.database.InsertUserCredentials(context.TODO(), cr); err != nil {
		return nil, errors.NewGRPC(
			"failed to insert record in db: "+err.Error(), // temporary measure REFACTOR
			"auth.server.SignUp() -> s.database.InsertUserCredentials()",
			err.Error(),
			codes.InvalidArgument,
		)
	}

	// insert user personal info in page service
	user := &pb.User{
		Id:         1,
		FirstName:  "fn",
		SecondName: "sn",
		Email:      req.GetEmail(),
	}

	// generate jwt token
	tokenStr, err := jwt.GenerateToken(user.Id)
	if err != nil {
		return nil, errors.NewGRPC(
			"failed to generate JWT token: "+err.Error(), // temporary measure REFACTOR
			"auth.server.SignUp() -> jwt.GenerateToken()",
			err.Error(),
			codes.Code(codes.Internal),
		)
	}

	return &pb.SignUpResponse{User: user, TokenString: tokenStr}, nil
}

// SignIn receives credentials for the authorization and return the authorized user
func (s *Server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {

	return nil, nil
}
