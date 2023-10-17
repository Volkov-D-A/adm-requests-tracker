package service

import (
	"context"
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type Config struct {
	Key string
}

type TSRRepository interface {
	Create(ctx context.Context, user *models.User) (int, error)
}

type TSRService struct {
	userRepository TSRRepository
	conf           *Config
}

func New(userRepository TSRRepository, conf *Config) *TSRService {
	return &TSRService{
		userRepository: userRepository,
		conf:           conf,
	}
}

func (s *TSRService) CreateUser(ctx context.Context, user *models.User, token string) (int32, error) {
	fmt.Println(hashPassword(user.Password))
	return 1, nil
}
