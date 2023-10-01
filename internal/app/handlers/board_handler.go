package handlers

import (
	"server/internal/service"
)

type BoardHandler struct {
	bs service.IBoardService
}
