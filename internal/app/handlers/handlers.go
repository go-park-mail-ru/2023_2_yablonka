package handlers

import "server/internal/service"

func NewAuthHandler(as service.IAuthService, us service.IUserAuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
	}
}

func NewBoardHandler(bs service.IBoardService, as service.IAuthService) *BoardHandler {
	return &BoardHandler{
		authService:  as,
		boardService: bs,
	}
}
