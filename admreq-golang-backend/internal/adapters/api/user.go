package api

import (
	"context"
	"encoding/json"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	CreateUser(user *models.UserCreate, ur *models.UserRole) (string, error)
	Auth(creds *models.UserAuth) (*models.UserRole, error)
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

func (i *UserApi) UserAuth(ctx context.Context, req *tsr.UserAuthRequest) (*tsr.UserAuthResponse, error) {
	creds := &models.UserAuth{
		Login:    req.Login,
		Password: req.Password,
	}
	ur, err := i.userService.Auth(creds)
	if err != nil {
		switch err {
		case models.ErrUnauthenticated:
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "Error gettign authorization: %v", err)
		}
	}

	roleJson, err := json.Marshal(ur)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error marshalling data: %v", err)
	}

	token, err := utils.EncryptToken(i.config.Key, string(roleJson))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error encrypting token: %v", err)
	}

	return &tsr.UserAuthResponse{
		Token: token,
	}, nil
}

func (i *UserApi) RegisterUser(ctx context.Context, req *tsr.RegisterUserRequest) (*tsr.RegisterUserResponse, error) {
	var ur models.UserRole
	usr := &models.UserCreate{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Login:     req.Login,
		Password:  req.Password,
		Role:      req.Role,
	}

	str, err := utils.DecryptToken(i.config.Key, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error decrypting token: %v", err)
	}

	err = json.Unmarshal([]byte(str), &ur)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error unmarshalling userdata: %v", err)
	}

	uuid, err := i.userService.CreateUser(usr, &ur)
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
	return nil, nil
}
