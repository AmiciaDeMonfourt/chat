package orm

import (
	"context"
	"pawpawchat/internal/model/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *gormUserRepository {
	return &gormUserRepository{db}
}

func (r *gormUserRepository) Create(ctx context.Context, ucr *domain.UserCredentials) error {
	return r.db.WithContext(ctx).Clauses(clause.Returning{Columns: []clause.Column{clause.PrimaryColumn}}).Create(ucr).Error
}
