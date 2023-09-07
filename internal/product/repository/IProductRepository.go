package repository

import (
	"context"
	"goshop/internal/product/dto"
	"goshop/internal/product/model"
	"goshop/pkg/paging"
)

type IProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	// Update(ctx context.Context, product *model.Product) error
	ListProducts(ctx context.Context, dto *dto.ListProductReqDto) ([]*model.Product, *paging.Pagination, error)
	// GetProductByID(ctx context.Context, id string) (*model.Product, error)
}
