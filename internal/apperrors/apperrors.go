package apperrors

import (
	"errors"
	"net/http"
)

// TODO User, Board related errors
var (
	ErrUserNotFound  = errors.New("no user that matches the provided credentials")
	ErrWrongPassword = errors.New("wrong password")
	ErrUserExists    = errors.New("user with this email already exists")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var ErrorMap = map[error]ErrorResponse{
	ErrUserNotFound: {
		Code:    http.StatusNotFound,
		Message: "Пользователя не существует",
	},
	ErrWrongPassword: {
		Code:    http.StatusUnauthorized,
		Message: "Указан неправильный пароль",
	},
	ErrUserExists: {
		Code:    http.StatusConflict,
		Message: "Пользователь с таким адресом почты уже существует",
	},
}
