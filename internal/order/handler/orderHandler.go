package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	"goshop/internal/order/dto"
	"goshop/internal/order/service"
	"goshop/pkg/response"
	"goshop/pkg/utils"
	"net/http"
)

type OrderHandler struct {
	service service.IOrderService
}

func NewOrderHandler(service service.IOrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req dto.PlaceOrderReqDto
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err, "Invalid parameters")
		response.Error(c, 400, err, "Invalid parameters")
		return
	}

	req.UserID = c.GetString("userId")
	if req.UserID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	order, err := h.service.SaveOrder(c, &req)
	if err != nil {
		logger.Error("Failed to create order: ", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.OrderDto
	utils.Copy(&res, &order)
	response.Success(c, http.StatusOK, res)
}
