package repository

import (
	"context"
	"errors"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/auth/repository/orm"
	"pawpawchat/pkg/auth/repository/sql"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(context.Context, *domain.UserCredentials) error
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return orm.NewGormUserRepository(db)
}

func NewSqlxUserRepository(db *sqlx.DB) UserRepository {
	return sql.NewSqlxUserRepository(db)
}

func NewUserRepository(db any) (UserRepository, error) {
	switch db := db.(type) {
	case *gorm.DB:
		return NewGormUserRepository(db), nil
	case *sqlx.DB:
		return NewSqlxUserRepository(db), nil
	default:
		return nil, errors.New("unsupported database type")
	}
}
