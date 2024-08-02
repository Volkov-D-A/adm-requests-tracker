package models

type UserToken struct {
	ID         string
	Rights     *UserRights
	Department string
}
