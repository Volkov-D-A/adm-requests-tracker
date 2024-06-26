package api

import (
	"context"

	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TSRService interface {
	AddTSR(ctsr *models.CreateTSR) (string, error)
	TSREmployee(etsr *models.SetEmployee, token *models.UserToken) error
	TSRImportance(itsr *models.SetImportant, token *models.UserToken) error
	FinishTSR(ftsr *models.FinishTSR, token *models.UserToken) error
	GetListTickets(mode string, token *models.UserToken) ([]models.ListTicketResponse, error)
	SetComment(comment *models.CommentAdd) error
	GetComments(token *models.UserToken, tsrid string) ([]models.ResponseComments, error)
	GetFullTsrInfo(token *models.UserToken, tsrid string) (*models.FullTsrInfo, error)
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
			return nil, status.Errorf(codes.Internal, "error deleting user: %v", err)
		}
	}

	return &tsr.CreateTSRResponse{
		Uuid: res,
	}, nil
}

func (t *TSRApi) EmployeeTSR(ctx context.Context, req *tsr.EmployeeTSRRequest) (*tsr.EmployeeTSRResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	etsr := &models.SetEmployee{
		UserID: req.EmployeeUuid,
		TSRId:  req.TsrUuid,
	}
	err = t.tsrService.TSREmployee(etsr, ut)
	if err != nil {
		switch err {
		case models.ErrUnauthorized, models.ErrUserNotEmployee:
			return nil, status.Error(codes.PermissionDenied, err.Error())
		case models.ErrTicketNotExist, models.ErrUserNotExist:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "error setting employee for task: %v", err)
		}
	}
	return &tsr.EmployeeTSRResponse{}, nil
}

func (t *TSRApi) ImportanceTSR(ctx context.Context, req *tsr.ImportanceTSRRequest) (*tsr.ImportanceTSRResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	itsr := &models.SetImportant{
		TSRId:     req.TsrUuid,
		Important: req.Important,
	}
	err = t.tsrService.TSRImportance(itsr, ut)
	if err != nil {
		switch err {
		case models.ErrUnauthorized:
			return nil, status.Error(codes.PermissionDenied, err.Error())
		case models.ErrTicketNotExist:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "error setting importance for task: %v", err)
		}
	}
	return &tsr.ImportanceTSRResponse{}, nil
}

func (t *TSRApi) FinishTSR(ctx context.Context, req *tsr.FinishTSRRequest) (*tsr.FinishTSRResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	ftsr := &models.FinishTSR{
		TSRId: req.TsrUuid,
	}

	err = t.tsrService.FinishTSR(ftsr, ut)
	if err != nil {
		switch err {
		case models.ErrTicketNotExist:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "error finishing ticket: %v", err)
		}
	}
	return &tsr.FinishTSRResponse{}, nil
}

func (t *TSRApi) GetListTickets(ctx context.Context, req *tsr.GetListTicketRequest) (*tsr.GetListTicketResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	res, err := t.tsrService.GetListTickets(req.Mode, ut)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting tickets: %v", err)
	}

	result := make([]*tsr.GetListTicketResponse_Ticket, len(res))
	for z, x := range res {
		var eminitials string
		if x.EmployeFirstname.Valid {
			eminitials = x.EmployeLastname.String + " " + string([]rune(x.EmployeFirstname.String)[0]) + "." + string([]rune(x.EmployeSurname.String)[0]) + "."
		} else {
			eminitials = "-"
		}
		result[z] = &tsr.GetListTicketResponse_Ticket{
			Id:               x.ID,
			Text:             x.Text,
			CreatedAt:        timestamppb.New(x.CreatedAt),
			UserInitials:     x.UserLastname + " " + string([]rune(x.UserFirstname)[0]) + "." + string([]rune(x.UserSurname)[0]) + ".",
			EmployeeInitials: eminitials,
		}
	}

	return &tsr.GetListTicketResponse{Tickets: result}, nil
}

func (t *TSRApi) SetTsrComment(ctx context.Context, req *tsr.SetTsrCommentRequest) (*tsr.SetTsrCommentResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	comment := &models.CommentAdd{
		UserID:      ut.ID,
		TsrID:       req.TsrId,
		TextComment: req.CommentText,
	}

	err = t.tsrService.SetComment(comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error setting cooment: %v", err)
	}
	return &tsr.SetTsrCommentResponse{}, nil
}

func (t *TSRApi) GetTsrCommnts(ctx context.Context, req *tsr.GetTsrCommentsRequest) (*tsr.GetTsrCommentsResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	res, err := t.tsrService.GetComments(ut, req.TsrId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting comments: %v", err)
	}

	if len(res) == 0 {
		return &tsr.GetTsrCommentsResponse{
			Count:    int32(0),
			Comments: nil,
		}, nil
	}

	result := make([]*tsr.GetTsrCommentsResponse_Comment, len(res))
	for z, x := range res {
		result[z] = &tsr.GetTsrCommentsResponse_Comment{
			Firstname:   x.Firstname,
			Lastname:    x.Lastname,
			Surname:     x.Surname,
			CommentText: x.TextComment,
			PostedAt:    timestamppb.New(x.PostedAt),
		}
	}

	return &tsr.GetTsrCommentsResponse{
		Count:    int32(len(res)),
		Comments: result,
	}, nil
}

func (t *TSRApi) GetFullTsrInfo(ctx context.Context, req *tsr.GetFullTsrInfoRequest) (*tsr.GetFullTsrInfoResponse, error) {
	var result tsr.GetFullTsrInfoResponse
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	res, err := t.tsrService.GetFullTsrInfo(ut, req.TsrId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting tsr info: %v", err)
	}

	if res.EmployeFirstname.Valid {
		result.EmployeeFirstname = res.EmployeFirstname.String
		result.EmployeeLastname = res.EmployeLastname.String
		result.EmployeeSurname = res.EmployeSurname.String
	}
	if res.FinishedAt.Valid {
		result.FinishedAt = timestamppb.New(res.FinishedAt.Time)
	}

	result.Id = res.ID
	result.Text = res.Text
	result.UserFirstname = res.UserFirstname
	result.UserLastname = res.UserLastname
	result.UserSurname = res.UserSurname
	result.UserDepartment = res.UserDepartment
	result.PostedAt = timestamppb.New(res.CreatedAt)
	result.Important = res.Important
	result.Finished = res.Finished

	return &result, nil
}
