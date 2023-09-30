package handlers

import (
	"net/http"

	"server/internal/service"
)

type IBoardHandler interface {
	GetBoard(w http.ResponseWriter, r *http.Request)
}

type BoardHandler struct {
	bs service.IBoardService
}

func NewBoardHandler(bs service.IBoardService) *BoardHandler {
	return &BoardHandler{
		bs: bs,
	}
}
