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

type FinishTSR struct {
	TSRId     string
	FinisText string
}

type TicketResponse struct {
	ID             string         `db:"id"`
	UserID         string         `db:"user_id"`
	EmployeeUserID sql.NullString `db:"employee_user_id"`
	Text           string         `db:"req_text"`
	FinishText     sql.NullString `db:"finished_comment"`
}
