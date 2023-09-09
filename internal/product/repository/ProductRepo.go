package repository

import (
	"context"
	"gorm.io/gorm"
	"goshop/internal/product/dto"
	"goshop/internal/product/model"
	"goshop/pkg/config"
	"goshop/pkg/paging"
	"log"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepo {
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatal(err)
	}
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) ListProducts(ctx context.Context, req *dto.ListProductReqDto) ([]*model.Product, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()
	query := r.DB
	order := "created_at"
	if req.Name != "" {
		query = r.DB.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Code != "" {
		query = r.DB.Where("code = ?", req.Code)
	}
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := query.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var products []*model.Product
	err := query.
		Limit(int(pagination.Limit)).
		Order(order).
		Find(&products).Error

	if err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (r *ProductRepo) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()

	var product model.Product
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo) Create(ctx context.Context, product *model.Product) error {
	ctx, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()

	if err := r.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepo) Update(ctx context.Context, product *model.Product) error {
	_, cancel := context.WithTimeout(ctx, config.DATABASE_TIMEOUT)
	defer cancel()

	if err := r.DB.Save(&product).Error; err != nil {
		return err
	}

	return nil
}
