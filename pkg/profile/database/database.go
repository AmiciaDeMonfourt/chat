package database

import (
	"context"
	"pawpawchat/internal/model/domain"
)

type Database interface {
	CreateProfile(context.Context, *domain.UserBiography) (*domain.User, error)
	GetProfileByID(context.Context, uint64) (*domain.User, error)
}
