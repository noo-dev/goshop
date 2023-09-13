package repository

import (
	"context"
	"gorm.io/gorm"
	"goshop/internal/product/model"
)

type IProductRepository interface {
	GetProductByID(ctx context.Context, id string) (*model.Product, error)
}

type ProductRepo struct {
	db *gorm.DB
}

func (r *ProductRepo) GetProductByID(ctx context.Context, id string) (*model.Product, error) {

	var product model.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
