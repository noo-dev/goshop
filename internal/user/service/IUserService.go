package service

import (
	"context"
	"goshop/internal/user/dto"
	"goshop/internal/user/model"
)

type IUserService interface {
	Register(ctx context.Context, dto *dto.RegisterReqDto) (*model.User, error)
	Login(ctx context.Context, req *dto.LoginReq) (*model.User, string, string, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	RefreshToken(ctx context.Context, userID string) (string, error)
	ChangePassword(ctx context.Context, id string, req *dto.ChangePasswordReqDto) error
}
