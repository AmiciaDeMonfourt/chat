package orm

import (
	"context"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/profile/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresProfileRepository struct {
	db *gorm.DB
}

func NewPostgresProfileRepository(db *gorm.DB) repository.Profile {
	return &PostgresProfileRepository{db}
}

func (p *PostgresProfileRepository) Create(ctx context.Context, userinfo *domain.UserBiography) (*domain.User, error) {
	return &domain.User{Biography: *userinfo}, p.db.WithContext(ctx).Clauses(clause.Returning{}).Create(userinfo).Error
}

func (p *PostgresProfileRepository) GetByID(ctx context.Context, id uint64) (*domain.User, error) {
	var user domain.User
	if err := p.db.Table("user_biographies").First(&user, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
