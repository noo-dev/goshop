package repository

import (
	"context"
	"gorm.io/gorm"
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
	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
	}
	var total int64
	if err := or.DB.Model(&model.Order{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var pagination = *paging.New(req.Page, req.Limit, total)

	var orders []*model.Order
	err := or.DB.Where("user_id = ?", req.UserId).
		Offset(int(pagination.Skip)).
		Limit(int(pagination.Limit)).
		Order(order).
		Find(&orders).
		Error
	if err != nil {
		return nil, nil, err
	}

	return orders, &pagination, nil
}
