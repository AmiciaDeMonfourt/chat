package sql

import (
	"context"
	"pawpawchat/internal/model/domain"

	"github.com/jmoiron/sqlx"
)

type sqlxUserRepository struct {
	db *sqlx.DB
}

func NewSqlxUserRepository(db *sqlx.DB) *sqlxUserRepository {
	return &sqlxUserRepository{db}
}

func (r *sqlxUserRepository) Create(ctx context.Context, credentials *domain.UserCredentials) error {
	return nil
}
