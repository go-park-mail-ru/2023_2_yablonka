package handlers

import (
	"net/http"

	"server/internal/service"
)

type IBoardHandler interface {
	GetBoard(w http.ResponseWriter, r *http.Request)
	// TODO VerifyAuth
	// TODO LogOut
}

// TODO IUserHandler

type BoardHandler struct {
	bs service.IBoardService
}

func NewBoardHandler(bs service.IBoardService) *BoardHandler {
	return &BoardHandler{
		bs: bs,
	}
}
