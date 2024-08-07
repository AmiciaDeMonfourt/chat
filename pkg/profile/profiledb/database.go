package profiledb

import (
	"context"
	"pawpawchat/internal/model/domain"
)

type Database interface {
	CreateProfile(context.Context, *domain.UserBiography) (*domain.User, error)
	GetProfileByID(context.Context, uint64) (*domain.User, error)
}

type Factory interface {
	OpenProfileDB(dsn string) (Database, error)
}
