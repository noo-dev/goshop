package service

import (
	"context"
	"goshop/internal/product/dto"
	"goshop/internal/product/model"
	"goshop/pkg/paging"
)

type IProductService interface {
	ListProducts(ctx context.Context, req *dto.ListProductReqDto) ([]*model.Product, *paging.Pagination, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	Create(ctx context.Context, req *dto.CreateProductReqDto) (*model.Product, error)
	Update(ctx context.Context, id string, req *dto.UpdateProductReqDto) (*model.Product, error)
}
