package database

import (
	"context"
	"pawpawchat/pkg/auth/model"
)

// Database defines the interface for database operations
type Database interface {
	InsertUserCredentials(context.Context, *model.UserCredentials) error
	CheckUserCredentials(context.Context, *model.UserCredentials) error
}
