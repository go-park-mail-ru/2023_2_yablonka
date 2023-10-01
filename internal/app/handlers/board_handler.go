package handlers

import (
	"server/internal/service"
)

type BoardHandler struct {
	bs service.IBoardService
}

func NewBoardHandler(bs service.IBoardService) *BoardHandler {
	return &BoardHandler{
		bs: bs,
	}
}
