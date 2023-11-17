package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	pg "github.com/volkov-d-a/adm-requests-tracker/pkg/PG"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
)

type userStorage struct {
	db *pg.PG
}

func NewUserStorage(db *pg.PG) *userStorage {
	return &userStorage{db: db}
}

func (r *userStorage) Create(user *models.UserCreate) (string, error) {
	var uuid string
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO requsers (first_name, last_name, user_role, user_login, user_pass) VALUES ($1, $2, $3, $4, $5) RETURNING id", user.FirstName, user.LastName, user.Role, user.Login, utils.HashPassword(user.Password)).Scan(&uuid)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return "", models.ErrUserAlreadyExists
			}
		}
		return "", err
	}
	return uuid, nil
}

func (r *userStorage) Auth(user *models.UserAuth) (*models.UserRole, error) {
	var role models.UserRole

	err := r.db.Pool.QueryRow(context.Background(), "SELECT user_login, user_role FROM requsers WHERE user_login = $1 AND user_pass = $2", user.Login, utils.HashPassword(user.Password)).Scan(&role.Login, &role.Role)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, models.ErrUnauthenticated
		default:
			return nil, fmt.Errorf("error while querying: %v", err)
		}
	}
	return &role, nil
}
