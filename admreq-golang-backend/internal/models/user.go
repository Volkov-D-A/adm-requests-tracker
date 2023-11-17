package models

type User struct {
	ID        string
	FirstName string
	LastName  string
	Login     string
	Password  string
	Role      string
}

type UserAuth struct {
	Login    string
	Password string
}

type UserRole struct {
	Login string
	Role  string
}

type UserCreate struct {
	FirstName string
	LastName  string
	Login     string
	Password  string
	Role      string
}
