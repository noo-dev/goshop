package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"goshop/internal/user/handler"
	"goshop/internal/user/repository"
	"goshop/internal/user/service"
)

func InitUserRoutes(r *gin.RouterGroup, DB *gorm.DB, validator *validator.Validate) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator, userRepo)
	userHandler := handler.NewUserHandler(userService)

	users := r.Group("/users")
	users.POST("/register", userHandler.Register)
}
