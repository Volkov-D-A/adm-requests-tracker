package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"github.com/volkov-d-a/adm-requests-tracker/internal/service"
)

type Implementation struct {
	tsr.UnimplementedTsrServiceServer
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

func (i *Implementation) RegisterUser(ctx context.Context, req *tsr.RegisterUserRequest) (*tsr.RegisterUserResponse, error) {

	usr := &models.User{
		Id:        req.User.Id,
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Login:     req.User.Login,
		Password:  req.User.Password,
		Role:      req.User.Role,
	}

	id, err := i.userService.CreateUser(ctx, usr, req.Token)
	if err != nil {
		return nil, err
	}

	return &tsr.RegisterUserResponse{
		Id: id,
	}, nil
}
