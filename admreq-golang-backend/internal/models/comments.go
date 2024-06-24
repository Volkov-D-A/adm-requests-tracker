package models

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
	FirstName   string
	LastName    string
	TextComment string
}
