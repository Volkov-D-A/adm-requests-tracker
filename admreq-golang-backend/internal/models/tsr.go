package models

import "time"

type TSR struct {
	ID             string
	UserID         string
	Text           string
	Created        time.Time
	Finished       time.Time
	EmployeeUserID string
	Comment        string
	IsImportant    bool
	IsFinished     bool
}
