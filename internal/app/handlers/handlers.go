package handlers

import (
	config "server/internal/config/session"
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
	UserHandler
	BoardHandler
	WorkspaceHandler
	ListHandler
	TaskHandler
}

// NewHandlerManager
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlerManager(dbConnection *pgxpool.Pool, config *config.SessionServerConfig) *HandlerManager {
	userStorage := postgresql.NewUserStorage(dbConnection)
	authStorage := postgresql.NewAuthStorage(dbConnection)
	boardStorage := postgresql.NewBoardStorage(dbConnection)
	workspaceStorage := postgresql.NewWorkspaceStorage(dbConnection)
	listStorage := postgresql.NewListStorage(dbConnection)
	taskStorage := postgresql.NewTaskStorage(dbConnection)

	userService := user.NewUserService(userStorage)
	authService := auth.NewAuthSessionService(*config, authStorage)
	boardService := board.NewBoardService(boardStorage)
	workspaceService := workspace.NewWorkspaceService(workspaceStorage)
	listService := list.NewListService(listStorage)
	taskService := task.NewTaskService(taskStorage)

	return &HandlerManager{
		UserHandler:      *NewUserHandler(authService, userService),
		BoardHandler:     *NewBoardHandler(authService, boardService),
		WorkspaceHandler: *NewWorkspaceHandler(workspaceService),
		ListHandler:      *NewListHandler(listService),
		TaskHandler:      *NewTaskHandler(taskService),
	}
}

// NewUserHandler
// возвращает AuthHandler с необходимыми сервисами
func NewUserHandler(as service.IAuthService, us service.IUserService) *UserHandler {
	return &UserHandler{
		as: as,
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
