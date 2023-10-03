package handlers

import (
	session "server/internal/config/session"
	"server/internal/service"
)

// HandlerManager
// объект со всеми хэндлерами приложения
type HandlerManager struct {
	AuthHandler
	BoardHandler
}

// NewHandlerManager
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlerManager(config session.SessionServerConfig,
	authService service.IAuthService,
	userAuthService service.IUserAuthService,
	//userUserService := user.NewUserService(userStorage),
	boardService service.IBoardService) *HandlerManager {
	return &HandlerManager{
		AuthHandler:  *NewAuthHandler(authService, userAuthService),
		BoardHandler: *NewBoardHandler(authService, boardService),
	}
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
