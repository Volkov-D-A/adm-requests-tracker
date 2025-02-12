package service

import (
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type UserStorage interface {
	RegisterUser(user *models.UserCreate) (string, error)
	UserAuth(user *models.UserAuth) (*models.UserResponse, error)
	DisableUser(uuid string) error
	GetUsers() ([]models.UserResponse, error)
	AddDepartment(ad *models.AddDepartment) (string, error)
	GetDepartments(gd *models.GetDepartment) ([]models.DepartmentResponse, error)
	ChangeUserPassword(uuid, password string) error
	RecordAction(act *models.ActionADD) error
	UpdateUserRight(ur *models.UserRight) error
}

type userService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *userService {
	return &userService{
		userStorage: userStorage,
	}
}

func (s *userService) RegisterUser(user *models.UserCreate, ut *models.UserToken) error {
	if !ut.Rights.Users {
		return models.ErrUnauthorized
	}

	if user.Firstname == "" || user.Lastname == "" || user.Login == "" || user.Password == "" || user.DepartmentID == "" {
		return models.ErrInvalidDataInRequest
	}

	uuid, err := s.userStorage.RegisterUser(user)
	if err != nil {
		return err
	}
	s.userStorage.RecordAction(&models.ActionADD{SubjectID: ut.UserID, ObjectID: uuid, Action: "UserAdd"})
	return nil
}

func (s *userService) UserAuth(user *models.UserAuth) (*models.UserResponse, error) {
	if user.Login == "" || user.Password == "" {
		return nil, models.ErrInvalidDataInRequest
	}

	resp, err := s.userStorage.UserAuth(user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) DisableUser(uuid string, ut *models.UserToken) error {
	if !ut.Rights.Users {
		return models.ErrUnauthorized
	}
	err := s.userStorage.DisableUser(uuid)
	if err != nil {
		return err
	}
	s.userStorage.RecordAction(&models.ActionADD{SubjectID: ut.UserID, ObjectID: uuid, Action: "UserDisable"})
	return nil
}

func (s *userService) GetUsers(ut *models.UserToken) ([]models.UserResponse, error) {
	if !ut.Rights.Users && !ut.Rights.Admin {
		return nil, models.ErrUnauthorized
	}

	resp, err := s.userStorage.GetUsers()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) AddDepartment(ad *models.AddDepartment, ut *models.UserToken) error {
	if !ut.Rights.Users {
		return models.ErrUnauthorized
	}
	uuid, err := s.userStorage.AddDepartment(ad)
	if err != nil {
		return err
	}
	s.userStorage.RecordAction(&models.ActionADD{SubjectID: ut.UserID, ObjectID: uuid, Action: "DepartmentAdd"})
	return nil
}

func (s *userService) GetDepartments(gd *models.GetDepartment, ut *models.UserToken) ([]models.DepartmentResponse, error) {
	if !ut.Rights.Users && !ut.Rights.Create {
		return nil, models.ErrUnauthorized
	}
	res, err := s.userStorage.GetDepartments(gd)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userService) ChangeUserPassword(uuid, password string, ut *models.UserToken) error {
	if !ut.Rights.Users {
		return models.ErrUnauthorized
	}

	err := s.userStorage.ChangeUserPassword(uuid, password)
	if err != nil {
		return err
	}
	s.userStorage.RecordAction(&models.ActionADD{SubjectID: ut.UserID, ObjectID: uuid, Action: "ChangePassword"})
	return nil

}

func (s *userService) UpdateUserRight(ur *models.UserRight, ut *models.UserToken) error {
	if !ut.Rights.Users {
		return models.ErrUnauthorized
	}

	err := s.userStorage.UpdateUserRight(ur)
	if err != nil {
		return err
	}
	return nil
}
