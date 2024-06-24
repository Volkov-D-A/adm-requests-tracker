package models

type User struct {
	ID         string
	Firstname  string
	Lastname   string
	Surname    string
	Department string
	Login      string
	Password   string
	Role       string //admin, employee, user
}

type UserAuth struct {
	Login    string
	Password string
}

type UserResponse struct {
	ID         string `db:"id"`
	Firstname  string `db:"firstname"`
	Lastname   string `db:"lastname"`
	Surname    string `db:"surname"`
	Department string `db:"department"`
	Login      string `db:"user_login"`
	Role       string `db:"user_role"`
}

type UserCreate struct {
	Firstname  string
	Lastname   string
	Surname    string
	Department string
	Login      string
	Password   string
	Role       string
}
