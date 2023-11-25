package handlers

import (
	"server/internal/service"

	"github.com/sirupsen/logrus"
)

// Handlers
// объект со всеми хэндлерами приложения
type Handlers struct {
	AuthHandler
	UserHandler
	BoardHandler
	WorkspaceHandler
	ListHandler
	TaskHandler
	CSATAnswerHandler
	CSATQuestionHandler
}

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		AuthHandler:         *NewAuthHandler(services.Auth, services.User, services.CSRF),
		UserHandler:         *NewUserHandler(services.User),
		BoardHandler:        *NewBoardHandler(services.Auth, services.Board),
		WorkspaceHandler:    *NewWorkspaceHandler(services.Workspace),
		ListHandler:         *NewListHandler(services.List),
		TaskHandler:         *NewTaskHandler(services.Task),
		CSATAnswerHandler:   *NewCSATAnswerHandler(services.CSATAnswer, services.CSATQuestion),
		CSATQuestionHandler: *NewCSATQuestionHandler(services.CSATQuestion),
	}
}

// NewAuthHandler
// возвращает AuthHandler с необходимыми сервисами
func NewAuthHandler(as service.IAuthService, us service.IUserService, cs service.ICSRFService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
		cs: cs,
	}
}

// NewCSATHandler
// возвращает CSATHandler с необходимыми сервисами
func NewCSATQuestionHandler(qs service.ICSATQuestionService) *CSATQuestionHandler {
	return &CSATQuestionHandler{
		qs: qs,
	}
}

// NewCSATHandler
// возвращает CSATHandler с необходимыми сервисами
func NewCSATAnswerHandler(as service.ICSATSAnswerService, qs service.ICSATQuestionService) *CSATAnswerHandler {
	return &CSATAnswerHandler{
		as: as,
		qs: qs,
	}
}

// NewUserHandler
// возвращает UserHandler с необходимыми сервисами
func NewUserHandler(us service.IUserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

// NewBoardHandler
// возвращает BoardHandler с необходимыми сервисами
func NewBoardHandler(as service.IAuthService, bs service.IBoardService) *BoardHandler {
	return &BoardHandler{
		as: as,
		bs: bs,
	}
}

// NewWorkspaceHandler
// возвращает WorkspaceHandler с необходимыми сервисами
func NewWorkspaceHandler(ws service.IWorkspaceService) *WorkspaceHandler {
	return &WorkspaceHandler{
		ws: ws,
	}
}

// NewListHandler
// возвращает ListHandler с необходимыми сервисами
func NewListHandler(ls service.IListService) *ListHandler {
	return &ListHandler{
		ls: ls,
	}
}

// NewTaskHandler
// возвращает TaskHandler с необходимыми сервисами
func NewTaskHandler(ts service.ITaskService) *TaskHandler {
	return &TaskHandler{
		ts: ts,
	}
}

func handlerDebugLog(logger *logrus.Logger, function string, message string) {
	logger.
		WithFields(logrus.Fields{
			"route_node": "handler",
			"function":   function,
		}).
		Debug(message)
}
