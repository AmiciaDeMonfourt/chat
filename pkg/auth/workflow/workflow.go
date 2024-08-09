package workflow

import (
	"pawpawchat/pkg/auth/repository"
	"pawpawchat/pkg/auth/workflow/dbcontext"
)

type UnitOfWorkflow interface {
	GetUserRepository() repository.UserRepository
	WithTransaction(func() error) error
}

type unitOfWorkflow struct {
	db dbcontext.DBContext
}

func NewUnitOfWorkflow(db dbcontext.DBContext) UnitOfWorkflow {
	return &unitOfWorkflow{db: db}
}

func (u *unitOfWorkflow) GetUserRepository() repository.UserRepository {
	return u.db.GetUserRepository()
}

func (u *unitOfWorkflow) WithTransaction(foo func() error) error {
	var err error
	if err = u.db.Begin(); err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = u.db.Rollback()
		}
	}()

	if err = foo(); err != nil {
		if err := u.db.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := u.db.Commit(); err != nil {
		if err = u.db.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil
}
