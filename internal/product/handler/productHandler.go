package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/quangdangfit/gocommon/logger"
	"goshop/internal/product/dto"
	"goshop/internal/product/service"
	"goshop/pkg/response"
	"goshop/pkg/utils"
	"net/http"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(
	service service.IProductService,
) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (p *ProductHandler) ListProducts(c *gin.Context) {
	var req dto.ListProductReqDto
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	var res dto.ListProductResDto
	products, pagination, err := p.service.ListProducts(c, &req)
	if err != nil {
		logger.Error("Failed to get list of products: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}
	utils.Copy(&res.Products, &products)
	res.Pagination = pagination
	response.Success(c, http.StatusOK, res)
}

func (p *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductReqDto
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, 400, err, "invalid parameters")
		return
	}

	product, err := p.service.Create(c, &req)
	if err != nil {
		logger.Errorf("Failed to create product: %s\n", err.Error())

		if errors.As(err, &validator.ValidationErrors{}) {
			var errorBag []string
			for _, e := range err.(validator.ValidationErrors) {
				errorBag = append(errorBag, fmt.Sprintf("Error: validation for '%s' failed on the '%s' tag", e.Field(), e.Tag()))
			}
			c.JSON(400, gin.H{
				"errors":  errorBag,
				"result":  nil,
				"message": "Validation errors",
			})
			// response.Error(c, 400, err, "Validation error")
		} else {
			response.Error(c, 500, err, "Something went wrong")
		}
		return
	}

	var res dto.ProductDto
	utils.Copy(&res, &product)
	response.Success(c, 200, res)

}
