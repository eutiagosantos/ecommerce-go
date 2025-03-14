package repositories

import (
	"api-golang/internal/core/domain"
)

type UserRepositoryInterface interface {
	Create(user *domain.User) error
	FindAll() ([]domain.User, error)
	FindById(id string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
	FindByEmail(email string) (*domain.User, error)
	// Login(email string, password string) (*domain.User, error)
}
