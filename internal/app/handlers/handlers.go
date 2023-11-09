package handlers

import (
	config "server/internal/config"
	"server/internal/service"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	list "server/internal/service/list"
	task "server/internal/service/task"
	user "server/internal/service/user"
	workspace "server/internal/service/workspace"
	"server/internal/storage/postgresql"

	"github.com/jackc/pgx/v5/pgxpool"
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

	authService := auth.NewAuthService(config, authStorage)
	userService := user.NewUserService(userStorage)
	boardService := board.NewBoardService(boardStorage)
	workspaceService := workspace.NewWorkspaceService(workspaceStorage)
	listService := list.NewListService(listStorage)
	taskService := task.NewTaskService(taskStorage)

	return &HandlerManager{
		AuthHandler:      *NewAuthHandler(authService, userService),
		UserHandler:      *NewUserHandler(userService),
		BoardHandler:     *NewBoardHandler(authService, boardService),
		WorkspaceHandler: *NewWorkspaceHandler(workspaceService),
		ListHandler:      *NewListHandler(listService),
		TaskHandler:      *NewTaskHandler(taskService),
	}
}

// NewAuthHandler
// возвращает AuthHandler с необходимыми сервисами
func NewAuthHandler(as service.IAuthService, us service.IUserService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
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
