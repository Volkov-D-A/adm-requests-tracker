package repository

import (
	"context"
	"sync"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type TSRRepository struct {
	m sync.RWMutex
}

func New() *TSRRepository {
	return &TSRRepository{}
}

func (r *TSRRepository) Create(ctx context.Context, user *models.User) (int, error) {
	return 0, nil
}
