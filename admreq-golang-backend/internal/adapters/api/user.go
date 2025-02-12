package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	RegisterUser(user *models.UserCreate, ut *models.UserToken) error
	UserAuth(creds *models.UserAuth) (*models.UserResponse, error)
	DisableUser(uuid string, ut *models.UserToken) error
	GetUsers(ut *models.UserToken) ([]models.UserResponse, error)
	AddDepartment(ad *models.AddDepartment, ut *models.UserToken) error
	GetDepartments(gd *models.GetDepartment, ut *models.UserToken) ([]models.DepartmentResponse, error)
	ChangeUserPassword(uuid, password string, ut *models.UserToken) error
	UpdateUserRight(ur *models.UserRight, ut *models.UserToken) error
}

type UserApi struct {
	tsr.UnimplementedUserServiceServer
	userService UserService
	config      *UserConfig
}

type UserConfig struct {
	Key string
}

func NewUserApi(userService UserService, config *UserConfig) *UserApi {
	return &UserApi{
		userService: userService,
		config:      config,
	}
}

func (i *UserApi) GetUsers(ctx context.Context, req *tsr.GetUsersRequest) (*tsr.GetUsersResponse, error) {
	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	res, err := i.userService.GetUsers(ut)
	switch err {
	case nil:
		break
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrUserNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error gettign userlist: %v", err)
	}

	result := make([]*tsr.GetUsersResponseUser, len(res))
	for z, x := range res {
		result[z] = &tsr.GetUsersResponseUser{
			Uuid:           x.ID,
			Firstname:      x.Firstname,
			Lastname:       x.Lastname,
			Surname:        x.Surname,
			DepartmentId:   x.DepartmentID,
			DepartmentName: x.DepartmentName,
			Login:          x.Login,
			UserRights: &tsr.Rights{
				Create:   x.Create,
				Employee: x.Employee,
				Admin:    x.Admin,
				Users:    x.Users,
				Archiv:   x.Archiv,
				Stat:     x.Stat,
			},
		}
	}

	return &tsr.GetUsersResponse{
		Users: result}, nil
}

func (i *UserApi) UserAuth(ctx context.Context, req *tsr.UserAuthRequest) (*tsr.UserAuthResponse, error) {
	creds := &models.UserAuth{
		Login:    req.Login,
		Password: req.Password,
	}
	resp, err := i.userService.UserAuth(creds)
	if err != nil {
		switch err {
		case models.ErrUnauthenticated:
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case models.ErrInvalidDataInRequest:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "unhandled error gettign authentication: %v", err)
		}
	}

	token, err := getUserToken(&models.UserToken{UserID: resp.ID, Rights: &models.UserRights{Create: resp.Create, Employee: resp.Employee, Admin: resp.Admin, Users: resp.Users, Archiv: resp.Archiv, Stat: resp.Stat}, Department: resp.DepartmentID}, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error gettign token: %v", err)
	}

	return &tsr.UserAuthResponse{
		Uuid:           resp.ID,
		Firstname:      resp.Firstname,
		Lastname:       resp.Lastname,
		Surname:        resp.Surname,
		DepartmentId:   resp.DepartmentID,
		DepartmentName: resp.DepartmentName,
		Login:          resp.Login,
		UserRights: &tsr.Rights{
			Create:   resp.Create,
			Employee: resp.Employee,
			Admin:    resp.Admin,
			Users:    resp.Users,
			Archiv:   resp.Archiv,
			Stat:     resp.Stat,
		},
		Token: token,
	}, nil
}

func (i *UserApi) RegisterUser(ctx context.Context, req *tsr.RegisterUserRequest) (*tsr.RegisterUserResponse, error) {

	usr := &models.UserCreate{
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Surname:      req.Surname,
		DepartmentID: req.DepartmentId,
		Login:        req.Login,
		Password:     req.Password,
		Rights: &models.UserRights{
			Create:   req.UserRights.Create,
			Employee: req.UserRights.Employee,
			Admin:    req.UserRights.Admin,
			Users:    req.UserRights.Users,
			Archiv:   req.UserRights.Archiv,
			Stat:     req.UserRights.Stat,
		},
	}

	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	err = i.userService.RegisterUser(usr, ut)

	switch err {
	case nil:
		return &tsr.RegisterUserResponse{}, nil
	case models.ErrUserAlreadyExists:
		return nil, status.Error(codes.AlreadyExists, err.Error())
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrInvalidDataInRequest:
		return nil, status.Error(codes.InvalidArgument, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error creating user: %v", err)
	}

}

func (i *UserApi) DisableUser(ctx context.Context, req *tsr.DisableUserRequest) (*tsr.DisableUserResponse, error) {
	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	err = i.userService.DisableUser(req.Uuid, ut)

	switch err {
	case nil:
		return &tsr.DisableUserResponse{}, nil
	case models.ErrUserNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error disabling user: %v", err)
	}

}

func (i *UserApi) AddDepartment(ctx context.Context, req *tsr.AddDepartmentRequest) (*tsr.AddDepartmentResponse, error) {
	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	ad := &models.AddDepartment{
		DepartmentName:   req.DepartmentName,
		DepartmentDoWork: req.DepartmentDowork,
	}
	err = i.userService.AddDepartment(ad, ut)
	switch err {
	case nil:
		return &tsr.AddDepartmentResponse{}, nil
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrRowAlreadyExists:
		return nil, status.Error(codes.AlreadyExists, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "error adding department: %v", err)
	}

}

func (i *UserApi) GetDepartments(ctx context.Context, req *tsr.GetDepartmentsRequest) (*tsr.GetDepartmentsResponse, error) {
	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	res, err := i.userService.GetDepartments(&models.GetDepartment{Mode: req.Mode}, ut)
	switch err {
	case nil:
		break
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrDepartmentsNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error getting list departments: %v", err)
	}

	result := make([]*tsr.GetDepartmentsResponse_Department, len(res))
	for z, x := range res {
		result[z] = &tsr.GetDepartmentsResponse_Department{
			Uuid:       x.ID,
			Department: x.DepartmentName,
			DoWork:     x.DepartmnetDoWork,
		}
	}
	return &tsr.GetDepartmentsResponse{Departments: result}, nil
}

func (i *UserApi) ChangeUserPassword(ctx context.Context, req *tsr.ChangeUserPasswordRequest) (*tsr.ChangeUserPasswordResponse, error) {
	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	err = i.userService.ChangeUserPassword(req.Uuid, req.Password, ut)
	switch err {
	case nil:
		return &tsr.ChangeUserPasswordResponse{}, nil
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrUserNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error changing password: %v", err)
	}
}

func (i *UserApi) UpdateUserRight(ctx context.Context, req *tsr.UpdateUserRightRequest) (*tsr.UpdateUserRightResponse, error) {
	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	err = i.userService.UpdateUserRight(&models.UserRight{RightName: req.Name, RightValue: req.Value, UserUUID: req.UserUuid}, ut)

	switch err {
	case nil:
		return &tsr.UpdateUserRightResponse{}, nil
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrInvalidDataInRequest:
		return nil, status.Error(codes.InvalidArgument, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled update rights error: %v", err)
	}
}
