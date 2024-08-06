package database

import (
	"context"
	"pawpawchat/internal/model/domain"
)

type Database interface {
	CreateProfile(context.Context, *domain.UserPersonalInfo) (*domain.User, error)
}
