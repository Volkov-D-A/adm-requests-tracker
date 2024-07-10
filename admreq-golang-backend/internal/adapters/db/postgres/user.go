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
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO requsers (firstname, lastname, surname, department, user_role, user_login, user_pass) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", user.Firstname, user.Lastname, user.Surname, user.Department, user.Role, user.Login, utils.HashPassword(user.Password)).Scan(&uuid)
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

func (r *userStorage) Auth(user *models.UserAuth) (*models.UserResponse, error) {
	var resp models.UserResponse

	err := r.db.Pool.QueryRow(context.Background(), "SELECT id, firstname, lastname, surname, department, user_login, user_role FROM requsers WHERE user_login = $1 AND user_pass = $2 AND user_disabled = FALSE", user.Login, utils.HashPassword(user.Password)).Scan(&resp.ID, &resp.Firstname, &resp.Lastname, &resp.Surname, &resp.Department, &resp.Login, &resp.Role)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, models.ErrUnauthenticated
		default:
			return nil, fmt.Errorf("error while querying: %v", err)
		}
	}
	return &resp, nil
}

func (r *userStorage) Delete(uuid string) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE requsers SET user_disabled = TRUE WHERE id = $1", uuid)
	if err != nil {
		return fmt.Errorf("error deleting user %v: %v", uuid, err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrUserNotExist
	}
	return nil
}

func (r *userStorage) GetUsers() ([]models.UserResponse, error) {

	rws, err := r.db.Pool.Query(context.Background(), "SELECT id, firstname, lastname, surname, department, user_login, user_role FROM requsers WHERE user_disabled = FALSE")

	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.UserResponse])
	if err != nil {
		return nil, err
	}

	return users, nil
}
