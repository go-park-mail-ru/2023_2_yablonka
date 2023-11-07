package handlers

import (
	config "server/internal/config/session"
	"server/internal/service"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
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
}

// NewHandlerManager
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlerManager(dbConnection *pgxpool.Pool, config *config.SessionServerConfig) *HandlerManager {
	userStorage := postgresql.NewUserStorage(dbConnection)
	authStorage := postgresql.NewAuthStorage(dbConnection)
	boardStorage := postgresql.NewBoardStorage(dbConnection)
	workspaceStorage := postgresql.NewWorkspaceStorage(dbConnection)

	userService := user.NewUserService(userStorage)
	authService := auth.NewAuthSessionService(*config, authStorage)
	boardService := board.NewBoardService(boardStorage)
	workspaceService := workspace.NewWorkspaceService(workspaceStorage)

	return &HandlerManager{
		UserHandler:      *NewUserHandler(authService, userService),
		BoardHandler:     *NewBoardHandler(authService, boardService),
		WorkspaceHandler: *NewWorkspaceHandler(workspaceService),
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

// NewBoardHandler
// возвращает BoardHandler с необходимыми сервисами
func NewWorkspaceHandler(ws service.IWorkspaceService) *WorkspaceHandler {
	return &WorkspaceHandler{
		ws: ws,
	}
}
