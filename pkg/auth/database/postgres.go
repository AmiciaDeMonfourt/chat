package database

import (
	"context"
	"pawpawchat/internal/model/domain"

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

func (p *PostgresDB) InsertUserCredentials(ctx context.Context, tx *gorm.DB, cr *domain.UserCredentials) error {
	return tx.WithContext(ctx).Create(cr).Error
}

func (p *PostgresDB) CheckUserCredentials(ctx context.Context, cr *domain.UserCredentials) error {
	return p.db.WithContext(ctx).Where(cr).First(nil).Error
}

func (p *PostgresDB) Begin() *gorm.DB {
	return p.db.Begin()
}
