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
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO reqtickets (user_id, req_text, target_department) VALUES ($1, $2, $3) RETURNING id", ctsr.UserID, ctsr.Text, ctsr.TargetDepartment).Scan(&uid)
	if err != nil {
		return "", fmt.Errorf("error adding ticket: %v", err)
	}
	return uid, nil
}

func (r *tsrStorage) TSREmployee(etsr *models.SetEmployee) error {
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

	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_finished = TRUE, finished_at = CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg' WHERE id = $1 AND employee_user_id = $2", ftsr.TSRId, employee_id)
	if err != nil {
		return fmt.Errorf("error while finishing ticket: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) ApplyTSR(atsr *models.ApplyTSR, user_id string) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_applied = TRUE WHERE id = $1 AND user_id = $2", atsr.TSRId, user_id)
	if err != nil {
		return fmt.Errorf("error while finishing ticket: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) GetListTickets(mode, uuid, dep_uuid string) ([]models.ListTicketResponse, error) {
	var query string
	switch mode {
	case "user":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE user_id = '%s' AND req_applied = FALSE", uuid)
	case "employee":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE employee_user_id = '%s' AND req_finished = FALSE", uuid)
	case "archive":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE req_applied = TRUE AND target_department = '%s'", dep_uuid)
	default:
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE req_applied = FALSE AND target_department = '%s'", dep_uuid)
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

func (r *tsrStorage) AddComment(comment *models.CommentAdd) (string, error) {
	var uuid string
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO reqcomments (req_id, user_id, comm_text) VALUES ($1, $2, $3) RETURNING id", comment.TsrID, comment.UserID, comment.TextComment).Scan(&uuid)
	if err != nil {
		return "", fmt.Errorf("error adding comment: %v", err)
	}
	return uuid, nil
}

func (r *tsrStorage) GetComments(tsrid string) ([]models.ResponseComments, error) {
	rws, err := r.db.Query(context.Background(), "SELECT reqcomments.id, comm_text, firstname, lastname, surname, created_at FROM reqcomments LEFT JOIN requsers ON reqcomments.user_id = requsers.id WHERE reqcomments.req_id = $1 ORDER BY created_at ASC", tsrid)

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

	return comments, nil
}

func (r *tsrStorage) GetFullTsrInfo(tsrid string) (*models.FullTsrInfo, error) {
	row, err := r.db.Query(context.Background(), "SELECT reqtickets.id, req_text, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname, created_at, finished_at, req_important, req_finished, req_applied FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE reqtickets.id = $1", tsrid)
	if err != nil {
		return nil, fmt.Errorf("error getting tsr data: %v", err)
	}

	tsrdaata, err := pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[models.FullTsrInfo])

	switch err {
	case nil:
		return &tsrdaata, nil
	case pgx.ErrNoRows:
		return nil, models.ErrTicketNotExist
	default:
		return nil, fmt.Errorf("error getting tsr data: %v", err)
	}
}

func (r *tsrStorage) RecordAction(act *models.ActionADD) error {
	_, err := r.db.Pool.Exec(context.Background(), "INSERT INTO actions (action_subject, action_object, action_string, action_info) VALUES ($1, $2, $3, $4)", act.SubjectID, act.ObjectID, act.Action, act.Info)
	if err != nil {
		return err
	}
	return nil
}
