package dto

import (
	"goshop/pkg/paging"
	"time"
)

type ProductDto struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListProductReqDto struct {
	Name      string `json:"name,omitempty" form:"name"`
	Code      string `json:"code,omitempty" form:"code""`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type ListProductResDto struct {
	Products   []*ProductDto      `json:"products"`
	Pagination *paging.Pagination `json:"pagination"`
}

type CreateProductReqDto struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price"`
}

type UpdateProductReqDto struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty" validate:"gte=0"`
}
