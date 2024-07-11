package models

import (
	"database/sql"
	"time"
)

type TSR struct {
	ID             string
	UserID         string
	Text           string
	Created        time.Time
	Finished       time.Time
	EmployeeUserID string
	FinishText     string
	IsImportant    bool
	IsFinished     bool
}

type CreateTSR struct {
	UserID string
	Text   string
}

type SetEmployee struct {
	UserID string
	TSRId  string
}

type SetImportant struct {
	TSRId     string
	Important bool
}

type FinishTSR struct {
	TSRId string
}

type ApplyTSR struct {
	TSRId string
}

type ListTicketResponse struct {
	ID               string         `db:"id"`
	Text             string         `db:"req_text"`
	CreatedAt        time.Time      `db:"created_at"`
	UserFirstname    string         `db:"user_firstname"`
	UserLastname     string         `db:"user_lastname"`
	UserSurname      string         `db:"user_surname"`
	EmployeFirstname sql.NullString `db:"employe_firstname"`
	EmployeLastname  sql.NullString `db:"employe_lastname"`
	EmployeSurname   sql.NullString `db:"employe_surname"`
	Important        bool           `db:"req_important"`
	Finished         bool           `db:"req_finished"`
}

type FullTsrInfo struct {
	ID               string         `db:"id"`
	Text             string         `db:"req_text"`
	UserFirstname    string         `db:"user_firstname"`
	UserLastname     string         `db:"user_lastname"`
	UserSurname      string         `db:"user_surname"`
	UserDepartment   string         `db:"user_department"`
	EmployeFirstname sql.NullString `db:"employee_firstname"`
	EmployeLastname  sql.NullString `db:"employee_lastname"`
	EmployeSurname   sql.NullString `db:"employee_surname"`
	CreatedAt        time.Time      `db:"created_at"`
	FinishedAt       sql.NullTime   `db:"finished_at"`
	Important        bool           `db:"req_important"`
	Finished         bool           `db:"req_finished"`
	Applied          bool           `db:"req_applied"`
}
