package repository

import (
	"context"
	"sync"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

var _ UserRepository = (*repository)(nil)

type repository struct {
	m sync.RWMutex
}

func NewUserRepository() *repository {
	return &repository{}
}

func (r *repository) Create(ctx context.Context, user *models.User) (int, error) {
	return 0, nil
}
