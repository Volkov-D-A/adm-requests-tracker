package models

type UserToken struct {
	UserID     string
	Rights     *UserRights
	Department string
}
