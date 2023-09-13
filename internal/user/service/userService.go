package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/quangdangfit/gocommon/logger"
	"golang.org/x/crypto/bcrypt"
	"goshop/internal/user/dto"
	"goshop/internal/user/model"
	"goshop/internal/user/repository"
	"goshop/pkg/jtoken"
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

func (us *UserService) Register(ctx context.Context, req *dto.RegisterReqDto) (*model.User, error) {
	if err := us.validator.Struct(req); err != nil {
		return nil, err
	}

	var user model.User
	utils.Copy(&user, req)
	err := us.repo.Create(ctx, &user)
	if err != nil {
		logger.Errorf("Register.Create fail, error: %v", err)
		return nil, err
	}

	return &user, nil
}

func (us *UserService) Login(ctx context.Context, req *dto.LoginReq) (*model.User, string, string, error) {
	if err := us.validator.Struct(req); err != nil {
		return nil, "", "", err
	}
	user, err := us.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
		return nil, "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", errors.New("wrong password")
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}

	accessToken := jtoken.GenerateAccessToken(tokenData)
	refreshToken := jtoken.GenerateRefreshToken(tokenData)

	return user, accessToken, refreshToken, nil
}

func (us *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := us.repo.GetUserById(ctx, id)
	if err != nil {
		logger.Errorf("UserService.GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}
	return user, nil
}

func (us *UserService) RefreshToken(ctx context.Context, userID string) (string, error) {
	user, err := us.repo.GetUserById(ctx, userID)
	if err != nil {
		logger.Errorf("RefreshToken.GetUserByID fail, id: %s, error: %s", userID, err)
		return "", err
	}

	tokenClaims := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}
	accessToken := jtoken.GenerateAccessToken(tokenClaims)
	return accessToken, nil
}

func (us *UserService) ChangePassword(ctx context.Context, id string, req *dto.ChangePasswordReqDto) error {
	if err := us.validator.Struct(req); err != nil {
		return err
	}
	user, err := us.repo.GetUserById(ctx, id)
	if err != nil {
		logger.Errorf("ChangePassword.GetUserById fail, id: %s, error: %s", id, err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New("wrong password")
	}

	user.Password = utils.HashAndSalt([]byte(req.NewPassword))
	err = us.repo.Update(ctx, user)
	if err != nil {
		logger.Errorf("ChangePassword.Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}
