package database

import (
	"context"
	"pawpawchat/internal/model/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB(db *gorm.DB) Database {
	return &PostgresDB{db: db}
}

func (p *PostgresDB) CreateProfile(ctx context.Context, userinfo *domain.UserPersonalInfo) (*domain.User, error) {
	tx := p.db.WithContext(ctx).Clauses(clause.Returning{}).Create(userinfo)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &domain.User{PersonalInfo: *userinfo}, nil
}
