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

func (r *tsrStorage) TSRImportance(itsr *models.SetImportant) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_important = $1 WHERE id = $2", itsr.Important, itsr.TSRId)
	if err != nil {
		return fmt.Errorf("error updating reqtickets: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) FinishTSR(ftsr *models.FinishTSR, employee_id string) error {

	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_finished = TRUE WHERE id = $1 AND employee_user_id = $2", ftsr.TSRId, employee_id)
	if err != nil {
		return fmt.Errorf("error while finishing ticket: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) GetListTickets(mode, uuid string) ([]models.ListTicketResponse, error) {
	var query string
	switch mode {
	case "user":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, p2.firstname AS employe_firstname, p2.lastname AS employe_lastname, p2.surname AS employe_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id WHERE user_id = '%s' AND req_finished = FALSE", uuid)
	case "employee":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, p2.firstname AS employe_firstname, p2.lastname AS employe_lastname, p2.surname AS employe_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id WHERE employee_user_id = '%s') AND req_finished = FALSE", uuid)
	default:
		query = "SELECT reqtickets.id, req_text, created_at, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, p2.firstname AS employe_firstname, p2.lastname AS employe_lastname, p2.surname AS employe_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id WHERE req_finished = FALSE"
	}

	rws, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying tickets: %v", err)
	}

	tickets, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.ListTicketResponse])
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

func (r *tsrStorage) GetComments(tsrid string) ([]models.ResponseComments, error) {
	rws, err := r.db.Query(context.Background(), "SELECT comm_text, first_name, last_name, created_at FROM reqcomments LEFT JOIN requsers ON reqcomments.user_id = requsers.id WHERE reqcomments.req_id = $1 ORDER BY created_at ASC", tsrid)

	switch err {
	case nil:
		break
	case pgx.ErrNoRows:
		nullcomments := make([]models.ResponseComments, 0)
		return nullcomments, nil
	default:
		return nil, fmt.Errorf("error querying comments: %v", err)
	}

	comments, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.ResponseComments])
	if err != nil {
		return nil, fmt.Errorf("error collecting comments: %v", err)
	}

	fmt.Println(comments)
	return comments, nil
}
