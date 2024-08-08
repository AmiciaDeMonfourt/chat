package controller

import (
	"context"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/profile/repository"
)

type Profile struct {
	profile repository.Profile
}

func NewProfile(profile repository.Profile) *Profile {
	return &Profile{profile}
}

func (c *Profile) Create(ctx context.Context, bio *domain.UserBiography) (*domain.User, error) {
	return c.profile.Create(ctx, bio)
}

func (c *Profile) GetByID(ctx context.Context, id uint64) (*domain.User, error) {
	return c.profile.GetByID(ctx, id)
}
