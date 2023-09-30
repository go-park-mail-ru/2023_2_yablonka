package handlers

import (
	"net/http"

	authservice "server/internal/service/auth"
	userservice "server/internal/service/user"
)

type IBoardHandler interface {
	GetBoard(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	// TODO VerifyAuth
	// TODO LogOut
}

// TODO IUserHandler

type BoardHandler struct {
	as authservice.IAuthService
	us userservice.IUserAuthService
}

func NewBoardHandler(as authservice.IAuthService, us userservice.IUserAuthService) BoardHandler {
	return BoardHandler{
		as: as,
		us: us,
	}
}
