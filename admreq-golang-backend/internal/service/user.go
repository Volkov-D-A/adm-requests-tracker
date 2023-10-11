package service

import (
	"context"
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"github.com/volkov-d-a/adm-requests-tracker/internal/repository"
)

var _ UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	conf           *Config
}

func NewUserService(userRepository repository.UserRepository, conf *Config) *service {
	return &service{
		userRepository: userRepository,
		conf:           conf,
	}
}

func (s *service) CreateUser(ctx context.Context, user *models.User, token string) (int32, error) {
	fmt.Println(hashPassword(user.Password))
	return 1, nil
}
