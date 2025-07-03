package service

import (
	"github.com/wahyuutomoputra/order-management/models"
	"github.com/wahyuutomoputra/order-management/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Register(user *models.User) error {
	// Bisa tambahkan business logic lain di sini
	return s.repo.Create(user)
}

func (s *UserService) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := s.repo.DB().First(&user, id).Error
	return &user, err
}
