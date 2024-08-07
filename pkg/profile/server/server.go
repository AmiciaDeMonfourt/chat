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
	userinfo := &domain.UserBiography{
		FirstName:  req.GetUserbio().GetFirstname(),
		SecondName: req.GetUserbio().GetSecondname(),
	}

	_, err := s.db.CreateProfile(ctx, userinfo)
	if err != nil {
		return nil, err
	}

	user, err := s.db.GetProfileByID(ctx, uint64(userinfo.UserID))
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateResponse{
		User: &pb.User{
			Userid: userinfo.UserID,
			Userbio: &pb.UserBiography{
				Firstname:  user.Biography.FirstName,
				Secondname: user.Biography.SecondName,
			},
		},
	}

	slog.Info("new profile has been created", "info", resp)

	return resp, nil
}
