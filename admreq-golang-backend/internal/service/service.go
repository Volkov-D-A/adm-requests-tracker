package service

import (
	"context"

	models "github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User, token string) (int32, error)
}

type Config struct {
	Key string
}
