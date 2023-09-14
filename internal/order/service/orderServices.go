package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"goshop/internal/order/dto"
	"goshop/internal/order/model"
	"goshop/internal/order/repository"
	"goshop/pkg/customerrors"
	"goshop/pkg/paging"
	"goshop/pkg/utils"
)

type IOrderService interface {
	SaveOrder(ctx context.Context, req *dto.PlaceOrderReqDto) (*model.Order, error)
	GetOrderByID(ctx context.Context, id string) (*model.Order, error)
	GetMyOrders(ctx context.Context, orderID, userID string) (*model.Order, error)
}

type OrderService struct {
	validator   validator.Validate
	repo        repository.IOrderRepository
	productRepo repository.IProductRepository
}

func NewOrderService(
	validator validator.Validate,
	repo repository.IOrderRepository,
	productRepo repository.IProductRepository,
) *OrderService {
	return &OrderService{
		validator:   validator,
		repo:        repo,
		productRepo: productRepo,
	}
}

func (s *OrderService) SaveOrder(ctx context.Context, req *dto.PlaceOrderReqDto) (*model.Order, error) {

	err := s.validator.Struct(&req)
	if err != nil {
		return nil, &customerrors.ClientError{WrappedError: err}
	}

	var orderLines []*model.OrderLine
	utils.Copy(&orderLines, &req.Lines)

	productMap, err := FetchAndMapProductsToOrderLines(ctx, s.productRepo, orderLines)
	if err != nil {
		return nil, err
	}

	order, err := s.repo.CreateOrder(ctx, req.UserID, orderLines)
	if err != nil {
		return nil, err
	}

	for _, line := range order.Lines {
		line.Product = productMap[line.ProductID]
	}

	return order, nil
}

func (s *OrderService) GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, *paging.Pagination, error) {
	return nil, nil, nil
}

func FetchAndMapProductsToOrderLines(
	ctx context.Context,
	repo repository.IProductRepository, orderLines []*model.OrderLine) (map[string]*model.Product, error) {
	productMap := make(map[string]*model.Product)
	for _, line := range orderLines {
		product, err := repo.GetProductByID(ctx, line.ProductID)
		if err != nil {
			return nil, err
		}
		line.Price = product.Price * float64(line.Quantity)
		productMap[line.ProductID] = product
	}
	return productMap, nil
}
