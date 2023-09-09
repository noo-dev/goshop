package service

import (
	"context"
	"goshop/internal/user/dto"
	"goshop/internal/user/model"
)

type IUserService interface {
	Register(ctx context.Context, dto *dto.RegisterReqDto) (*model.User, error)
}
