package dbcontext

import (
	"fmt"
	"pawpawchat/pkg/auth/repository"
	"reflect"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type DBContext interface {
	GetUserRepository() repository.UserRepository
	Begin() error
	Commit() error
	Rollback() error
}

func MustNew(db any) DBContext {
	switch db := db.(type) {
	case *gorm.DB:
		return &gormHandler{db: db}
	case *sqlx.DB:
		return &sqlxHandler{db: db}
	default:
		panic(fmt.Errorf("unknown database type: %v", (reflect.TypeOf(db))))
	}
}
