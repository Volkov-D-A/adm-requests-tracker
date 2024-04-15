package service

import (
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type TSRStorage interface {
	Create(ctsr *models.CreateTSR) (string, error)
}

type tsrService struct {
	tsrStorage TSRStorage
}

func NewTSRService(tsrStorage TSRStorage) *tsrService {
	return &tsrService{
		tsrStorage: tsrStorage,
	}
}

func (s *tsrService) AddTSR(ctsr *models.CreateTSR) (string, error) {
	res, err := s.tsrStorage.Create(ctsr)
	if err != nil {
		return "", err
	}
	return res, nil
}
