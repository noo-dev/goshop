package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"goshop/internal/order/dto"
	"goshop/internal/order/model"
	"goshop/pkg/paging"
	"goshop/pkg/utils"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, userID string, lines []*model.OrderLine) (*model.Order, error)
}

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

func (or *OrderRepo) CreateOrder(ctx context.Context, userID string, lines []*model.OrderLine) (*model.Order, error) {
	var order model.Order

	transactionHandler := func() error {
		//////// CREATE ORDERS ///////////
		var totalPrice float64
		for _, line := range lines {
			totalPrice += line.Price
		}
		order.TotalPrice = totalPrice
		order.UserID = userID

		if err := or.DB.Create(&order).Error; err != nil {
			return err
		}

		//////// CREATE ORDER LINES ///////////
		for _, line := range lines {
			line.OrderID = order.ID
		}
		if err := or.DB.CreateInBatches(&lines, len(lines)).Error; err != nil {
			return err
		}

		utils.Copy(&order.Lines, &lines)
		return nil
	}

	tx := or.DB.Begin()
	if err := transactionHandler(); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (or *OrderRepo) GetOrderByID(ctx context.Context, id string, preload bool) (*model.Order, error) {
	var order model.Order
	if err := or.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (or *OrderRepo) GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, *paging.Pagination, error) {
	if req.Status != "" {

	}

	userID := req.UserId
	limit := req.Limit
	var pagination *paging.Pagination
	or.DB.Where("user_id = ?", userID).Order(clause.Or).Offset().Limit(limit)
}
