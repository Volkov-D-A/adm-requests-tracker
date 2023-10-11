package service

import (
	"context"

	models "github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type UserService interface {
	Create(ctx context.Context, user *models.User, token string) (int32, error)
}
