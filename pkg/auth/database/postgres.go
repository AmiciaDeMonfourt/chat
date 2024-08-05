package database

import (
	"context"
	"pawpawchat/pkg/auth/model"

	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPosgresDB(db *gorm.DB) Database {
	return &PostgresDB{
		db: db,
	}
}

func (p *PostgresDB) InsertUserCredentials(ctx context.Context, cr *model.UserCredentials) error {
	return p.db.WithContext(ctx).Create(cr).Error
}

func (p *PostgresDB) CheckUserCredentials(ctx context.Context, cr *model.UserCredentials) error {
	return p.db.WithContext(ctx).Where(cr).First(nil).Error
}
