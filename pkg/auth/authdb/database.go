package authdb

import (
	"context"
	"pawpawchat/internal/model/domain"

	"gorm.io/gorm"
)

// Database defines the interface for database operations
type Database interface {
	//InsertUserCredentials ...
	InsertUserCredentials(context.Context, *gorm.DB, *domain.UserCredentials) error

	// CheckUserCredentials ...
	CheckUserCredentials(context.Context, *domain.UserCredentials) error

	// CheckUserCredentials
	Begin() *gorm.DB
}
