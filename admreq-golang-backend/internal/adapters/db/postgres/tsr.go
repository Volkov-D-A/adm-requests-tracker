package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func (r *tsrStorage) TSREmployee(etsr *models.SetEmployee) error {
	var role string

	err := r.db.QueryRow(context.Background(), "SELECT (user_role) FROM requsers WHERE id = $1", etsr.UserID).Scan(&role)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return models.ErrUserNotExist
		default:
			return fmt.Errorf("error checking user role: %v", err)
		}
	}

	if role == "user" {
		return models.ErrUserNotEmployee
	}

	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET employee_user_id = $1 WHERE id = $2", etsr.UserID, etsr.TSRId)
	if err != nil {
		return fmt.Errorf("error updating reqtickets: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) FinishTSR(ftsr *models.FinishTSR, employee_id string) error {

	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_finished = TRUE, req_text = $1 WHERE id = $2 AND employee_user_id = $3", ftsr.FinisText, ftsr.TSRId, employee_id)
	if err != nil {
		return fmt.Errorf("error while finishing ticket: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) GetTickets(mode, uuid string) ([]models.TicketResponse, error) {
	var query string
	switch mode {
	case "user":
		query = fmt.Sprintf("SELECT id, user_id, req_text, employee_user_id, finished_comment FROM reqtickets WHERE user_id = '%s' AND req_finished = FALSE", uuid)
	case "employee":
		query = fmt.Sprintf("SELECT id, user_id, req_text, employee_user_id, finished_comment FROM reqtickets WHERE (user_id = '%s' OR employee_user_id = '%s') AND req_finished = FALSE", uuid, uuid)
	default:
		query = "SELECT id, user_id, req_text, employee_user_id, finished_comment FROM reqtickets WHERE req_finished = FALSE"
	}

	rws, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying tickets: %v", err)
	}

	tickets, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.TicketResponse])
	if err != nil {
		return nil, fmt.Errorf("error collecting tickets: %v", err)
	}

	return tickets, nil
}

func (r *tsrStorage) AddComment(comment *models.CommentAdd) error {
	ct, err := r.db.Pool.Exec(context.Background(), "INSERT INTO reqcomments (req_id, user_id, comm_text) VALUES ($1, $2, $3)", comment.TsrID, comment.UserID, comment.TextComment)
	if err != nil {
		return fmt.Errorf("error adding comment: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return fmt.Errorf("comments are not inserted")
	}
	return nil
}
