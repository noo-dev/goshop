package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"goshop/internal/product/handler"
	"goshop/internal/product/repository"
	"goshop/internal/product/service"
)

func InitProductRoutes(r *gin.RouterGroup, DB *gorm.DB, validator *validator.Validate) {
	productRepo := repository.NewProductRepository(DB)
	productSvc := service.NewProductService(validator, productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	productRoutes := r.Group("/products")
	{
		productRoutes.GET("", productHandler.ListProducts)
		productRoutes.GET("/:id", productHandler.GetProductById)
		productRoutes.POST("", productHandler.CreateProduct)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
	}
}
