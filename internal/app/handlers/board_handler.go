package handlers

import (
	"net/http"

	boardservice "server/internal/service/board"
)

type IBoardHandler interface {
	GetBoard(w http.ResponseWriter, r *http.Request)
	// TODO VerifyAuth
	// TODO LogOut
}

// TODO IUserHandler

type BoardHandler struct {
	bs boardservice.IBoardService
}

func NewBoardHandler(bs boardservice.IBoardService) *BoardHandler {
	return &BoardHandler{
		bs: bs,
	}
}
