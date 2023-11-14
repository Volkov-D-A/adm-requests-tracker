package service

import (
	"context"
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
)

type UserStorage interface {
	Create(ctx context.Context, user *models.User) (string, error)
}

type userService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *userService {
	return &userService{
		userStorage: userStorage,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User, token string) (string, error) {
	fmt.Println(utils.HashPassword(user.Password))
	return "1", nil
}
