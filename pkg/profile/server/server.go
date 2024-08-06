package server

import (
	"context"
	"log/slog"
	pb "pawpawchat/generated/proto/profilepb"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/profile/database"
)

type ProfileServer struct {
	pb.UnimplementedProfileServer
	db database.Database
}

func New(db database.Database) *ProfileServer {
	return &ProfileServer{db: db}
}

func (s *ProfileServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	userinfo := &domain.UserPersonalInfo{
		FirstName:  req.GetUserinfo().GetFirstname(),
		SecondName: req.GetUserinfo().GetSecondname(),
	}

	user, err := s.db.CreateProfile(ctx, userinfo)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateResponse{
		User: &pb.User{
			Userid: userinfo.UserID,
			Userinfo: &pb.UserPersonalInfo{
				Firstname:  user.PersonalInfo.FirstName,
				Secondname: user.PersonalInfo.SecondName,
			},
		},
	}

	slog.Info("new profile has been created", "info", resp)

	return resp, nil
}
