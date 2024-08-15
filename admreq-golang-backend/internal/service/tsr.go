package service

import (
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
)

type TSRStorage interface {
	CreateTSR(ctsr *models.CreateTSR) (string, error)
	EmployeeTSR(etsr *models.SetEmployee) error
	ImportanceTSR(itsr *models.SetImportant) error
	FinishTSR(ftsr *models.FinishTSR) error
	ApplyTSR(atsr *models.ApplyTSR) error
	GetListTickets(mode, uuid, dep_uuid string) ([]models.ListTicketResponse, error)
	AddTsrComment(comment *models.CommentAdd) (string, error)
	GetTsrComments(tsrid string) ([]models.ResponseComments, error)
	GetFullTsrInfo(tsrid string) (*models.FullTsrInfo, error)
	RecordAction(act *models.ActionADD) error
	GetDepartmentsList() ([]models.Department, error)
	GetStatByDepartment(req *models.StatByDepartmentReq) (*models.StatByDepartment, error)
	GetEmployeeList(target_dep string) ([]models.Employee, error)
	GetStatByEmployee(req *models.StatByEmployeeReq) (*models.StatByEmployee, error)
	CheckTSROwn(user_uuid, tsr_uuid, mode string) (bool, error)
}

type tsrService struct {
	tsrStorage TSRStorage
}

func NewTSRService(tsrStorage TSRStorage) *tsrService {
	return &tsrService{
		tsrStorage: tsrStorage,
	}
}

func (s *tsrService) CreateTSR(ctsr *models.CreateTSR, token *models.UserToken) error {
	if !token.Rights.Create {
		return models.ErrUnauthorized
	}

	if ctsr.TargetDepartment == "" || ctsr.Text == "" {
		return models.ErrInvalidDataInRequest
	}
	uuid, err := s.tsrStorage.CreateTSR(ctsr)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: ctsr.UserID, ObjectID: uuid, Action: "TsrAdd"})
	return nil
}

func (s *tsrService) EmployeeTSR(etsr *models.SetEmployee, token *models.UserToken) error {
	if !token.Rights.Admin {
		return models.ErrUnauthorized
	}
	err := s.tsrStorage.EmployeeTSR(etsr)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: etsr.TSRId, Action: "SetEmployee", Info: etsr.UserID})
	return nil
}

func (s *tsrService) ImportanceTSR(itsr *models.SetImportant, token *models.UserToken) error {
	if !token.Rights.Admin {
		return models.ErrUnauthorized
	}
	err := s.tsrStorage.ImportanceTSR(itsr)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: itsr.TSRId, Action: "SetImportance"})
	return nil
}

func (s *tsrService) FinishTSR(ftsr *models.FinishTSR, token *models.UserToken) error {
	if !token.Rights.Employee {
		return models.ErrUnauthorized
	}

	chk, err := s.tsrStorage.CheckTSROwn(token.ID, ftsr.TSRId, "employee")
	if err != nil {
		return err
	}
	if !chk {
		return models.ErrUserNotEmployee
	}
	err = s.tsrStorage.FinishTSR(ftsr)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: ftsr.TSRId, Action: "TsrFinish"})
	return nil

}

func (s *tsrService) ApplyTSR(atsr *models.ApplyTSR, token *models.UserToken) error {
	if !token.Rights.Create {
		return models.ErrUnauthorized
	}

	chk, err := s.tsrStorage.CheckTSROwn(token.ID, atsr.TSRId, "user")
	if err != nil {
		return err
	}
	if !chk {
		return models.ErrUserNotOwnTicket
	}

	err = s.tsrStorage.ApplyTSR(atsr)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: token.ID, ObjectID: atsr.TSRId, Action: "TsrApply"})
	return nil
}

func (s *tsrService) GetListTickets(mode string, token *models.UserToken) ([]models.ListTicketResponse, error) {
	if !token.Rights.Create && !token.Rights.Employee && !token.Rights.Admin && !token.Rights.Archiv {
		return nil, models.ErrUnauthorized
	}
	res, err := s.tsrStorage.GetListTickets(mode, token.ID, token.Department)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *tsrService) AddTsrComment(comment *models.CommentAdd, token *models.UserToken) error {
	if !token.Rights.Admin {
		if !token.Rights.Employee {
			if !token.Rights.Create {
				return models.ErrUnauthorized
			}
			chk, err := s.tsrStorage.CheckTSROwn(comment.UserID, comment.TsrID, "user")
			if err != nil {
				return err
			}
			if !chk {
				return models.ErrUserNotOwnTicket
			}
		}
		chk, err := s.tsrStorage.CheckTSROwn(comment.UserID, comment.TsrID, "employee")
		if err != nil {
			return err
		}
		if !chk {
			return models.ErrUserNotEmployee
		}
	}

	if comment.TextComment == "" {
		return models.ErrInvalidDataInRequest
	}

	uuid, err := s.tsrStorage.AddTsrComment(comment)
	if err != nil {
		return err
	}
	s.tsrStorage.RecordAction(&models.ActionADD{SubjectID: comment.UserID, ObjectID: comment.TsrID, Action: "AddComment", Info: uuid})
	return nil
}

func (s *tsrService) GetTsrComments(token *models.UserToken, tsrid string) ([]models.ResponseComments, error) {
	if !token.Rights.Admin {
		if !token.Rights.Employee {
			if !token.Rights.Create {
				return nil, models.ErrUnauthorized
			}
			chk, err := s.tsrStorage.CheckTSROwn(token.ID, tsrid, "user")
			if err != nil {
				return nil, err
			}
			if !chk {
				return nil, models.ErrUserNotOwnTicket
			}
		}
		chk, err := s.tsrStorage.CheckTSROwn(token.ID, tsrid, "employee")
		if err != nil {
			return nil, err
		}
		if !chk {
			return nil, models.ErrUserNotEmployee
		}
	}
	res, err := s.tsrStorage.GetTsrComments(tsrid)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *tsrService) GetFullTsrInfo(token *models.UserToken, tsrid string) (*models.FullTsrInfo, error) {
	if !token.Rights.Admin && !token.Rights.Archiv {
		if !token.Rights.Employee {
			if !token.Rights.Create {
				return nil, models.ErrUnauthorized
			}
			chk, err := s.tsrStorage.CheckTSROwn(token.ID, tsrid, "user")
			if err != nil {
				return nil, err
			}
			if !chk {
				return nil, models.ErrUserNotOwnTicket
			}
		}
		chk, err := s.tsrStorage.CheckTSROwn(token.ID, tsrid, "employee")
		if err != nil {
			return nil, err
		}
		if !chk {
			return nil, models.ErrUserNotEmployee
		}
	}

	res, err := s.tsrStorage.GetFullTsrInfo(tsrid)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *tsrService) GetTsrStat(token *models.UserToken, target_dep string) (*models.FullStat, error) {
	if !token.Rights.Stat {
		return nil, models.ErrUnauthorized
	}

	deps, err := s.tsrStorage.GetDepartmentsList()
	if err != nil {
		return nil, err
	}

	empls, err := s.tsrStorage.GetEmployeeList(target_dep)
	if err != nil {
		return nil, err
	}

	byDepartment := make([]*models.StatByDepartment, 0)

	for _, y := range deps {
		res, err := s.tsrStorage.GetStatByDepartment(&models.StatByDepartmentReq{TargetDepartmentUUID: target_dep, SourceDepartmentUUID: y.ID})
		if err != nil {
			return nil, err
		}
		res.DepartmentName = y.DepartmentName
		byDepartment = append(byDepartment, res)
	}

	byEmployee := make([]*models.StatByEmployee, 0)

	for _, y := range empls {
		res, err := s.tsrStorage.GetStatByEmployee(&models.StatByEmployeeReq{EmplotyeeUUID: y.ID})
		if err != nil {
			return nil, err
		}
		res.EmployeeName = y.Lastname + " " + y.Firstname + " " + y.Surname
		byEmployee = append(byEmployee, res)
	}

	return &models.FullStat{ByDepartment: byDepartment, ByEmployee: byEmployee}, nil
}
