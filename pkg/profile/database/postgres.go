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

func (p *PostgresDB) CreateProfile(ctx context.Context, userinfo *domain.UserBiography) (*domain.User, error) {
	return &domain.User{Biography: *userinfo}, p.db.WithContext(ctx).Clauses(clause.Returning{}).Create(userinfo).Error
}

func (p *PostgresDB) GetProfileByID(ctx context.Context, id uint64) (*domain.User, error) {
	var user domain.User
	if err := p.db.Table("user_biographies").First(&user, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
