package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"github.com/volkov-d-a/adm-requests-tracker/internal/service"
	"google.golang.org/grpc/status"
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
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Login:     req.User.Login,
		Password:  req.User.Password,
		Role:      req.User.Role,
	}

	id, err := i.userService.Create(ctx, usr, req.Token)
	if err != nil {
		return nil, status.Error(6, "Token not valid")
	}

	return &tsr.RegisterUserResponse{
		Id: id,
	}, nil
}
