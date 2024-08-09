package handler

import (
	"context"
	"fmt"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/profile/controller"
)

type Profile struct {
	profileController *controller.Profile
}

func NewProfile(profile *controller.Profile) *Profile {
	return &Profile{profile}
}

func (p *Profile) Create(ctx context.Context, req *profilepb.CreateRequest) (*profilepb.CreateResponse, error) {
	biography, err := generateUserBiographyFromCreateRequest(req)
	if err != nil {
		return nil, err
	}

	createdProfile, err := p.profileController.Create(ctx, biography)
	if err != nil {
		return nil, err
	}

	return generateCreateResponse(createdProfile), nil
}

func generateUserBiographyFromCreateRequest(req *profilepb.CreateRequest) (*domain.UserBiography, error) {
	if req.GetUserbio() != nil {
		firstname := req.GetUserbio().GetFirstname()
		if firstname == "" {
			return nil, fmt.Errorf("missing first name")
		}

		secondname := req.GetUserbio().GetSecondname()
		if secondname == "" {
			return nil, fmt.Errorf("missing second name")
		}

		return &domain.UserBiography{
			FirstName:  firstname,
			SecondName: secondname,
		}, nil
	}

	return nil, fmt.Errorf("missing user biography")
}

func generateCreateResponse(user *domain.User) *profilepb.CreateResponse {
	return &profilepb.CreateResponse{
		User: &profilepb.User{
			Userid: user.Biography.UserID,
			Userbio: &profilepb.UserBiography{
				Firstname:  user.Biography.FirstName,
				Secondname: user.Biography.SecondName,
			},
		},
	}
}
