package repository

import (
	"github.com/wahyuutomoputra/order-management/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *OrderRepository) FindProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *OrderRepository) WithTx(tx *gorm.DB) *OrderRepository {
	return &OrderRepository{db: tx}
}

func (r *OrderRepository) DB() *gorm.DB {
	return r.db
}
