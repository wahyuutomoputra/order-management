package service

import (
	"errors"

	"github.com/wahyuutomoputra/order-management/dto"
	"github.com/wahyuutomoputra/order-management/models"
	"github.com/wahyuutomoputra/order-management/repository"
	"gorm.io/gorm"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) CreateOrder(userID uint, items []dto.OrderItemInput) (*models.Order, error) {
	var resultOrder *models.Order
	err := s.repo.DB().Transaction(func(tx *gorm.DB) error {
		repoTx := s.repo.WithTx(tx)
		order := models.Order{
			UserID:    userID,
			CreatedAt: int64(0), // set di handler
		}
		var orderItems []models.OrderItem
		for _, item := range items {
			product, err := repoTx.FindProductByID(item.ProductID)
			if err != nil {
				return errors.New("product not found")
			}
			if product.Stock < item.Quantity {
				return errors.New("Insufficient stock for product: " + product.Name)
			}
			product.Stock -= item.Quantity
			if err := repoTx.UpdateProduct(product); err != nil {
				return err
			}
			orderItems = append(orderItems, models.OrderItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			})
		}
		order.Items = orderItems
		if err := repoTx.Create(&order); err != nil {
			return err
		}
		resultOrder = &order
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resultOrder, nil
}

func (s *OrderService) GetOrderHistory(userID uint) ([]models.Order, error) {
	return s.repo.FindByUser(userID)
}
