package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/internal/pkg/dto"
)

// Ошибки, связанные с конфигурацией сервера
var (
	// ErrEnvNotFound ошибка: переданный при конфигурации .env файл не был найден
	ErrEnvNotFound = errors.New("invalid .env location was provided to the server setup")
	// ErrJWTSecretMissing ошибка: в полученном конфиге нет JWT ключа для подписи
	ErrJWTSecretMissing = errors.New("no JWT secret was found in the environment")
	// ErrSessionDurationMissing ошибка: в полученном конфиге нет времени жизни сессии
	ErrSessionDurationMissing = errors.New("no session duration settings were found in the environment")
	// ErrSessionNullDuration ошибка: в полученном конфиге время жизни сессии меньше секунды
	ErrSessionNullDuration = errors.New("provided session duration is zero seconds")
	// ErrSessionIDLengthMissing ошибка: в полученном конфиге нет длины ID сессии
	ErrSessionIDLengthMissing = errors.New("session ID length is missing")
	// ErrSessionNullIDLength ошибка: в полученном конфиге длина ID сессии равна нулю
	ErrSessionNullIDLength = errors.New("session ID length is zero")
)

// Ошибки, связанные с авторизацией
var (
	// ErrUserNotFound ошибка: нет пользователя с такими данными
	ErrUserNotFound = errors.New("no user that matches the provided credentials")
	// ErrUserNotFound ошибка: неправильный пароль
	ErrWrongPassword = errors.New("wrong password")
	// ErrUserAlreadyExists ошибка: пользователь с таким адресом почты уже существует
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	// ErrUserNotCreated ошибка: не удалось добавить пользователя в БД
	ErrUserNotCreated = errors.New("user couldn't be created")
	// ErrUserNotUpdated ошибка: не удалось обновить пользователя в БД
	ErrUserNotUpdated = errors.New("user couldn't be updated")
	// ErrUserNotDeleted ошибка: не удалось обновить пользователя в БД
	ErrUserNotDeleted = errors.New("user couldn't be deleted")
)

// Ошибки, связанные с AuthService
var (
	// ErrJWTWrongMethod ошибка: у полученного JWT неправильный метод подписи
	ErrJWTWrongMethod = errors.New("provided token has the wrong signing method")
	// ErrTokenNotGenerated ошибка: у полученного JWT неправильный метод подписи
	ErrTokenNotGenerated = errors.New("the system had trouble generating a session token")
	// ErrJWTInvalidToken ошибка: полученный JWT не валиден
	ErrJWTInvalidToken = errors.New("provided token is invalid")
	// ErrJWTMissingClaim ошибка: у полученного JWT отсутствует необходимое поле
	ErrJWTMissingClaim = errors.New("provided token is missing a required claim")
	// ErrJWTOldToken ошибка: время действия JWT истекло
	ErrJWTOldToken = errors.New("provided token has expired")
	// ErrSessionExpired ошибка: время действия сессии истекло
	ErrSessionExpired = errors.New("user session has expired")
	// ErrSessionNotFound ошибка: полученной сессии нет в хранилище
	ErrSessionNotFound = errors.New("no session found for provided session ID")
	// ErrSessionNotCreated ошибка: полученной сессии нет в хранилище
	ErrSessionNotCreated = errors.New("session couldn't be created")
)

// Ошибки, связанные с сервером
var (
	// ErrCouldNotBuildQuery ошибка: не удалось сформировать SQL запрос
	ErrCouldNotBuildQuery = errors.New("error building an SQL query")
)

// Ошибки, связанные с BoardService
var (
	// ErrBoardNotFound ошибка: доски с полученным ID не существует
	ErrBoardNotFound = errors.New("no board found for provided board ID")
)

// Ошибки, связанные с WorkspaceService
var (
	// ErrWorkspaceNotDeleted ошибка: не удалось создать рабочее прострнство в БД
	ErrWorkspaceNotCreated = errors.New("workspace couldn't be created")
	// ErrWorkspaceNotDeleted ошибка: не удалось получить рабочее прострнство в БД
	ErrCouldNotGetWorkspace = errors.New("workspace couldn't be retreived")
	// ErrWorkspaceNotDeleted ошибка: не удалось удалить рабочее прострнство в БД
	ErrWorkspaceNotDeleted = errors.New("user couldn't be deleted")
)

// Ошибки, связанные с ListService
var (
	// ErrListNotCreated ошибка: не удалось создать список в БД
	ErrListNotCreated = errors.New("list couldn't be created")
	// ErrListNotUpdated ошибка: не удалось получить список в БД
	ErrListNotUpdated = errors.New("list couldn't be updated")
	// ErrListNotDeleted ошибка: не удалось удалить список в БД
	ErrListNotDeleted = errors.New("list couldn't be deleted")
)

// Ошибки, связанные с TaskService
var (
	// ErrTaskNotCreated ошибка: не удалось создать задание в БД
	ErrTaskNotCreated = errors.New("task couldn't be created")
	// ErrTaskNotUpdated ошибка: не удалось получить задание в БД
	ErrTaskNotUpdated = errors.New("task couldn't be updated")
	// ErrTaskNotDeleted ошибка: не удалось удалить задание в БД
	ErrTaskNotDeleted = errors.New("task couldn't be deleted")
)

// ErrorResponse
// структура для обёртки ошибок приложения в ответ бэкэнд-сервера со статусом
type ErrorResponse struct {
	Code    int    `json:"-"`
	Message string `json:"error_response"`
}

// WrongLoginResponse
// заглушка для ответа 401 без разглашения имплементации чувствительных процессов
var WrongLoginResponse = ErrorResponse{
	Code:    http.StatusUnauthorized,
	Message: "Неверный адрес почты или пароль",
}

// BadRequestResponse
// заглушка для ответа 401 без разглашения имплементации чувствительных процессов
var BadRequestResponse = ErrorResponse{
	Code:    http.StatusBadRequest,
	Message: "Ошибка авторизации",
}

// GenericUnauthorizedResponse
// заглушка для ответа 401 без разглашения имплементации чувствительных процессов
var GenericUnauthorizedResponse = ErrorResponse{
	Code:    http.StatusUnauthorized,
	Message: "Ошибка авторизации",
}

// InternalServerErrorResponse
// заглушка для ответа 500 без разглашения имплементации чувствительных процессов
var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "Ошибка сервера",
}

// InternalServerErrorResponse
// заглушка для ответа 409 без разглашения имплементации чувствительных процессов
var StatusConflictResponse = ErrorResponse{
	Code:    http.StatusConflict,
	Message: "Пользователь с таким адресом почты уже существует",
}

// ErrorMap
// карта для связи ошибок приложения и ответа бэкэнд-сервера
var ErrorMap = map[error]ErrorResponse{
	ErrUserNotFound:           WrongLoginResponse,
	ErrWrongPassword:          WrongLoginResponse,
	ErrUserAlreadyExists:      StatusConflictResponse,
	ErrUserNotCreated:         InternalServerErrorResponse,
	ErrJWTSecretMissing:       InternalServerErrorResponse,
	ErrTokenNotGenerated:      InternalServerErrorResponse,
	ErrJWTWrongMethod:         GenericUnauthorizedResponse,
	ErrJWTInvalidToken:        GenericUnauthorizedResponse,
	ErrJWTMissingClaim:        GenericUnauthorizedResponse,
	ErrJWTOldToken:            GenericUnauthorizedResponse,
	ErrSessionDurationMissing: InternalServerErrorResponse,
	ErrSessionNullDuration:    InternalServerErrorResponse,
	ErrSessionIDLengthMissing: InternalServerErrorResponse,
	ErrSessionNullIDLength:    InternalServerErrorResponse,
	ErrSessionNotCreated:      InternalServerErrorResponse,
	ErrSessionExpired:         GenericUnauthorizedResponse,
	ErrCouldNotBuildQuery:     InternalServerErrorResponse,
	ErrSessionNotFound:        GenericUnauthorizedResponse,
	ErrWorkspaceNotCreated:    InternalServerErrorResponse,
	ErrCouldNotGetWorkspace:   InternalServerErrorResponse,
	ErrWorkspaceNotDeleted:    InternalServerErrorResponse,
	ErrBoardNotFound:          InternalServerErrorResponse,
}

func ErrorJSON(err ErrorResponse) []byte {
	response := dto.JSONResponse{
		Body: err,
	}
	jsonResponse, _ := json.Marshal(response)
	return jsonResponse
}
