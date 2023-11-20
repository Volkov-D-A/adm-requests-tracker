package service

import (
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type UserStorage interface {
	Create(user *models.UserCreate) (string, error)
	Auth(user *models.UserAuth) (*models.UserResponse, error)
	Delete(uuid string) error
	GetUsers() ([]models.UserResponse, error)
}

type userService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *userService {
	return &userService{
		userStorage: userStorage,
	}
}

func (s *userService) Create(user *models.UserCreate, ur *models.UserRole) (string, error) {
	if ur.Role != "admin" {
		return "", models.ErrUnauthorized
	}
	uuid, err := s.userStorage.Create(user)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (s *userService) Auth(user *models.UserAuth) (*models.UserResponse, error) {
	resp, err := s.userStorage.Auth(user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) Delete(uuid string, ur *models.UserRole) error {
	if ur.Role != "admin" {
		return models.ErrUnauthorized
	}
	err := s.userStorage.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUsers(ur *models.UserRole) ([]models.UserResponse, error) {
	resp, _ := s.userStorage.GetUsers()
	return resp, nil
}
