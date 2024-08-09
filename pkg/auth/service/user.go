package service

import (
	"context"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/auth/workflow"
	"pawpawchat/utils/encrypt"
)

type UserService interface {
	SignUp(context.Context, *domain.UserCredentials) error
}

func NewUserService(unit workflow.UnitOfWorkflow) UserService {
	return &userService{unit}
}

type userService struct {
	unitOfWork workflow.UnitOfWorkflow
}

func (s *userService) SignUp(ctx context.Context, user *domain.UserCredentials) error {
	hashpass, err := encrypt.EncryptString(user.Password)
	if err != nil {
		return err
	}

	user.Password = ""
	user.HashPass = hashpass

	return s.unitOfWork.WithTransaction(func() error {
		if err := s.unitOfWork.GetUserRepository().Create(ctx, user); err != nil {
			return err
		}
		return nil
	})
}
