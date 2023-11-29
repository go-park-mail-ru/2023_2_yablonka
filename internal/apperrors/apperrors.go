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
	// ErrDatabasePWMissing ошибка: в полученном конфиге нет пароля от БД
	ErrDatabasePWMissing = errors.New("database PW is missing")
	// ErrInvalidLoggingLevel ошибка: в полученном конфиге указан неправильный уровень логгирования
	ErrInvalidLoggingLevel = errors.New("incorrect logging level provided (accepted values -- debug, info, warning, error)")
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
	// ErrCouldNotGetUser ошибка: не удалось получить задание в БД
	ErrCouldNotGetUser = errors.New("couldn't get User")
)

// Ошибки, связанные с CSAT
var (
	// ErrCouldNotCreateQuestion ошибка: не удалось создать вопрос CSAT в БД
	ErrCouldNotCreateQuestion = errors.New("couldn't create question")
	// ErrCouldNotGetQuestions ошибка: не удалось получить все вопросы CSAT в БД
	ErrCouldNotGetQuestions = errors.New("couldn't get all questions")
	// ErrCouldNotGetQuestionType ошибка: не удалось получить тип вопроса CSAT в БД
	ErrCouldNotGetQuestionType = errors.New("couldn't get question type")
	// ErrCouldNotGetQuestionType ошибка: не удалось обновить вопрос CSAT в БД
	ErrQuestionNotUpdated = errors.New("couldn't обновить question")
	// ErrCouldNotGetQuestionType ошибка: не удалось удалить вопрос CSAT в БД
	ErrQuestionNotDeleted = errors.New("couldn't delete question")
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
	// ErrCSRFNotFound ошибка: полученного CSRF нет в хранилище
	ErrCSRFNotFound = errors.New("provided CSRF not found in storage")
)

// Ошибки, связанные с Commentervice
var (
	// ErrCouldNotGetTaskComments ошибка: нельзя получить комментарии задания
	ErrCouldNotGetTaskComments = errors.New("couldn't get the comments from task")
)

// Ошибки, связанные с сервером
var (
	// ErrCouldNotBuildQuery ошибка: не удалось сформировать SQL запрос
	ErrCouldNotBuildQuery       = errors.New("error building an SQL query")
	ErrCouldNotStartTransaction = errors.New("error starting a transaction")
	ErrCouldNotCollectRows      = errors.New("couldn't collect rows")
)

// Ошибки, связанные с BoardService
var (
	// ErrWorkspaceNotDeleted ошибка: не удалось создать рабочее прострнство в БД
	ErrBoardNotCreated = errors.New("board couldn't be created")
	// ErrWorkspaceNotDeleted ошибка: не удалось получить рабочее прострнство в БД
	ErrBoardNotUpdated = errors.New("board couldn't be updated")
	// ErrWorkspaceNotDeleted ошибка: не удалось удалить рабочее прострнство в БД
	ErrBoardNotDeleted = errors.New("board couldn't be deleted")
	// ErrCouldNotGetBoard ошибка: доски с полученным ID не существует
	ErrCouldNotGetBoard = errors.New("could not retrieve board")
	// ErrCouldNotGetBoardUsers ошибка: пользователей у доски с полученным ID не существует
	ErrCouldNotGetBoardUsers = errors.New("could not retrieve board users")
	// ErrNoBoardAccess ошибка: у запрашивающего пользователя нет доступа к полученной доске
	ErrNoBoardAccess = errors.New("user has no access to board")
	// ErrCouldNotAddBoardUser ошибка: не удалось добавить пользователя на доску
	ErrCouldNotAddBoardUser = errors.New("couldn't add user to board")
	// ErrCouldNotRemoveBoardUser ошибка: не удалось добавить пользователя на доску
	ErrCouldNotRemoveBoardUser = errors.New("couldn't remove user from board")
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
	// ErrListNotDeleted ошибка: не удалось получить список из БД
	ErrCouldNotGetList = errors.New("could not get list")
	// ErrListNotCreated ошибка: не удалось создать список в БД
	ErrListNotCreated = errors.New("list couldn't be created")
	// ErrListNotUpdated ошибка: не удалось получить список в БД
	ErrListNotUpdated = errors.New("list couldn't be updated")
	// ErrListNotDeleted ошибка: не удалось удалить список в БД
	ErrListNotDeleted = errors.New("list couldn't be deleted")
)

// Ошибки, связанные с ChecklistService
var (
	// ErrCouldNotGetChecklist ошибка: не удалось получить чеклист из БД
	ErrCouldNotGetChecklist = errors.New("could not get checklist")
	// ErrChecklistNotCreated ошибка: не удалось создать чеклист в БД
	ErrChecklistNotCreated = errors.New("checklist couldn't be created")
	// ErrChecklistNotUpdated ошибка: не удалось получить чеклист в БД
	ErrChecklistNotUpdated = errors.New("checklist couldn't be updated")
	// ErrChecklistNotDeleted ошибка: не удалось удалить чеклист в БД
	ErrChecklistNotDeleted = errors.New("checklist couldn't be deleted")
)

// Ошибки, связанные с ChecklistItemService
var (
	// ErrCouldNotGetChecklist ошибка: не удалось получить элемент чеклиста из БД
	ErrCouldNotGetChecklistItem = errors.New("could not get checklist item")
	// ErrChecklistItemNotCreated ошибка: не удалось создать элемент чеклиста в БД
	ErrChecklistItemNotCreated = errors.New("checklist item couldn't be created")
	// ErrChecklistItemNotUpdated ошибка: не удалось получить элемент чеклиста в БД
	ErrChecklistItemNotUpdated = errors.New("checklist item couldn't be updated")
	// ErrChecklistItemNotDeleted ошибка: не удалось удалить элемент чеклиста в БД
	ErrChecklistItemNotDeleted = errors.New("checklist item couldn't be deleted")
)

// Ошибки, связанные с TaskService
var (
	// ErrTaskNotCreated ошибка: не удалось создать задание в БД
	ErrTaskNotCreated = errors.New("task couldn't be created")
	// ErrTaskNotUpdated ошибка: не удалось получить задание в БД
	ErrTaskNotUpdated = errors.New("task couldn't be updated")
	// ErrTaskNotDeleted ошибка: не удалось удалить задание в БД
	ErrTaskNotDeleted = errors.New("task couldn't be deleted")
	// ErrCouldNotGetTask ошибка: не удалось получить задание в БД
	ErrCouldNotGetTask = errors.New("couldn't get task")
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
// заглушка для ответа 400 без разглашения имплементации чувствительных процессов
var BadRequestResponse = ErrorResponse{
	Code:    http.StatusBadRequest,
	Message: "Ошибка запроса",
}

// GenericUnauthorizedResponse
// заглушка для ответа 401 без разглашения имплементации чувствительных процессов
var GenericUnauthorizedResponse = ErrorResponse{
	Code:    http.StatusUnauthorized,
	Message: "Ошибка авторизации",
}

// ForbiddenResponse
// заглушка для ответа 403 без разглашения имплементации чувствительных процессов
var ForbiddenResponse = ErrorResponse{
	Code:    http.StatusForbidden,
	Message: "Ошибка доступа",
}

// InternalServerErrorResponse
// заглушка для ответа 500 без разглашения имплементации чувствительных процессов
var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "Ошибка сервера",
}

// StatusConflictResponse
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
	ErrBoardNotCreated:        InternalServerErrorResponse,
	ErrBoardNotUpdated:        InternalServerErrorResponse,
	ErrBoardNotDeleted:        InternalServerErrorResponse,
	ErrCouldNotGetBoard:       InternalServerErrorResponse,
	ErrNoBoardAccess:          ForbiddenResponse,
	ErrTaskNotCreated:         InternalServerErrorResponse,
	ErrTaskNotUpdated:         InternalServerErrorResponse,
	ErrTaskNotDeleted:         InternalServerErrorResponse,
	ErrCouldNotGetTask:        InternalServerErrorResponse,
	ErrListNotCreated:         InternalServerErrorResponse,
	ErrListNotUpdated:         InternalServerErrorResponse,
	ErrListNotDeleted:         InternalServerErrorResponse,
}

func ErrorJSON(err ErrorResponse) []byte {
	response := dto.JSONResponse{
		Body: err,
	}
	jsonResponse, _ := json.Marshal(response)
	return jsonResponse
}

func ReturnError(err ErrorResponse, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(err.Code)
	response := ErrorJSON(err)
	_, _ = w.Write(response)
	r.Body.Close()
}
