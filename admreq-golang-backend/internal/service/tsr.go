package service

import (
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type TSRStorage interface {
	Create(ctsr *models.CreateTSR) (string, error)
	TSREmployee(etsr *models.SetEmployee) error
	TSRImportance(itsr *models.SetImportant) error
	FinishTSR(ftsr *models.FinishTSR, employee_id string) error
	GetListTickets(mode, uuid string) ([]models.ListTicketResponse, error)
	AddComment(comment *models.CommentAdd) error
	GetComments(tsrid string) ([]models.ResponseComments, error)
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

func (s *tsrService) TSRImportance(itsr *models.SetImportant, token *models.UserToken) error {
	if token.Role != "admin" {
		return models.ErrUnauthorized
	}
	err := s.tsrStorage.TSRImportance(itsr)
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

func (s *tsrService) GetListTickets(mode string, token *models.UserToken) ([]models.ListTicketResponse, error) {
	res, err := s.tsrStorage.GetListTickets(mode, token.ID) //TODO сделать проверку ролей по токену
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *tsrService) SetComment(comment *models.CommentAdd) error {
	err := s.tsrStorage.AddComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *tsrService) GetComments(token *models.UserToken, tsrid string) ([]models.ResponseComments, error) {
	//TODO проверка роли пользователя или принадлежности токена
	res, err := s.tsrStorage.GetComments(tsrid)
	if err != nil {
		return nil, fmt.Errorf("error while getting comments: %v", err)
	}

	return res, nil
}
