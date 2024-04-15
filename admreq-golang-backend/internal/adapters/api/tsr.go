package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TSRService interface {
	AddTSR(ctsr *models.CreateTSR) (string, error)
}

type TSRApi struct {
	tsr.UnimplementedTSRServiceServer
	tsrService TSRService
	config     *TSRConfig
}

type TSRConfig struct {
	Key string
}

func NewTSRApi(tsrService TSRService, config *TSRConfig) *TSRApi {
	return &TSRApi{
		tsrService: tsrService,
		config:     config,
	}
}

func (t *TSRApi) CreateTSR(ctx context.Context, req *tsr.CreateTSRRequest) (*tsr.CreateTSRResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	ctsr := &models.CreateTSR{
		UserID: ut.ID,
		Text:   req.Text,
	}

	res, err := t.tsrService.AddTSR(ctsr)
	if err != nil {
		switch err {
		case models.ErrUserNotExist:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "error deleting user")
		}
	}

	return &tsr.CreateTSRResponse{
		Uuid: res,
	}, nil
}
