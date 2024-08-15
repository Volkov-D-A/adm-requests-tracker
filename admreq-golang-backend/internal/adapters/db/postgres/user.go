package storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	pg "github.com/volkov-d-a/adm-requests-tracker/pkg/PG"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
)

type userStorage struct {
	db *pg.PG
}

func NewUserStorage(db *pg.PG) *userStorage {
	return &userStorage{db: db}
}

func (r *userStorage) RegisterUser(user *models.UserCreate) (string, error) {
	var rights_uuid string
	var user_uuid string
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO rights (create_tsr, employee_tsr, admin_tsr, admin_users, archiv_tsr, stat_tsr) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", user.Rights.Create, user.Rights.Employee, user.Rights.Admin, user.Rights.Users, user.Rights.Archiv, user.Rights.Stat).Scan(&rights_uuid)
	if err != nil {
		return "", fmt.Errorf("error while adding user rights: %v", err)
	}
	err = r.db.Pool.QueryRow(context.Background(), "INSERT INTO requsers (firstname, lastname, surname, department, user_rights, user_login, user_pass) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", user.Firstname, user.Lastname, user.Surname, user.DepartmentID, rights_uuid, user.Login, utils.HashPassword(user.Password)).Scan(&user_uuid)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return "", models.ErrUserAlreadyExists
			}
		}
		_, err = r.db.Pool.Exec(context.Background(), "DELETE FROM rights WHERE id = $1", rights_uuid)
		if err != nil {
			return "", fmt.Errorf("error deleting rights: %v", err)
		}
		return "", fmt.Errorf("error while adding user: %v", err)
	}
	return user_uuid, nil
}

func (r *userStorage) UserAuth(user *models.UserAuth) (*models.UserResponse, error) {
	var resp models.UserResponse

	rows, err := r.db.Pool.Query(context.Background(), "SELECT requsers.id, firstname, lastname, surname, departments.id AS department_id, departments.department_name AS department_name, user_login, create_tsr, employee_tsr, admin_tsr, admin_users, archiv_tsr, stat_tsr FROM requsers LEFT JOIN departments ON departments.id = requsers.department LEFT JOIN rights ON rights.id = requsers.user_rights WHERE user_login = $1 AND user_pass = $2 AND user_disabled = FALSE", user.Login, utils.HashPassword(user.Password))
	switch err {
	case nil:
		break
	case pgx.ErrNoRows:
		return nil, models.ErrUnauthenticated
	default:
		return nil, fmt.Errorf("error while auth select: %v", err)
	}

	resp, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.UserResponse])
	if err != nil {
		return nil, fmt.Errorf("error collecting data auth: %v", err)
	}

	return &resp, nil

}

func (r *userStorage) DisableUser(uuid string) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE requsers SET user_disabled = TRUE WHERE id = $1", uuid)
	if err != nil {
		return fmt.Errorf("error querying deleting user %v: %v", uuid, err)
	}
	if ct.RowsAffected() == 0 {
		return models.ErrUserNotExist
	}
	return nil
}

func (r *userStorage) GetUsers() ([]models.UserResponse, error) {

	rws, err := r.db.Pool.Query(context.Background(), "SELECT requsers.id, firstname, lastname, surname, departments.id AS department_id, departments.department_name AS department_name, user_login, create_tsr, employee_tsr, admin_tsr, admin_users, archiv_tsr, stat_tsr FROM requsers LEFT JOIN departments ON departments.id = requsers.department LEFT JOIN rights ON rights.id = requsers.user_rights WHERE user_disabled = FALSE")

	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}

	if rws.CommandTag().RowsAffected() == 0 {
		return nil, models.ErrUserNotExist
	}

	users, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.UserResponse])
	if err != nil {
		return nil, fmt.Errorf("error collecting users data: %v", err)
	}

	return users, nil
}

func (r *userStorage) AddDepartment(ad *models.AddDepartment) (string, error) {
	var uuid string
	err := r.db.Pool.QueryRow(context.Background(), "INSERT INTO departments (department_name, department_dowork) VALUES ($1, $2) RETURNING id", ad.DepartmentName, ad.DepartmentDoWork).Scan(&uuid)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return "", models.ErrRowAlreadyExists
			}
		}
		return "", fmt.Errorf("unhandled error while adding department: %v", err)
	}
	return uuid, nil
}

func (r *userStorage) GetDepartments(gd *models.GetDepartment) ([]models.DepartmentResponse, error) {
	var request string
	switch gd.Mode {
	case "user":
		request = "SELECT id, department_name, department_dowork FROM departments WHERE department_active = TRUE AND department_dowork = TRUE"
	case "admin":
		request = "SELECT id, department_name, department_dowork FROM departments WHERE department_active = TRUE"
	}

	rws, err := r.db.Pool.Query(context.Background(), request)
	if err != nil {
		return nil, err
	}

	if rws.CommandTag().RowsAffected() == 0 {
		return nil, models.ErrDepartmentsNotExist
	}

	departments, err := pgx.CollectRows(rws, pgx.RowToStructByName[models.DepartmentResponse])
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *userStorage) ChangeUserPassword(uuid, password string) error {
	ct, err := r.db.Pool.Exec(context.Background(), "UPDATE requsers SET user_pass = $1 WHERE id = $2", utils.HashPassword(password), uuid)
	if ct.RowsAffected() == 0 {
		return models.ErrUserNotExist
	}
	if err != nil {
		return fmt.Errorf("error updating password in db: %v", err)
	}
	return nil
}

func (r *userStorage) RecordAction(act *models.ActionADD) error {
	_, err := r.db.Pool.Exec(context.Background(), "INSERT INTO actions (action_subject, action_object, action_string, action_info) VALUES ($1, $2, $3, $4)", act.SubjectID, act.ObjectID, act.Action, act.Info)
	if err != nil {
		return err
	}
	return nil
}

func (r *userStorage) UpdateUserRight(ur *models.UserRight) error {
	var column string
	right := strconv.FormatBool(ur.RightValue)
	switch ur.RightName {
	case "create":
		column = "create_tsr"
	case "employee":
		column = "employee_tsr"
	case "admin":
		column = "admin_tsr"
	case "users":
		column = "admin_users"
	case "archiv":
		column = "archiv_tsr"
	case "stat":
		column = "stat_tsr"
	default:
		return models.ErrInvalidDataInRequest
	}

	query := fmt.Sprintf("UPDATE rights SET %s = %s FROM (SELECT user_rights FROM requsers WHERE id = '%s') AS subquery WHERE rights.id = subquery.user_rights", column, right, ur.UserUUID)
	ct, err := r.db.Pool.Exec(context.Background(), query)
	if ct.RowsAffected() == 0 {
		return models.ErrInvalidDataInRequest
	}
	if err != nil {
		return fmt.Errorf("unhandled error while updating rights: %v", err)
	}
	return nil
}
