package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/quangdangfit/gocommon/logger"
	"goshop/internal/user/dto"
	"goshop/internal/user/model"
	"goshop/internal/user/repository"
	"goshop/pkg/utils"
)

type UserService struct {
	validator *validator.Validate
	repo      repository.IUserRepository
}

func NewUserService(
	validator *validator.Validate,
	repo repository.IUserRepository,
) *UserService {
	return &UserService{
		validator: validator,
		repo:      repo,
	}
}

func (u *UserService) Register(ctx context.Context, req *dto.RegisterReqDto) (*model.User, error) {
	if err := u.validator.Struct(req); err != nil {
		return nil, err
	}

	var user model.User
	utils.Copy(&user, req)
	err := u.repo.Create(ctx, &user)
	if err != nil {
		logger.Errorf("Register.Create fail, error: %v", err)
		return nil, err
	}

	return &user, nil
}
