package handlers

import (
	session "server/internal/config/session"
	"server/internal/service"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	user "server/internal/service/user"
	"server/internal/storage/in_memory"
)

// HandlerManager
// объект со всеми хэндлерами приложения
type HandlerManager struct {
	AuthHandler  AuthHandler
	BoardHandler BoardHandler
}

// NewAuthHandler
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlerManager(config session.SessionServerConfig) (*HandlerManager, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}
	userStorage := in_memory.NewUserStorage()
	authStorage := in_memory.NewAuthStorage()
	boardStorage := in_memory.NewBoardStorage()

	authService := auth.NewAuthSessionService(config, authStorage)
	userAuthService := user.NewAuthUserService(userStorage)
	//userUserService := user.NewUserService(userStorage)
	boardService := board.NewBoardService(boardStorage)

	var manager HandlerManager
	manager.AuthHandler = *NewAuthHandler(authService, userAuthService)
	manager.BoardHandler = *NewBoardHandler(authService, boardService)
	return &manager, nil
}

// NewAuthHandler
// возвращает AuthHandler с необходимыми сервисами
func NewAuthHandler(as service.IAuthService, us service.IUserAuthService) *AuthHandler {
	return &AuthHandler{
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
