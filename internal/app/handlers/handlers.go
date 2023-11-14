package handlers

import (
	config "server/internal/config"
	"server/internal/service"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	csrf "server/internal/service/csrf"
	list "server/internal/service/list"
	task "server/internal/service/task"
	user "server/internal/service/user"
	workspace "server/internal/service/workspace"
	"server/internal/storage/postgresql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// HandlerManager
// объект со всеми хэндлерами приложения
type HandlerManager struct {
	AuthHandler
	UserHandler
	BoardHandler
	WorkspaceHandler
	ListHandler
	TaskHandler
}

// NewHandlerManager
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlerManager(dbConnection *pgxpool.Pool, config config.SessionConfig) *HandlerManager {
	userStorage := postgresql.NewUserStorage(dbConnection)
	authStorage := postgresql.NewAuthStorage(dbConnection)
	boardStorage := postgresql.NewBoardStorage(dbConnection)
	workspaceStorage := postgresql.NewWorkspaceStorage(dbConnection)
	listStorage := postgresql.NewListStorage(dbConnection)
	taskStorage := postgresql.NewTaskStorage(dbConnection)
	csrfStorage := postgresql.NewCSRFStorage(dbConnection)

	authService := auth.NewAuthService(config, authStorage)
	userService := user.NewUserService(userStorage)
	boardService := board.NewBoardService(boardStorage)
	workspaceService := workspace.NewWorkspaceService(workspaceStorage)
	listService := list.NewListService(listStorage)
	taskService := task.NewTaskService(taskStorage)
	csrfService := csrf.NewCSRFService(config, csrfStorage)

	return &HandlerManager{
		AuthHandler:      *NewAuthHandler(authService, userService, csrfService),
		UserHandler:      *NewUserHandler(userService),
		BoardHandler:     *NewBoardHandler(authService, boardService),
		WorkspaceHandler: *NewWorkspaceHandler(workspaceService),
		ListHandler:      *NewListHandler(listService),
		TaskHandler:      *NewTaskHandler(taskService),
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
