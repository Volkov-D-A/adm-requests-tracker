package models

import "time"

type Comment struct {
	ID          string
	UserID      string
	TsrID       string
	TextComment string
}

type CommentAdd struct {
	UserID      string
	TsrID       string
	TextComment string
}

type ResponseComments struct {
	FirstName   string    `db:"comm_text"`
	LastName    string    `db:"first_name"`
	TextComment string    `db:"last_name"`
	PostedAt    time.Time `db:"created_at"`
}
