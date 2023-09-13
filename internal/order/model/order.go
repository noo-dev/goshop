package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"goshop/pkg/utils"
	"time"
)

type OrderStatus string

const (
	ORDER_STATUS_NEW         OrderStatus = "new"
	ORDER_STATUS_IN_PROGRESS OrderStatus = "in-progress"
	ORDER_STATUS_DONE        OrderStatus = "done"
	ORDER_STATUS_CANCELLED   OrderStatus = "cancelled"
)

type Order struct {
	ID         string    `json:"id" gorm:"unique;not null;index;primary_key"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"index"`
	Code       string    `json:"code"`
	UserID     string    `json:"user_id"`
	User       *User
	Lines      []*OrderLine `json:"lines"`
	TotalPrice float64      `json:"total_price"`
	Status     OrderStatus  `json:"status"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New().String()
	order.Code = utils.GenerateCode("SO")

	if order.Status == "" {
		order.Status = ORDER_STATUS_NEW
	}

	return nil
}
