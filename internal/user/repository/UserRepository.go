package repository

import (
	"context"
	"gorm.io/gorm"
	"goshop/internal/user/model"
	"goshop/pkg/config"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()

	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()

	if err := r.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()

	var user model.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()

	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
