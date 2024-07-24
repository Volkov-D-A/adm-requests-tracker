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
	ApplyTSR(atsr *models.ApplyTSR, user_id string) error
	GetListTickets(mode, uuid, dep_uuid string) ([]models.ListTicketResponse, error)
	AddComment(comment *models.CommentAdd) (string, error)
	GetComments(tsrid string) ([]models.ResponseComments, error)
	GetFullTsrInfo(tsrid string) (*models.FullTsrInfo, error)
	RecordAction(act *models.ActionADD) error
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
	uuid, err := s.tsrStorage.Create(ctsr)
	if err != nil {
		return "", err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: ctsr.UserID, ObjectID: uuid, Action: "TsrAdd"})
	return uuid, nil
}

func (s *tsrService) TSREmployee(etsr *models.SetEmployee, token *models.UserToken) error {
	if token.Role != "admin" {
		return models.ErrUnauthorized
	}
	err := s.tsrStorage.TSREmployee(etsr)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: etsr.TSRId, Action: "SetEmployee", Info: etsr.UserID})
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
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: itsr.TSRId, Action: "SetImportance"})
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
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: ftsr.TSRId, Action: "TsrFinish"})
	return nil
}

func (s *tsrService) ApplyTSR(atsr *models.ApplyTSR, token *models.UserToken) error {
	//TODO сделать проверку принадлежности тикета по токену
	err := s.tsrStorage.ApplyTSR(atsr, token.ID)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: atsr.TSRId, Action: "TsrFinish"})
	return nil
}

func (s *tsrService) GetListTickets(mode string, token *models.UserToken) ([]models.ListTicketResponse, error) {
	res, err := s.tsrStorage.GetListTickets(mode, token.ID, token.Department) //TODO сделать проверку ролей по токену
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *tsrService) SetComment(comment *models.CommentAdd) error {
	uuid, err := s.tsrStorage.AddComment(comment)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: comment.UserID, ObjectID: comment.TsrID, Action: "TsrFinish", Info: uuid})
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

func (s *tsrService) GetFullTsrInfo(token *models.UserToken, tsrid string) (*models.FullTsrInfo, error) {
	//TODO проверка роли пользователя или принадлежности задачи

	res, err := s.tsrStorage.GetFullTsrInfo(tsrid)
	if err != nil {
		return nil, fmt.Errorf("error while getting comments: %v", err)
	}

	return res, nil
}
