package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type TSRService interface {
	AddTSR(ctx context.Context, tsr *models.TSR) (string, error)
}

type TSRApi struct {
	tsr.UnimplementedTSRServiceServer
	tsrService TSRService
}

func NewTSRApi(tsrService TSRService) *TSRApi {
	return &TSRApi{tsrService: tsrService}
}

func (t *TSRApi) CreateTSR(ctx context.Context, req *tsr.CreateTSRRequest) (*tsr.CreateTSRResponse, error) {
	return &tsr.CreateTSRResponse{
		Uuid: "jkbjkbnjgb",
	}, nil
}
