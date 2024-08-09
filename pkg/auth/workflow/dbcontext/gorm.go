package dbcontext

import (
	"pawpawchat/pkg/auth/repository"

	"gorm.io/gorm"
)

type gormHandler struct {
	db *gorm.DB
	tx *gorm.DB
	ur repository.UserRepository
}

func (h *gormHandler) Begin() error {
	tx := h.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	h.tx = tx
	return nil
}

func (h *gormHandler) Commit() error {
	if h.tx == nil {
		return nil
	}
	if err := h.tx.Commit().Error; err != nil {
		return err
	}
	h.tx = nil
	return nil
}

func (h *gormHandler) Rollback() error {
	if h.tx == nil {
		return nil
	}
	if err := h.tx.Rollback().Error; err != nil {
		return err
	}
	h.tx = nil
	return nil
}

func (h *gormHandler) GetUserRepository() repository.UserRepository {
	if h.ur == nil {
		h.ur = repository.NewGormUserRepository(h.db)
	}
	return h.ur
}
