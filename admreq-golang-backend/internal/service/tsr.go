package service

import (
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type TSRStorage interface {
	Create(ctsr *models.CreateTSR) (string, error)
	TSREmployee(etsr *models.SetEmployee) error
	FinishTSR(ftsr *models.FinishTSR, employee_id string) error
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

func (s *tsrService) TSREmployee(etsr *models.SetEmployee, token *models.UserToken) error {
	if token.Role != "admin" {
		return models.ErrUnauthorized
	}
	err := s.tsrStorage.TSREmployee(etsr)
	if err != nil {
		return err
	}
	return nil
}

func (s *tsrService) FinishTSR(ftsr *models.FinishTSR, token *models.UserToken) error {
	if token.Role == "user" {
		return models.ErrUnauthorized
	}

	err := s.tsrStorage.FinishTSR(ftsr, token.ID)
	if err != nil {
		return err
	}
	return nil
}
