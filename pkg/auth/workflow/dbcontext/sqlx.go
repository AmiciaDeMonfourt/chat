package dbcontext

import (
	"pawpawchat/pkg/auth/repository"

	"github.com/jmoiron/sqlx"
)

type sqlxHandler struct {
	db *sqlx.DB
	tx *sqlx.DB
}

func (h *sqlxHandler) Begin() error {
	_ = h.tx
	return nil
}

func (h *sqlxHandler) Commit() error {
	return nil
}

func (h *sqlxHandler) Rollback() error {
	return nil
}

func (h *sqlxHandler) GetUserRepository() repository.UserRepository {
	return nil
}
