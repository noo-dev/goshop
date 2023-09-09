package handler

import (
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
	}

	var res dto.RegisterResDto
	utils.Copy(&res.User, &user)

	response.Success(c, http.StatusCreated, res)
}
