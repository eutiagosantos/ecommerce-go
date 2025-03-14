package usecases

import (
	"api-golang/infra/repositories"
	"api-golang/internal/core/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCases struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserUseCases(userRepository repositories.UserRepositoryInterface) *UserUseCases {
	return &UserUseCases{UserRepository: userRepository}
}

func (u *UserUseCases) FindByEmail(email string) (*domain.User, error) {
	userFound, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return userFound, nil
}

func (u *UserUseCases) Create(user *domain.User) error {
	userFound, _ := u.FindByEmail(user.Email)

	if userFound != nil {
		return fmt.Errorf("user already exists")
	}
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = passwordHash

	return u.UserRepository.Create(user)
}

func (u *UserUseCases) FindAll() ([]domain.User, error) {
	return u.UserRepository.FindAll()
}

func (u *UserUseCases) FindById(id string) (*domain.User, error) {
	return u.UserRepository.FindById(id)
}

func (u *UserUseCases) Update(user *domain.User) error {
	return u.UserRepository.Update(user)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
