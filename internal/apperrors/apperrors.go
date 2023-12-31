package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"server/internal/pkg/dto"
)

// Ошибки, связанные с конфигурацией сервера
var (
	// ErrEnvNotFound ошибка: переданный при конфигурации .env файл не был найден
	ErrEnvNotFound = errors.New("invalid .env location was provided to the server setup")
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

// Ошибки, связанные с GRPC
var (
	// ErrGRPCServerError ошибка: ошибка сервера
	ErrGRPCServerError = errors.New("GRPC server error")
)

// Ошибки, связанные с пользователями
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
	// ErrAvatarGone ошибка: удаление пустого аватара
	ErrAvatarGone = errors.New("deleting an empty avatar")
)

// Ошибки, связанные с CSAT
var (
	// ErrCouldNotCreateQuestion ошибка: не удалось создать ответ CSAT в БД
	ErrCouldNotCreateAnswer = errors.New("couldn't create answer")
	// ErrCouldNotCreateQuestion ошибка: не удалось создать вопрос CSAT в БД
	ErrCouldNotCreateQuestion = errors.New("couldn't create question")
	// ErrCouldNotGetQuestions ошибка: не удалось получить все вопросы CSAT в БД
	ErrCouldNotGetQuestions = errors.New("couldn't get all questions")
	// ErrCouldNotGetQuestionType ошибка: не удалось получить тип вопроса CSAT в БД
	ErrCouldNotGetQuestionType = errors.New("couldn't get question type")
	// ErrCouldNotGetQuestionType ошибка: не удалось обновить вопрос CSAT в БД
	ErrQuestionNotUpdated = errors.New("couldn't update question")
	// ErrCouldNotGetQuestionType ошибка: не удалось удалить вопрос CSAT в БД
	ErrQuestionNotDeleted = errors.New("couldn't delete question")
	// ErrAnswerRatingTooBig ошибка: у полученного ответа рейтинг выше доступного для вопроса
	ErrAnswerRatingTooBig = errors.New("provided rating is too big for this question type")
	// ErrCouldNotStoreAnswer ошибка: не удалось сохранить ответ
	ErrCouldNotStoreAnswer = errors.New("couldn't store answer")
)

// Ошибки, связанные с AuthService
var (
	// ErrTokenNotGenerated ошибка: не удалось сгенерировать токен
	ErrTokenNotGenerated = errors.New("failed to generate token")
	// ErrSessionExpired ошибка: время действия сессии истекло
	ErrSessionExpired = errors.New("user session has expired")
	// ErrSessionNotFound ошибка: полученной сессии нет в хранилище
	ErrSessionNotFound = errors.New("no session found for provided session ID")
	// ErrSessionNotCreated ошибка: полученной сессии нет в хранилище
	ErrSessionNotCreated = errors.New("session couldn't be created")
	// ErrCSRFExpired ошибка: время действия сессии истекло
	ErrCSRFExpired = errors.New("user CSRF has expired")
	// ErrCSRFNotFound ошибка: полученной сессии нет в хранилище
	ErrCSRFNotFound = errors.New("no CSRF found for provided CSRF ID")
	// ErrCSRFNotCreated ошибка: полученной сессии нет в хранилище
	ErrCSRFNotCreated = errors.New("CSRF couldn't be created")
)

// Ошибки, связанные с Commentervice
var (
	// ErrCouldNotGetTaskComments ошибка: нельзя получить комментарии задания
	ErrCouldNotGetTaskComments = errors.New("couldn't get the comments from task")
	// ErrCouldNotGetComments ошибка: нельзя получить комментарии
	ErrCouldNotGetComments = errors.New("couldn't get the comments")
)

// Ошибки, связанные с БД
var (
	// ErrCouldNotBuildQuery ошибка: не удалось сформировать SQL запрос
	ErrCouldNotBuildQuery = errors.New("error building an SQL query")
	// ErrCouldNotExecuteQuery ошибка: не удалось выполнить SQL запрос
	ErrCouldNotExecuteQuery = errors.New("error executing an SQL query")
	// ErrCouldNotExecuteQuery ошибка: не удалось закрыть запрос
	ErrCouldNotCloseQuery = errors.New("error closing SQL query")
	// ErrCouldNotBeginTransaction ошибка: не удалось начать транзакцию
	ErrCouldNotBeginTransaction = errors.New("error starting a transaction")
	// ErrCouldNotCollectRows ошибка: не удалось получить строки
	ErrCouldNotCollectRows = errors.New("couldn't collect rows")
	// ErrCouldNotScanRows ошибка: не удалось отсканировать строки
	ErrCouldNotScanRows = errors.New("couldn't scan rows")
	// ErrCouldNotRollback ошибка: не удалось отменить транзакцию
	ErrCouldNotRollback = errors.New("couldn't rollback a transaction")
	// ErrCouldNotCommit ошибка: не удалось коммитнуть изменения в бд
	ErrCouldNotCommit = errors.New("failed to commit database changes")
)

// Ошибки, связанные с сервером
var (
	// ErrFailedToCreateFile ошибка: не удалось создать файл
	ErrFailedToCreateFile = errors.New("failed to create file")
	// ErrFailedToSaveFile ошибка: не удалось сохранить файл
	ErrFailedToSaveFile = errors.New("failed to save file")
	// ErrFailedToDeleteFile ошибка: не удалось удалить файл
	ErrFailedToDeleteFile = errors.New("failed to delete file")
	// ErrNoLoggerFound ошибка: не удалось получить логгер из контекста
	ErrNoLoggerFound = errors.New("no logger in context")
	// ErrNoRequestIDFound ошибка: не удалось получить id запроса из контекста
	ErrNoRequestIDFound = errors.New("no request id in context")
)

// Ошибки, связанные с BoardService
var (
	// ErrWorkspaceNotDeleted ошибка: не удалось создать рабочее прострнство в БД
	ErrBoardNotCreated = errors.New("board couldn't be created")
	// ErrUserAlreadyInBoard ошибка: пользователь уже есть в доске
	ErrUserAlreadyInBoard = errors.New("user already in board")
	// ErrUserNotInBoard ошибка: пользователя нет в доске
	ErrUserNotInBoard = errors.New("user not in board")
	// ErrWorkspaceNotDeleted ошибка: не удалось получить рабочее прострнство в БД
	ErrBoardNotUpdated = errors.New("board couldn't be updated")
	// ErrWorkspaceNotDeleted ошибка: не удалось удалить рабочее прострнство в БД
	ErrBoardNotDeleted = errors.New("board couldn't be deleted")
	// ErrCouldNotGetBoard ошибка: доски с полученным ID не существует
	ErrCouldNotGetBoard = errors.New("could not retrieve board")
	// ErrCouldNotGetBoardUsers ошибка: пользователей у доски с полученным ID не существует
	ErrCouldNotGetBoardUsers = errors.New("could not retrieve board users")
	// ErrCouldNotGetBoardTags ошибка: не удалось полить тэги доски
	ErrCouldNotGetBoardTags = errors.New("could not retrieve board tags")
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
	// ErrCouldNotGetWorkspace ошибка: не удалось получить рабочее прострнство в БД
	ErrCouldNotGetWorkspace = errors.New("workspace couldn't be retreived")
	// ErrWorkspaceNotDeleted ошибка: не удалось удалить рабочее прострнство в БД
	ErrWorkspaceNotDeleted = errors.New("user couldn't be deleted")
)

// Ошибки, связанные с TagService
var (
	// ErrTagNotCreated ошибка: не удалось создать тэг
	ErrTagNotCreated = errors.New("tag couldn't be created")
	// ErrTagNotCreated ошибка: не удалось обновить тэг
	ErrTagNotUpdated = errors.New("tag couldn't be updated")
	// ErrTagNotAddedToTask ошибка: не удалось добавить тэг к заданию
	ErrTagNotAddedToTask = errors.New("tag couldn't be added to task")
	// ErrTagNotRemovedFromTask ошибка: не удалось удалить тэг из задания
	ErrTagNotRemovedFromTask = errors.New("tag couldn't be removed from task")
	// ErrBoardTagConnectionNotCreated ошибка: не удалось создать связь между тэгом и доской
	ErrBoardTagConnectionNotCreated = errors.New("could not create connection between tag and board")
	// ErrTaskTagConnectionNotCreated ошибка: не удалось создать связь между тэгом и заданием
	ErrTaskTagConnectionNotCreated = errors.New("could not create connection between tag and task")
)

// Ошибки, связанные с ListService
var (
	// ErrListNotDeleted ошибка: не удалось получить список из БД
	ErrCouldNotGetList = errors.New("could not get list")
	// ErrListNotDeleted ошибка: не удалось изменить порядок в списке
	ErrCouldNotChangeListOrder = errors.New("could not change list order")
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
	// ErrListNotDeleted ошибка: не удалось изменить порядок в списке
	ErrCouldNotChangeTaskOrder = errors.New("could not change task order")
	// ErrTaskNotCreated ошибка: не удалось создать задание в БД
	ErrTaskNotCreated = errors.New("task couldn't be created")
	// ErrTaskNotUpdated ошибка: не удалось получить задание в БД
	ErrTaskNotUpdated = errors.New("task couldn't be updated")
	// ErrTaskNotDeleted ошибка: не удалось удалить задание в БД
	ErrTaskNotDeleted = errors.New("task couldn't be deleted")
	// ErrCouldNotGetTask ошибка: не удалось получить задание в БД
	ErrCouldNotGetTask = errors.New("couldn't get task")
	// ErrCouldNotGetTaskFiles ошибка: не удалось получить задание в БД
	ErrCouldNotGetTaskFiles = errors.New("couldn't get task files")
	// ErrCouldNotAddTaskUser ошибка: не удалось добавить пользователя на карточку
	ErrCouldNotAddTaskUser = errors.New("couldn't add user to task")
	// ErrCouldNotRemoveTaskUser ошибка: не удалось удалить пользователя с карточки
	ErrCouldNotRemoveTaskUser = errors.New("couldn't remove user from task")
	// ErrUserAlreadyInTask ошибка: пользователь уже есть в задании
	ErrUserAlreadyInTask = errors.New("user already in task")
	// ErrUserNotInTask ошибка: пользователь уже есть в задании
	ErrUserNotInTask = errors.New("user not in task")
)

// Ошибки, связанные с CommentService
var (
	// ErrCommentNotCreated ошибка: не удалось создать элемент чеклиста в БД
	ErrCommentNotCreated = errors.New("comment couldn't be created")
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

// GoneResponse
// заглушка для ответа 410 без разглашения имплементации
var GoneResponse = ErrorResponse{
	Code:    http.StatusGone,
	Message: "Ресурс уже был удалён",
}

// ErrorMap
// карта для связи ошибок приложения и ответа бэкэнд-сервера
var ErrorMap = map[error]ErrorResponse{
	ErrGRPCServerError:              InternalServerErrorResponse,
	ErrUserNotFound:                 WrongLoginResponse,
	ErrWrongPassword:                WrongLoginResponse,
	ErrUserAlreadyExists:            StatusConflictResponse,
	ErrUserNotCreated:               InternalServerErrorResponse,
	ErrUserNotUpdated:               InternalServerErrorResponse,
	ErrTokenNotGenerated:            InternalServerErrorResponse,
	ErrCSRFNotFound:                 GenericUnauthorizedResponse,
	ErrCSRFNotCreated:               InternalServerErrorResponse,
	ErrCSRFExpired:                  GenericUnauthorizedResponse,
	ErrSessionDurationMissing:       InternalServerErrorResponse,
	ErrSessionNullDuration:          InternalServerErrorResponse,
	ErrSessionIDLengthMissing:       InternalServerErrorResponse,
	ErrSessionNullIDLength:          InternalServerErrorResponse,
	ErrSessionNotCreated:            InternalServerErrorResponse,
	ErrSessionExpired:               GenericUnauthorizedResponse,
	ErrSessionNotCreated:            InternalServerErrorResponse,
	ErrSessionNotFound:              GenericUnauthorizedResponse,
	ErrWorkspaceNotCreated:          InternalServerErrorResponse,
	ErrCouldNotGetWorkspace:         InternalServerErrorResponse,
	ErrWorkspaceNotDeleted:          InternalServerErrorResponse,
	ErrBoardNotCreated:              InternalServerErrorResponse,
	ErrBoardNotUpdated:              InternalServerErrorResponse,
	ErrBoardNotDeleted:              InternalServerErrorResponse,
	ErrCouldNotGetBoard:             InternalServerErrorResponse,
	ErrNoBoardAccess:                ForbiddenResponse,
	ErrCouldNotAddBoardUser:         InternalServerErrorResponse,
	ErrCouldNotRemoveBoardUser:      InternalServerErrorResponse,
	ErrCouldNotAddTaskUser:          InternalServerErrorResponse,
	ErrCouldNotRemoveTaskUser:       InternalServerErrorResponse,
	ErrTaskNotCreated:               InternalServerErrorResponse,
	ErrTaskNotUpdated:               InternalServerErrorResponse,
	ErrTaskNotDeleted:               InternalServerErrorResponse,
	ErrCouldNotGetTask:              InternalServerErrorResponse,
	ErrCouldNotGetTaskFiles:         InternalServerErrorResponse,
	ErrListNotCreated:               InternalServerErrorResponse,
	ErrListNotUpdated:               InternalServerErrorResponse,
	ErrListNotDeleted:               InternalServerErrorResponse,
	ErrCouldNotGetChecklist:         InternalServerErrorResponse,
	ErrChecklistNotCreated:          InternalServerErrorResponse,
	ErrChecklistNotUpdated:          InternalServerErrorResponse,
	ErrChecklistNotDeleted:          InternalServerErrorResponse,
	ErrCouldNotGetChecklistItem:     InternalServerErrorResponse,
	ErrChecklistItemNotCreated:      InternalServerErrorResponse,
	ErrChecklistItemNotUpdated:      InternalServerErrorResponse,
	ErrChecklistItemNotDeleted:      InternalServerErrorResponse,
	ErrCommentNotCreated:            InternalServerErrorResponse,
	ErrUserAlreadyInBoard:           StatusConflictResponse,
	ErrUserNotInBoard:               StatusConflictResponse,
	ErrUserAlreadyInTask:            StatusConflictResponse,
	ErrUserNotInTask:                StatusConflictResponse,
	ErrFailedToCreateFile:           InternalServerErrorResponse,
	ErrFailedToSaveFile:             InternalServerErrorResponse,
	ErrFailedToDeleteFile:           InternalServerErrorResponse,
	ErrAnswerRatingTooBig:           BadRequestResponse,
	ErrCouldNotChangeTaskOrder:      BadRequestResponse,
	ErrCouldNotChangeListOrder:      BadRequestResponse,
	ErrCouldNotStoreAnswer:          InternalServerErrorResponse,
	ErrCouldNotGetQuestions:         InternalServerErrorResponse,
	ErrCouldNotCreateQuestion:       InternalServerErrorResponse,
	ErrQuestionNotUpdated:           InternalServerErrorResponse,
	ErrAvatarGone:                   GoneResponse,
	ErrCouldNotCreateAnswer:         InternalServerErrorResponse,
	ErrCouldNotGetComments:          InternalServerErrorResponse,
	ErrTagNotCreated:                InternalServerErrorResponse,
	ErrBoardTagConnectionNotCreated: InternalServerErrorResponse,
	ErrTaskTagConnectionNotCreated:  InternalServerErrorResponse,
	ErrCouldNotGetBoardTags:         InternalServerErrorResponse,
	ErrCouldNotBuildQuery:           InternalServerErrorResponse,
	ErrCouldNotExecuteQuery:         InternalServerErrorResponse,
	ErrCouldNotCloseQuery:           InternalServerErrorResponse,
	ErrCouldNotBeginTransaction:     InternalServerErrorResponse,
	ErrCouldNotCollectRows:          InternalServerErrorResponse,
	ErrCouldNotScanRows:             InternalServerErrorResponse,
	ErrCouldNotRollback:             InternalServerErrorResponse,
	ErrCouldNotCommit:               InternalServerErrorResponse,
	ErrNoLoggerFound:                InternalServerErrorResponse,
}

func ErrorJSON(err ErrorResponse) []byte {
	response := dto.JSONResponse{
		Body: err,
	}
	jsonResponse, _ := json.Marshal(response)
	return jsonResponse
}

func ReturnError(err ErrorResponse, w http.ResponseWriter, r *http.Request) {
	log.Println("Returning error response", err)
	w.WriteHeader(err.Code)
	log.Println("Wrote header, code:", err.Code)
	response := ErrorJSON(err)
	_, _ = w.Write(response)
	log.Println("Wrote response, text:", err.Message)
	r.Body.Close()
	log.Println("Request body closed")
}
