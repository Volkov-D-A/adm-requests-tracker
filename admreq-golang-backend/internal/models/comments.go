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
	Firstname   string    `db:"firstname"`
	Lastname    string    `db:"lastname"`
	Surname     string    `db:"surname"`
	TextComment string    `db:"comm_text"`
	PostedAt    time.Time `db:"created_at"`
}
