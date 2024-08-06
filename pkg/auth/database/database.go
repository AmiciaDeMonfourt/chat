package database

import (
	"context"
	"pawpawchat/internal/model/domain"

	"gorm.io/gorm"
)

// Database defines the interface for database operations
type Database interface {
	InsertUserCredentials(context.Context, *gorm.DB, *domain.UserCredentials) error
	CheckUserCredentials(context.Context, *domain.UserCredentials) error
	Begin() *gorm.DB
}
