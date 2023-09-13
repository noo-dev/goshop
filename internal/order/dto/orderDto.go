package dto

import "goshop/pkg/paging"

type OrderDto struct {
	ID         string          `json:"id"`
	Code       string          `json:"code"`
	Lines      []*OrderLineDto `json:"lines"`
	TotalPrice float64         `json:"total_price"`
	Status     string          `json:"status"`
}

type OrderLineDto struct {
	Product  ProductDto `json:"product,omitempty"`
	Quantity uint       `json:"quantity"`
	Price    float64    `json:"price"`
}

type PlaceOrderReqDto struct {
	UserID string                 `json:"user_id" validate:"required"`
	Lines  []PlaceOrderLineReqDto `json:"lines" validate:"required,gt=0,lte=5"`
}

type PlaceOrderLineReqDto struct {
	ProductID string `json:"product_id,omitempty" validate:"required"`
	Quantity  uint   `json:"quantity,omitempty" validate:"required"`
}

type ListOrderReq struct {
	UserId    string `json:"-"`
	Code      string `json:"code,omitempty" form:"code"`
	Status    string `json:"status,omitempty" form:"status"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type ListOrderRes struct {
	Orders     []*OrderDto        `json:"orders,omitempty"`
	Pagination *paging.Pagination `json:"pagination,omitempty"`
}
