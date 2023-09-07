package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/quangdangfit/gocommon/logger"
	"goshop/internal/product/dto"
	"goshop/internal/product/model"
	"goshop/internal/product/repository"
	"goshop/pkg/paging"
	"goshop/pkg/utils"
)

type ProductService struct {
	repo      repository.IProductRepository
	validator *validator.Validate
}

func NewProductService(
	validator *validator.Validate,
	repo repository.IProductRepository,
) *ProductService {
	return &ProductService{
		repo:      repo,
		validator: validator,
	}
}

func (p *ProductService) ListProducts(ctx context.Context, req *dto.ListProductReqDto) ([]*model.Product, *paging.Pagination, error) {
	products, pagination, err := p.repo.ListProducts(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (p *ProductService) Create(ctx context.Context, req *dto.CreateProductReqDto) (*model.Product, error) {

	if err := p.validator.Struct(req); err != nil {

		return nil, err
	}

	var product model.Product
	utils.Copy(&product, req)

	err := p.repo.Create(ctx, &product)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &product, nil
}
