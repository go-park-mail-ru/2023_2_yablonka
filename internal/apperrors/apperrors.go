package apperrors

import (
	"errors"
	"net/http"
)

// TODO User, Board related errors
var (
	ErrEnvNotFound            = errors.New("invalid .env location was provided to the server setup")
	ErrUserNotFound           = errors.New("no user that matches the provided credentials")
	ErrWrongPassword          = errors.New("wrong password")
	ErrUserAlreadyExists      = errors.New("user with this email already exists")
	ErrJWTSecretMissing       = errors.New("no JWT secret was found in the environment")
	ErrSessionDurationMissing = errors.New("no session duration settings were found in the environment")
	ErrSessionNullDuration    = errors.New("provided session duration is zero seconds")
	ErrSessionIDLengthMissing = errors.New("session ID length is missing")
	ErrSessionNullIDLength    = errors.New("session ID length is zero")
	ErrSessionExpired         = errors.New("user session has expired")
	ErrJWTWrongMethod         = errors.New("provided token has the wrong signing method")
	ErrJWTInvalidToken        = errors.New("provided token is invalid")
	ErrJWTMissingClaim        = errors.New("provided token is missing a required claim")
	ErrJWTOldToken            = errors.New("provided token has expired")
	ErrSessionNotFound        = errors.New("no session found for provided session ID")
	ErrBoardNotFound          = errors.New("no board found for provided user ID")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var GenericUnauthorizedResponse = ErrorResponse{
	Code:    http.StatusUnauthorized,
	Message: "Ошибка авторизации",
}

var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "Ошибка сервера",
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
	ErrUserAlreadyExists: {
		Code:    http.StatusConflict,
		Message: "Пользователь с таким адресом почты уже существует",
	},
	ErrJWTSecretMissing:       InternalServerErrorResponse,
	ErrJWTWrongMethod:         GenericUnauthorizedResponse,
	ErrJWTInvalidToken:        GenericUnauthorizedResponse,
	ErrJWTMissingClaim:        GenericUnauthorizedResponse,
	ErrJWTOldToken:            GenericUnauthorizedResponse,
	ErrSessionDurationMissing: InternalServerErrorResponse,
	ErrSessionNullDuration:    InternalServerErrorResponse,
	ErrSessionIDLengthMissing: InternalServerErrorResponse,
	ErrSessionNullIDLength:    InternalServerErrorResponse,
	ErrSessionExpired:         GenericUnauthorizedResponse,
	ErrSessionNotFound:        GenericUnauthorizedResponse,
}
