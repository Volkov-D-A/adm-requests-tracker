package storage

import (
	"context"
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	pg "github.com/volkov-d-a/adm-requests-tracker/pkg/PG"
)

type tsrStorage struct {
	db *pg.PG
}

func NewTsrStorage(db *pg.PG) *tsrStorage {
	return &tsrStorage{db: db}
}

func (r *tsrStorage) Create(ctsr *models.CreateTSR) (string, error) {
	var uid string
	ct, err := r.db.Pool.Exec(context.Background(), "SELECT * FROM requsers WHERE id=$1", ctsr.UserID)
	if err != nil {
		return "", fmt.Errorf("error checking user: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return "", models.ErrUserNotExist
	}

	err = r.db.Pool.QueryRow(context.Background(), "INSERT INTO reqtickets (user_id, req_text) VALUES ($1, $2) RETURNING id", ctsr.UserID, ctsr.Text).Scan(&uid)
	if err != nil {
		return "", fmt.Errorf("error adding ticket: %v", err)
	}
	return uid, nil
}
