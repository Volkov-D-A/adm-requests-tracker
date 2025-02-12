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
	CreateTSR(ctsr *models.CreateTSR, token *models.UserToken) error
	EmployeeTSR(etsr *models.SetEmployee, token *models.UserToken) error
	ImportanceTSR(itsr *models.SetImportant, token *models.UserToken) error
	FinishTSR(ftsr *models.FinishTSR, token *models.UserToken) error
	ApplyTSR(atsr *models.ApplyTSR, token *models.UserToken) error
	RejectTSR(rtsr *models.RejectTSR, token *models.UserToken) error
	GetListTickets(mode string, token *models.UserToken) ([]models.ListTicketResponse, error)
	AddTsrComment(comment *models.CommentAdd, token *models.UserToken) error
	GetTsrComments(token *models.UserToken, tsrid string) ([]models.ResponseComments, error)
	GetFullTsrInfo(token *models.UserToken, tsrid string) (*models.FullTsrInfo, error)
	GetTsrStat(token *models.UserToken, target_dep string) (*models.FullStat, error)
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
		UserID:           ut.UserID,
		Text:             req.Text,
		TargetDepartment: req.TargetDep,
	}

	err = t.tsrService.CreateTSR(ctsr, ut)

	switch err {
	case nil:
		return &tsr.CreateTSRResponse{}, nil
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrInvalidDataInRequest:
		return nil, status.Error(codes.InvalidArgument, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error while creating tsr: %v", err)
	}
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
	err = t.tsrService.EmployeeTSR(etsr, ut)

	switch err {
	case nil:
		return &tsr.EmployeeTSRResponse{}, nil
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrTicketNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "error setting employee for task: %v", err)
	}

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
	err = t.tsrService.ImportanceTSR(itsr, ut)

	switch err {
	case nil:
		return &tsr.ImportanceTSRResponse{}, nil
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	case models.ErrTicketNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error setting importance for task: %v", err)
	}

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

	switch err {
	case nil:
		return &tsr.FinishTSRResponse{}, nil
	case models.ErrTicketNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	case models.ErrUnauthorized, models.ErrUserNotEmployee:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error finishing ticket: %v", err)
	}

}

func (t *TSRApi) ApplyTSR(ctx context.Context, req *tsr.ApplyTSRRequest) (*tsr.ApplyTSRResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	atsr := &models.ApplyTSR{
		TSRId: req.TsrUuid,
	}

	err = t.tsrService.ApplyTSR(atsr, ut)
	switch err {
	case nil:
		return &tsr.ApplyTSRResponse{}, nil
	case models.ErrTicketNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	case models.ErrUnauthorized, models.ErrUserNotOwnTicket:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error applying tsr: %v", err)
	}
}

func (t *TSRApi) RejectTSR(ctx context.Context, req *tsr.RejectTSRRequest) (*tsr.RejectTSRResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	rtsr := &models.RejectTSR{
		TSRId: req.TsrUuid,
	}

	err = t.tsrService.RejectTSR(rtsr, ut)
	switch err {
	case nil:
		return &tsr.RejectTSRResponse{}, nil
	case models.ErrUnauthorized, models.ErrUserNotOwnTicket:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error rejecting tsr: %v", err)
	}

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
		if x.EmployeeID.Valid {
			eminitials = x.EmployeeLastname.String + " " + string([]rune(x.EmployeeFirstname.String)[0]) + "." + string([]rune(x.EmployeeSurname.String)[0]) + "."
		} else {
			eminitials = ""
		}
		result[z] = &tsr.GetListTicketResponse_Ticket{
			Id:               x.ID,
			Text:             x.Text,
			CreatedAt:        timestamppb.New(x.CreatedAt),
			UserId:           x.UserID,
			UserInitials:     x.UserLastname + " " + string([]rune(x.UserFirstname)[0]) + "." + string([]rune(x.UserSurname)[0]) + ".",
			UserDepartment:   x.UserDepartment,
			EmployeeId:       x.EmployeeID.String,
			EmployeeInitials: eminitials,
			Important:        x.Important,
			Finished:         x.Finished,
			UnreadMessages:   x.UnreadMessages,
		}
	}

	return &tsr.GetListTicketResponse{Tickets: result}, nil
}

func (t *TSRApi) AddTsrComment(ctx context.Context, req *tsr.AddTsrCommentRequest) (*tsr.AddTsrCommentResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	comment := &models.CommentAdd{
		UserID:      ut.UserID,
		TsrID:       req.TsrId,
		TextComment: req.CommentText,
	}

	err = t.tsrService.AddTsrComment(comment, ut)
	switch err {
	case nil:
		return &tsr.AddTsrCommentResponse{}, nil
	case models.ErrUnauthorized, models.ErrUserNotEmployee, models.ErrUserNotOwnTicket:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "error setting cooment: %v", err)
	}

}

func (t *TSRApi) GetTsrComments(ctx context.Context, req *tsr.GetTsrCommentsRequest) (*tsr.GetTsrCommentsResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}
	res, err := t.tsrService.GetTsrComments(ut, req.TsrId)

	switch err {
	case nil:
		break
	case models.ErrUnauthorized, models.ErrUserNotEmployee, models.ErrUserNotOwnTicket:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
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
			CommId:      x.ID,
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
	switch err {
	case nil:
		break
	case models.ErrTicketNotExist:
		return nil, status.Error(codes.NotFound, err.Error())
	case models.ErrUnauthorized, models.ErrUserNotEmployee, models.ErrUserNotOwnTicket:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled error getting tsr info: %v", err)
	}

	if res.EmployeeID.Valid {
		result.EmployeeId = res.EmployeeID.String
		result.EmployeeFirstname = res.EmployeeFirstname.String
		result.EmployeeLastname = res.EmployeeLastname.String
		result.EmployeeSurname = res.EmployeeSurname.String
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
	result.Applied = res.Applied

	return &result, nil
}

func (t *TSRApi) GetTsrStat(ctx context.Context, req *tsr.GetTsrStatRequest) (*tsr.GetTsrStatResponse, error) {
	ut, err := getTokenData(req.Token, t.config.Key)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user rights: %v", err)
	}

	res, err := t.tsrService.GetTsrStat(ut, req.TargetDep)
	switch err {
	case nil:
		break
	case models.ErrUnauthorized:
		return nil, status.Error(codes.PermissionDenied, err.Error())
	default:
		return nil, status.Errorf(codes.Internal, "unhandled get statistic error: %v", err)
	}

	byDepartment := make([]*tsr.GetTsrStatResponseStatDep, len(res.ByDepartment))
	for z, x := range res.ByDepartment {
		byDepartment[z] = &tsr.GetTsrStatResponseStatDep{
			DepartmentName: x.DepartmentName,
			TsrInWork:      x.TsrInWork,
			TsrFinished:    x.TsrFinished,
			TsrApplyed:     x.TsrApplyed,
		}
	}

	byEmployee := make([]*tsr.GetTsrStatResponseStatEmployee, len(res.ByEmployee))
	for z, x := range res.ByEmployee {
		byEmployee[z] = &tsr.GetTsrStatResponseStatEmployee{
			EmployeeName: x.EmployeeName,
			TsrInWork:    x.TsrInWork,
			TsrFinished:  x.TsrFinished,
			TsrApplyed:   x.TsrApplyed,
		}
	}

	return &tsr.GetTsrStatResponse{ByDepartment: byDepartment, ByEmployee: byEmployee}, nil
}
