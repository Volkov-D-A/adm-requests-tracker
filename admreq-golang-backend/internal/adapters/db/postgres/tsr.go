package storage

import (
	"context"
	"fmt"
	"time"

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

func (r *tsrStorage) CreateTSR(ctsr *models.CreateTSR) (string, error) {
	var uid string
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO reqtickets (user_id, req_text, target_department) VALUES ($1, $2, $3) RETURNING id", ctsr.UserID, ctsr.Text, ctsr.TargetDepartment).Scan(&uid)
	if err != nil {
		return "", fmt.Errorf("error adding ticket: %v", err)
	}
	return uid, nil
}

func (r *tsrStorage) EmployeeTSR(etsr *models.SetEmployee) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET employee_user_id = $1 WHERE id = $2", etsr.UserID, etsr.TSRId)
	if err != nil {
		return fmt.Errorf("error updating reqtickets: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) SetTimeBefore(stb *models.SetTimeBefore) error {
	ct, err := r.db.Exec(context.Background(), "UPDATE reqtickets SET finish_before = $1 AT TIME ZONE 'Asia/Yekaterinburg' WHERE id = $2", stb.FinishBefore, stb.TSRId)
	if err != nil {
		return fmt.Errorf("error updating reqtickets: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) DelEmplOrTimeBefore(del *models.DelEmplOrTimeBefore) error {
	var sql string
	switch del.DelMode {
	case "employee":
		sql = fmt.Sprintf("UPDATE reqtickets SET employee_user_id = NULL WHERE id = '%s'", del.TSRId)
	case "timebefore":
		sql = fmt.Sprintf("UPDATE reqtickets SET finish_before = NULL WHERE id = '%s'", del.TSRId)
	default:
		return models.ErrInvalidDataInRequest
	}

	ct, err := r.db.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("error deleting employee or time before: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) ImportanceTSR(itsr *models.SetImportant) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_important = $1 WHERE id = $2", itsr.Important, itsr.TSRId)
	if err != nil {
		return fmt.Errorf("error updating reqtickets: %v", err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) CheckTSROwn(user_uuid, tsr_uuid, mode string) (bool, error) {
	var req string
	switch mode {
	case "user":
		req = fmt.Sprintf("SELECT * FROM reqtickets WHERE id = '%s' AND user_id = '%s'", tsr_uuid, user_uuid)
	case "employee":
		req = fmt.Sprintf("SELECT * FROM reqtickets WHERE id = '%s' AND employee_user_id = '%s'", tsr_uuid, user_uuid)
	}
	ct, err := r.db.Pool.Exec(context.Background(), req)
	if err != nil {
		return false, fmt.Errorf("error while requesting tsr by user or employee: %v", err)
	}
	if ct.RowsAffected() == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *tsrStorage) FinishTSR(ftsr *models.FinishTSR) error {

	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_finished = TRUE, finished_at = CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg' WHERE id = $1", ftsr.TSRId)
	if err != nil {
		return fmt.Errorf("error while finishing ticket: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) ApplyTSR(atsr *models.ApplyTSR) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_applied = TRUE WHERE id = $1", atsr.TSRId)
	if err != nil {
		return fmt.Errorf("error while finishing ticket: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return models.ErrTicketNotExist
	}
	return nil
}

func (r *tsrStorage) RejectTSR(rtsr *models.RejectTSR) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE reqtickets SET req_finished = FALSE, finished_at = NULL WHERE id = $1", rtsr.TSRId)
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
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, finish_before, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE user_id = '%s' AND req_applied = FALSE ORDER BY created_at ASC", uuid)
	case "employee":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, finish_before, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE employee_user_id = '%s' AND req_finished = FALSE ORDER BY created_at ASC", uuid)
	case "archive":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, finish_before, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE req_applied = TRUE AND target_department = '%s' ORDER BY created_at ASC", dep_uuid)
	case "admin":
		query = fmt.Sprintf("SELECT reqtickets.id, req_text, created_at, req_important, req_finished, finish_before, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE req_applied = FALSE AND target_department = '%s' ORDER BY created_at ASC", dep_uuid)
	default:
		return nil, models.ErrInvalidDataInRequest
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

func (r *tsrStorage) AddTsrComment(comment *models.CommentAdd) (string, error) {
	var uuid string
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO reqcomments (req_id, user_id, comm_text) VALUES ($1, $2, $3) RETURNING id", comment.TsrID, comment.UserID, comment.TextComment).Scan(&uuid)
	if err != nil {
		return "", fmt.Errorf("error adding comment: %v", err)
	}
	return uuid, nil
}

func (r *tsrStorage) GetTsrComments(tsrid string) ([]models.ResponseComments, error) {
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
	row, err := r.db.Query(context.Background(), "SELECT reqtickets.id, req_text, p1.id AS user_id, p1.firstname AS user_firstname, p1.lastname AS user_lastname, p1.surname AS user_surname, departments.department_name AS user_department, p2.id AS employee_id, p2.firstname AS employee_firstname, p2.lastname AS employee_lastname, p2.surname AS employee_surname, created_at, finished_at, finish_before, req_important, req_finished, req_applied FROM reqtickets LEFT JOIN requsers AS p1 ON p1.id = user_id LEFT JOIN requsers AS p2 ON p2.id = employee_user_id LEFT JOIN departments ON departments.id = p1.department WHERE reqtickets.id = $1", tsrid)
	if err != nil {
		return nil, fmt.Errorf("error getting tsr data: %v", err)
	}

	tsrdata, err := pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[models.FullTsrInfo])

	switch err {
	case nil:
		return &tsrdata, nil
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

func (r *tsrStorage) GetDepartmentsList() ([]models.Department, error) {
	rws, err := r.db.Pool.Query(context.Background(), "SELECT id, department_name FROM departments")
	if err != nil {
		return nil, fmt.Errorf("error querying departmrnts: %v", err)
	}

	deps, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.Department])
	if err != nil {
		return nil, fmt.Errorf("error collecting departmrnts: %v", err)
	}

	return deps, nil
}

func (r *tsrStorage) GetStatByDepartment(req *models.StatByDepartmentReq) (*models.StatByDepartment, error) {
	result := &models.StatByDepartment{}
	err := r.db.QueryRow(context.Background(), "SELECT count(reqtickets.id) FROM reqtickets LEFT JOIN requsers ON reqtickets.user_id = requsers.id LEFT JOIN departments ON requsers.department = departments.id WHERE target_department = $1 AND departments.id = $2 AND req_finished = FALSE", req.TargetDepartmentUUID, req.SourceDepartmentUUID).Scan(&result.TsrInWork)
	if err != nil {
		return nil, fmt.Errorf("error querying count in work tsr: %v", err)
	}
	err = r.db.QueryRow(context.Background(), "SELECT count(reqtickets.id) FROM reqtickets LEFT JOIN requsers ON reqtickets.user_id = requsers.id LEFT JOIN departments ON requsers.department = departments.id WHERE target_department = $1 AND departments.id = $2 AND req_finished = TRUE AND req_applied = FALSE", req.TargetDepartmentUUID, req.SourceDepartmentUUID).Scan(&result.TsrFinished)
	if err != nil {
		return nil, fmt.Errorf("error querying count finished tsr: %v", err)
	}
	err = r.db.QueryRow(context.Background(), "SELECT count(reqtickets.id) FROM reqtickets LEFT JOIN requsers ON reqtickets.user_id = requsers.id LEFT JOIN departments ON requsers.department = departments.id WHERE target_department = $1 AND departments.id = $2 AND req_applied = TRUE", req.TargetDepartmentUUID, req.SourceDepartmentUUID).Scan(&result.TsrApplyed)
	if err != nil {
		return nil, fmt.Errorf("error querying count applyed tsr: %v", err)
	}
	return result, nil
}

func (r *tsrStorage) GetEmployeeList(target_dep string) ([]models.Employee, error) {
	rws, err := r.db.Pool.Query(context.Background(), "SELECT requsers.id, firstname, lastname, surname FROM requsers LEFT JOIN rights ON requsers.user_rights = rights.id WHERE employee_tsr = TRUE AND department = $1", target_dep)
	if err != nil {
		return nil, fmt.Errorf("error querying employees by department: %v", err)
	}

	result, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.Employee])
	if err != nil {
		return nil, fmt.Errorf("error collecting employees by department data: %v", err)
	}
	return result, nil
}

func (r *tsrStorage) GetStatByEmployee(req *models.StatByEmployeeReq) (*models.StatByEmployee, error) {
	result := &models.StatByEmployee{}
	err := r.db.QueryRow(context.Background(), "SELECT count(reqtickets.id) FROM reqtickets WHERE employee_user_id = $1 AND req_finished = FALSE", req.EmplotyeeUUID).Scan(&result.TsrInWork)
	if err != nil {
		return nil, fmt.Errorf("error querying count in work tsr by epmloyee: %v", err)
	}
	err = r.db.QueryRow(context.Background(), "SELECT count(reqtickets.id) FROM reqtickets WHERE employee_user_id = $1 AND req_finished = TRUE AND req_applied = FALSE", req.EmplotyeeUUID).Scan(&result.TsrFinished)
	if err != nil {
		return nil, fmt.Errorf("error querying count finished tsr by epmloyee: %v", err)
	}
	err = r.db.QueryRow(context.Background(), "SELECT count(reqtickets.id) FROM reqtickets WHERE employee_user_id = $1 AND req_applied = TRUE", req.EmplotyeeUUID).Scan(&result.TsrApplyed)
	if err != nil {
		return nil, fmt.Errorf("error querying count applied tsr by epmloyee: %v", err)
	}
	return result, nil
}

func (r *tsrStorage) SetReadTicketDate(rtd *models.ReadTiketsDate) error {
	_, err := r.db.Exec(context.Background(), "INSERT INTO readticket (req_id, user_id, lastread) VALUES ($1, $2, CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg') ON CONFLICT (req_id, user_id) DO UPDATE SET lastread = CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg'", rtd.TSRId, rtd.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *tsrStorage) CheckUnreadComments(uc *models.UnreadComments) bool {
	rws := r.db.QueryRow(context.Background(), "SELECT created_at FROM reqcomments WHERE req_id = $1 ORDER BY created_at DESC LIMIT 1", uc.TSRId)

	var lastComment time.Time

	err := rws.Scan(&lastComment)
	if err == pgx.ErrNoRows {
		return false
	}

	var lastEnterTicket time.Time

	rws = r.db.QueryRow(context.Background(), "SELECT lastread FROM readticket WHERE req_id = $1 AND user_id = $2", uc.TSRId, uc.UserID)
	err = rws.Scan(&lastEnterTicket)
	if err == pgx.ErrNoRows {
		return true
	}

	return lastComment.Compare(lastEnterTicket) == 1
}
