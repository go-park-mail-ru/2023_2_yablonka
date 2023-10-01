package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"server/internal/service"
)

type BoardHandler struct {
	authService  service.IAuthService
	boardService service.IBoardService
}

func (bh BoardHandler) GetUserBoards(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	cookie, err := r.Cookie("tabula_user")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	token := cookie.Value

	user, err := bh.authService.VerifyAuth(ctx, token)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	boards, err := bh.boardService.GetUserBoards(ctx, *user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(boards)

	w.Write([]byte(json))
}
