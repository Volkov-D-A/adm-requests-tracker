package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User, token string) (string, error)
}

type UserApi struct {
	tsr.UnimplementedUserServiceServer
	userService UserService
}

func NewUserApi(userService UserService) *UserApi {
	return &UserApi{
		userService: userService,
	}
}

func (i *UserApi) UserAuth(ctx context.Context, req *tsr.UserAuthRequest) (*tsr.UserAuthResponse, error) {
	if req.Login == "user" && req.Password == "password" {
		return &tsr.UserAuthResponse{
			Token: "RTk,jbknjgolisyish",
		}, nil
	}
	return nil, status.Error(codes.Unauthenticated, "Invalid login or password")
}

func (i *UserApi) RegisterUser(ctx context.Context, req *tsr.RegisterUserRequest) (*tsr.RegisterUserResponse, error) {

	usr := &models.User{
		ID:        req.User.Uuid,
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Login:     req.User.Login,
		Password:  req.User.Password,
		Role:      req.User.Role,
	}

	uuid, err := i.userService.CreateUser(ctx, usr, req.Token)
	if err != nil {
		return nil, err
	}

	return &tsr.RegisterUserResponse{
		Uuid: uuid,
	}, nil
}
