package models

type Comment struct {
	ID       string
	UserID   string
	TsrID    string
	CommText string
}

type CommentAdd struct {
	UserID   string
	TsrID    string
	CommText string
}

type CommentGet struct {
	FirstName string
	LastName  string
	CommText  string
}
