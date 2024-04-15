package models

type User struct {
	ID        string
	FirstName string
	LastName  string
	Login     string
	Password  string
	Role      string //admin, employee, user
}

type UserAuth struct {
	Login    string
	Password string
}

type UserResponse struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Login     string `db:"user_login"`
	Role      string `db:"user_role"`
}

type UserRole struct {
	ID   string
	Role string
}

type UserCreate struct {
	FirstName string
	LastName  string
	Login     string
	Password  string
	Role      string
}
