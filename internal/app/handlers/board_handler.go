package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
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
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	token := cookie.Value

	user, err := bh.authService.VerifyAuth(ctx, token)

	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	ownedBoards, err := bh.boardService.GetUserOwnedBoards(ctx, *user)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}
	guestBoards, err := bh.boardService.GetUserGuestBoards(ctx, *user)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}
	userBoards := dto.UserTotalBoardInfo{
		OwnedBoards: ownedBoards,
		GuestBoards: guestBoards,
	}

	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	json, _ := json.Marshal(userBoards)
	response := fmt.Sprintf(`{"body": %s}`, string(json))

	w.Write([]byte(response))
}
