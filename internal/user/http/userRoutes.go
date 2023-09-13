package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"goshop/internal/user/handler"
	"goshop/internal/user/repository"
	"goshop/internal/user/service"
	"goshop/pkg/middleware"
)

func InitUserRoutes(r *gin.RouterGroup, DB *gorm.DB, validator *validator.Validate) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(validator, userRepo)
	userHandler := handler.NewUserHandler(userService)

	authMiddleware := middleware.JwtAuthMiddleware()
	refreshAuthMiddleware := middleware.JwtRefreshMiddleware()

	authRoute := r.Group("/auth")
	authRoute.POST("/register", userHandler.Register)
	authRoute.POST("/login", userHandler.Login)
	authRoute.GET("/me", authMiddleware, userHandler.GetMe)
	authRoute.GET("/refresh", refreshAuthMiddleware, userHandler.RefreshToken)
	authRoute.PUT("/change-password", authMiddleware, userHandler.ChangePassword)
}
