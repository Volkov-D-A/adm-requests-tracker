package models

import "errors"

var ErrUnauthenticated = errors.New("wrong login or password")
var ErrUnauthorized = errors.New("not enought permissions")
var ErrUserAlreadyExists = errors.New("user with this login already exists")
var ErrUserNotExist = errors.New("user with this uuid not exist")
var ErrTicketNotExist = errors.New("ticket with this uuid not exist")
var ErrUserNotEmployee = errors.New("user with this uuid not employee")
