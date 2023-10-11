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
	Key            string
}

func NewUserService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Create(ctx context.Context, user *models.User, token string) (int32, error) {
	fmt.Println(encryptToken(s.Key, token))

	return 0, nil
}
