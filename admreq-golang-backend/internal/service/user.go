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

func (s *userService) Create(user *models.UserCreate, ut *models.UserToken) (string, error) {
	if ut.Role != "admin" {
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

func (s *userService) Delete(uuid string, ut *models.UserToken) error {
	if ut.Role != "admin" {
		return models.ErrUnauthorized
	}
	err := s.userStorage.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUsers(ut *models.UserToken) ([]models.UserResponse, error) {
	resp, _ := s.userStorage.GetUsers()
	return resp, nil
}
