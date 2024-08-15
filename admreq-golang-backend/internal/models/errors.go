package models

import "errors"

var ErrUnauthenticated = errors.New("wrong login or password")
var ErrUnauthorized = errors.New("not enought permissions")
var ErrUserAlreadyExists = errors.New("user with this login already exists")
var ErrUserNotExist = errors.New("user with this uuid not exist")
var ErrDepartmentsNotExist = errors.New("queryed departments not exist")
var ErrTicketNotExist = errors.New("ticket with this uuid not exist")
var ErrUserNotEmployee = errors.New("user with this uuid not employee this ticket")
var ErrUserNotOwnTicket = errors.New("user with this uuid not create this ticket and cant apply it")
var ErrRowAlreadyExists = errors.New("row with this data unique and already exists")
var ErrInvalidDataInRequest = errors.New("invalid requested data")
