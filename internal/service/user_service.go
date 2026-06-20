package service

import (
	"errors"
	"strings"

	"user-service/internal/model"
	"user-service/internal/repository"
)

var (
	ErrNameRequired = errors.New("name is required")
	ErrInvalidAge   = errors.New("age must be greater than 0")
	ErrInvalidEmail = errors.New("a valid email is required")
	ErrUserNotFound = errors.New("user not found")
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(user model.User) (model.User, error) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if user.Name == "" {
		return model.User{}, ErrNameRequired
	}

	if user.Age <= 0 {
		return model.User{}, ErrInvalidAge
	}

	if user.Email == "" || !strings.Contains(user.Email, "@") {
		return model.User{}, ErrInvalidEmail
	}

	return s.repository.Create(user), nil
}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	user, found := s.repository.GetByID(id)
	if !found {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}
