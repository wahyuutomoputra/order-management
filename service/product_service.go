package service

import (
	"github.com/wahyuutomoputra/order-management/models"
	"github.com/wahyuutomoputra/order-management/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) FindAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) FindByID(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(product *models.Product) error {
	return s.repo.Delete(product)
}
