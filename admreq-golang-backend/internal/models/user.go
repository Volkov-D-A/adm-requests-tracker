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
	ID             string `db:"id"`
	Firstname      string `db:"firstname"`
	Lastname       string `db:"lastname"`
	Surname        string `db:"surname"`
	DepartmentID   string `db:"department_id"`
	DepartmentName string `db:"department_name"`
	Login          string `db:"user_login"`
	Role           string `db:"user_role"`
}

type UserCreate struct {
	Firstname    string
	Lastname     string
	Surname      string
	DepartmentID string
	Login        string
	Password     string
	Role         string
}

type AddDepartment struct {
	DepartmentName string
}

type GetDepartment struct {
	ID             string `db:"id"`
	DepartmentName string `db:"department_name"`
}
