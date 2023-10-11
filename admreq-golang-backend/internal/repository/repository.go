package repository

import (
	"context"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (int, error)
}
