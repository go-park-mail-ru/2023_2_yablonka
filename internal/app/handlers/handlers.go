package handlers

import (
	"net/http"
	"server/internal/service"
)

type IAuthHandler interface {
	LogIn(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
	// TODO VerifyAuth
	// TODO LogOut
}

func NewAuthHandler(as service.IAuthService, us service.IUserAuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
	}
}

type IBoardHandler interface {
	GetBoard(http.ResponseWriter, *http.Request)
}

func NewBoardHandler(bs service.IBoardService) *BoardHandler {
	return &BoardHandler{
		bs: bs,
	}
}
