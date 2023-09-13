package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	"goshop/internal/user/dto"
	"goshop/internal/user/service"
	"goshop/pkg/response"
	"goshop/pkg/utils"
	"net/http"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterReqDto
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, 400, err, "Invalid parameters")
		return
	}

	user, err := uh.service.Register(c, &req)
	if err != nil {
		logger.Error(err)
		response.Error(c, 500, err, "Something went wrong")
		return
	}

	var res dto.RegisterResDto
	utils.Copy(&res.User, &user)

	response.Success(c, http.StatusCreated, res)
}

func (uh *UserHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	user, accessToken, refreshToken, err := uh.service.Login(c, &req)
	if err != nil {
		logger.Error("Failed to login: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.LoginRes
	utils.Copy(&res.User, &user)
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	response.Success(c, http.StatusOK, res)
}

func (uh *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("Supply user id"), "Unauthorized")
		return
	}

	user, err := uh.service.GetUserByID(c, userID)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.UserDto
	utils.Copy(&res, &user)
	response.Success(c, http.StatusOK, res)
}

func (uh *UserHandler) RefreshToken(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	accessToken, err := uh.service.RefreshToken(c, userID)
	if err != nil {
		logger.Error("Failed to refresh token", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
	}

	res := dto.RefreshTokenResDto{
		AccessToken: accessToken,
	}
	response.Success(c, http.StatusOK, res)
}

func (uh *UserHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordReqDto
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get request body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userID := c.GetString("userId")
	err := uh.service.ChangePassword(c, userID, &req)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}
	response.Success(c, http.StatusOK, nil)
}
