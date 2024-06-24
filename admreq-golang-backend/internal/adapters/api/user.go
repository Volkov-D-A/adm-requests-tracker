package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Create(user *models.UserCreate, ut *models.UserToken) (string, error)
	Auth(creds *models.UserAuth) (*models.UserResponse, error)
	Delete(uuid string, ut *models.UserToken) error
	GetUsers(ut *models.UserToken) ([]models.UserResponse, error)
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
	if err != nil {
		switch err {
		case models.ErrUnauthorized:
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "Error gettign userlist: %v", err)
		}
	}

	result := make([]*tsr.GetUsersResponse_User, len(res))
	for z, x := range res {
		result[z] = &tsr.GetUsersResponse_User{
			Uuid:       x.ID,
			Firstname:  x.Firstname,
			Lastname:   x.Lastname,
			Surname:    x.Surname,
			Department: x.Department,
			Login:      x.Login,
			Role:       x.Role,
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
	resp, err := i.userService.Auth(creds)
	if err != nil {
		switch err {
		case models.ErrUnauthenticated:
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "Error gettign authorization: %v", err)
		}
	}

	token, err := getUserToken(&models.UserToken{ID: resp.ID, Role: resp.Role}, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error gettign token: %v", err)
	}

	return &tsr.UserAuthResponse{
		Uuid:       resp.ID,
		Firstname:  resp.Firstname,
		Lastname:   resp.Lastname,
		Surname:    resp.Surname,
		Department: resp.Department,
		Login:      resp.Login,
		Role:       resp.Role,
		Token:      token,
	}, nil
}

func (i *UserApi) RegisterUser(ctx context.Context, req *tsr.RegisterUserRequest) (*tsr.RegisterUserResponse, error) {

	usr := &models.UserCreate{
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		Surname:    req.Surname,
		Department: req.Department,
		Login:      req.Login,
		Password:   req.Password,
		Role:       req.Role,
	}

	ur, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	uuid, err := i.userService.Create(usr, ur)
	if err != nil {
		switch err {
		case models.ErrUserAlreadyExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case models.ErrUnauthorized:
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "error creating user: %v", err)
		}

	}

	return &tsr.RegisterUserResponse{
		Uuid: uuid,
	}, nil
}

func (i *UserApi) DeleteUser(ctx context.Context, req *tsr.DeleteUserRequest) (*tsr.DeleteUserResponse, error) {
	ut, err := getTokenData(req.Token, i.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	err = i.userService.Delete(req.Uuid, ut)
	if err != nil {
		switch err {
		case models.ErrUserNotExist:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "error deleting user")
		}
	}
	return &tsr.DeleteUserResponse{}, nil
}
