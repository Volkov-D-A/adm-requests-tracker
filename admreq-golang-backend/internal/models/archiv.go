package models

import (
	"database/sql"
	"time"
)

type ArchivTicketResponse struct {
	ID                string       `db:"id"`
	Text              string       `db:"req_text"`
	CreatedAt         time.Time    `db:"created_at"`
	FinishBefore      sql.NullTime `db:"finish_before"`
	FinishedAt        time.Time    `db:"finished_at"`
	UserFirstname     string       `db:"user_firstname"`
	UserLastname      string       `db:"user_lastname"`
	UserSurname       string       `db:"user_surname"`
	UserDepartment    string       `db:"user_department"`
	EmployeeFirstname string       `db:"employee_firstname"`
	EmployeeLastname  string       `db:"employee_lastname"`
	EmployeeSurname   string       `db:"employee_surname"`
	Important         bool         `db:"req_important"`
}

type ArchivResponse struct {
	Tickets []ArchivTicketResponse
	Pages   int32
}
