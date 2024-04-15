package api

import (
	"context"
	"encoding/json"
	"fmt"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Create(user *models.UserCreate, ur *models.UserRole) (string, error)
	Auth(creds *models.UserAuth) (*models.UserResponse, error)
	Delete(uuid string, ur *models.UserRole) error
	GetUsers(ur *models.UserRole) ([]models.UserResponse, error)
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
	ur, _ := i.getUserRole(req.Token)
	res, err := i.userService.GetUsers(ur)
	if err != nil {
		return nil, err
	}

	result := make([]*tsr.GetUsersResponse_User, len(res))
	for z, x := range res {
		result[z] = &tsr.GetUsersResponse_User{
			Uuid:      x.ID,
			FirstName: x.FirstName,
			LastName:  x.LastName,
			Login:     x.Login,
			Role:      x.Role,
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

	token, err := i.getUserToken(&models.UserRole{ID: resp.ID, Role: resp.Role})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error gettign token: %v", err)
	}

	return &tsr.UserAuthResponse{
		Uuid:      resp.ID,
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
		Login:     resp.Login,
		Role:      resp.Role,
		Token:     token,
	}, nil
}

func (i *UserApi) RegisterUser(ctx context.Context, req *tsr.RegisterUserRequest) (*tsr.RegisterUserResponse, error) {

	usr := &models.UserCreate{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Login:     req.Login,
		Password:  req.Password,
		Role:      req.Role,
	}

	ur, err := i.getUserRole(req.Token)
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

func (i *UserApi) DeleteUser(ctx context.Context, req *tsr.DeleteUserRequest) (*tsr.Empty, error) {
	ur, err := i.getUserRole(req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	err = i.userService.Delete(req.Uuid, ur)
	if err != nil {
		switch err {
		case models.ErrUserNotExist:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "error deleting user")
		}
	}
	return &tsr.Empty{}, nil
}

func (i *UserApi) getUserRole(token string) (*models.UserRole, error) {
	var ur models.UserRole
	str, err := utils.DecryptToken(i.config.Key, token)
	if err != nil {
		return nil, fmt.Errorf("error decrypting token: %v", err)
	}

	err = json.Unmarshal([]byte(str), &ur)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling userdata: %v", err)
	}

	fmt.Println(&ur)

	return &ur, nil
}

func (i *UserApi) getUserToken(ur *models.UserRole) (string, error) {
	roleJson, err := json.Marshal(ur)
	if err != nil {
		return "", fmt.Errorf("error marshalling data: %v", err)
	}

	token, err := utils.EncryptToken(i.config.Key, string(roleJson))
	if err != nil {
		return "", fmt.Errorf("error encrypting token: %v", err)
	}

	return token, nil
}
