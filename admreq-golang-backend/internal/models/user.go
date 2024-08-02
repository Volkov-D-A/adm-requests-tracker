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
	Create         bool   `db:"create_tsr"`
	Employee       bool   `db:"employee_tsr"`
	Admin          bool   `db:"admin_tsr"`
	Users          bool   `db:"admin_users"`
	Archiv         bool   `db:"archiv_tsr"`
	Stat           bool   `db:"stat_tsr"`
}

type UserCreate struct {
	Firstname    string
	Lastname     string
	Surname      string
	DepartmentID string
	Login        string
	Password     string
	Rights       *UserRights
}

type UserRights struct {
	Create   bool `db:"create_tsr"`
	Employee bool `db:"employee_tsr"`
	Admin    bool `db:"admin_tsr"`
	Users    bool `db:"admin_users"`
	Archiv   bool `db:"archiv_tsr"`
	Stat     bool `db:"stat_tsr"`
}

type AddDepartment struct {
	DepartmentName   string
	DepartmentDoWork bool
}

type GetDepartment struct {
	Mode string
}

type DepartmentResponse struct {
	ID               string `db:"id"`
	DepartmentName   string `db:"department_name"`
	DepartmnetDoWork bool   `db:"department_dowork"`
}
