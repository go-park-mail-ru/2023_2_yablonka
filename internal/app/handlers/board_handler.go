package handlers

import (
	"net/http"
	"server/internal/service"
)

type BoardHandler struct {
	as service.IAuthService
	bs service.IBoardService
}

func (bh BoardHandler) GetUserBoards(w http.ResponseWriter, r *http.Request) {

}
