package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wahyuutomoputra/order-management/dto"
	"github.com/wahyuutomoputra/order-management/service"
	"github.com/wahyuutomoputra/order-management/utils"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

// CreateOrderHandler godoc
// @Summary Create order
// @Tags Order
// @Accept json
// @Produce json
// @Param data body dto.OrderRequest true "Order data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Security BearerAuth
// @Router /orders [post]
func (h *OrderHandler) CreateOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.OrderRequest
		if err := c.ShouldBindJSON(&req); err != nil || len(req.Items) == 0 {
			utils.JSONError(c, 400, "Invalid request")
			return
		}
		userID, _ := c.Get("userID")
		order, err := h.OrderService.CreateOrder(userID.(uint), req.Items)
		if err != nil {
			utils.JSONError(c, 400, err.Error())
			return
		}
		order.CreatedAt = time.Now().Unix()
		utils.JSONCreated(c, order, "Order created")
	}
}

// OrderHistoryHandler godoc
// @Summary Get order history
// @Tags Order
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security BearerAuth
// @Router /orders/history [get]
func (h *OrderHandler) OrderHistoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		orders, err := h.OrderService.GetOrderHistory(userID.(uint))
		if err != nil {
			utils.JSONError(c, 500, "Failed to get orders")
			return
		}
		utils.JSONSuccess(c, orders, "Order history")
	}
}
