package service

import (
	"errors"

	"github.com/wahyuutomoputra/order-management/models"
	"github.com/wahyuutomoputra/order-management/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo}
}

type OrderItemInput struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}

func (s *OrderService) CreateOrder(userID uint, items []OrderItemInput) (*models.Order, error) {
	order := models.Order{
		UserID:    userID,
		CreatedAt: int64(0), // set di handler
	}
	var orderItems []models.OrderItem
	for _, item := range items {
		product, err := s.repo.FindProductByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}
		if product.Stock < item.Quantity {
			return nil, errors.New("Insufficient stock for product: " + product.Name)
		}
		product.Stock -= item.Quantity
		s.repo.UpdateProduct(product)
		orderItems = append(orderItems, models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}
	order.Items = orderItems
	return &order, s.repo.Create(&order)
}

func (s *OrderService) GetOrderHistory(userID uint) ([]models.Order, error) {
	return s.repo.FindByUser(userID)
}
