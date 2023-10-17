package repository

import (
	"context"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type Config struct {
	login    string
	password string
	host     string
	port     string
	database string
}

type TSRRepository struct {
	conf *Config
}

func New() *TSRRepository {
	return &TSRRepository{}
}

func (r *TSRRepository) Create(ctx context.Context, user *models.User) (int, error) {
	return 0, nil
}
