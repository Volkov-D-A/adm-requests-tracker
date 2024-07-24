package service

import (
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type UserStorage interface {
	Create(user *models.UserCreate) (string, error)
	Auth(user *models.UserAuth) (*models.UserResponse, error)
	Delete(uuid string) error
	GetUsers() ([]models.UserResponse, error)
	AddDepartment(ad *models.AddDepartment) error
	GetDepartments(gd *models.GetDepartment) ([]models.DepartmentResponse, error)
	ChangeUserPassword(uuid, password string) error
	RecordAction(act *models.ActionADD) error
}

type userService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *userService {
	return &userService{
		userStorage: userStorage,
	}
}

func (s *userService) Create(user *models.UserCreate, ut *models.UserToken) (string, error) {
	if ut.Role != "admin" {
		return "", models.ErrUnauthorized
	}
	uuid, err := s.userStorage.Create(user)
	if err != nil {
		s.userStorage.RecordAction(&models.ActionADD{SubjectID: ut.ID, ObjectID: "", Action: "UserAdd", Result: false})
		return "", err
	}
	s.userStorage.RecordAction(&models.ActionADD{SubjectID: ut.ID, ObjectID: uuid, Action: "UserAdd", Result: true})
	return uuid, nil
}

func (s *userService) Auth(user *models.UserAuth) (*models.UserResponse, error) {
	resp, err := s.userStorage.Auth(user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) Delete(uuid string, ut *models.UserToken) error {
	if ut.Role != "admin" {
		return models.ErrUnauthorized
	}
	err := s.userStorage.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUsers(ut *models.UserToken) ([]models.UserResponse, error) {
	if ut.Role != "admin" {
		return nil, models.ErrUnauthorized
	}

	resp, _ := s.userStorage.GetUsers()
	return resp, nil
}

func (s *userService) AddDepartment(ad *models.AddDepartment, ut *models.UserToken) error {
	if ut.Role != "admin" {
		return models.ErrUnauthorized
	}
	err := s.userStorage.AddDepartment(ad)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetDepartments(gd *models.GetDepartment, ut *models.UserToken) ([]models.DepartmentResponse, error) {
	if ut.Role != "admin" && gd.Mode == "admin" {
		return nil, models.ErrUnauthorized
	}
	res, err := s.userStorage.GetDepartments(gd)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userService) ChangeUserPassword(uuid, password string, ut *models.UserToken) error {
	if ut.Role != "admin" {
		return models.ErrUnauthorized
	}

	err := s.userStorage.ChangeUserPassword(uuid, password)
	switch err {
	case nil:
		return nil
	case models.ErrUserNotExist:
		return err
	default:
		return fmt.Errorf("error changing password: %v", err)
	}
}
