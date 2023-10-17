package controller

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type TSRService interface {
	CreateUser(ctx context.Context, user *models.User, token string) (int32, error)
}

type TSRController struct {
	tsr.UnimplementedTsrServiceServer
	tsrService TSRService
}

func New(tsrService TSRService) *TSRController {
	return &TSRController{
		tsrService: tsrService,
	}
}

func (i *TSRController) RegisterUser(ctx context.Context, req *tsr.RegisterUserRequest) (*tsr.RegisterUserResponse, error) {

	usr := &models.User{
		Id:        req.User.Id,
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Login:     req.User.Login,
		Password:  req.User.Password,
		Role:      req.User.Role,
	}

	id, err := i.tsrService.CreateUser(ctx, usr, req.Token)
	if err != nil {
		return nil, err
	}

	return &tsr.RegisterUserResponse{
		Id: id,
	}, nil
}
