package repository

import (
	"context"
	"goshop/internal/user/model"
)

type IUserRepository interface {
	Create(context.Context, *model.User) error
	/*Update(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)*/
}
