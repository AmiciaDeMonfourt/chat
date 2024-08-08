package repository

import (
	"context"
	"pawpawchat/internal/model/domain"
)

type Profile interface {
	Create(context.Context, *domain.UserBiography) (*domain.User, error)
	GetByID(context.Context, uint64) (*domain.User, error)
}
