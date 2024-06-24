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
}
