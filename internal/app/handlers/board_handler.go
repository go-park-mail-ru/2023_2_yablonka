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
	as service.IAuthService
	bs service.IBoardService
}

// TODO change the default data
// @Summary Log user into the system
// @Description Create new session or continue old one
// @ID login
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} nil
// @Router /api/v1/users/{id} [get]
func (bh BoardHandler) GetUserBoards(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// remake with r.Context
	// add timeout
	cookie, err := r.Cookie("tabula_user")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	token := cookie.Value

	user, err := bh.as.VerifyAuth(ctx, token)

	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}

	ownedBoards, err := bh.bs.GetUserOwnedBoards(ctx, *user)
	if err != nil {
		http.Error(w, apperrors.ErrorMap[err].Message, apperrors.ErrorMap[err].Code)
		return
	}
	guestBoards, err := bh.bs.GetUserGuestBoards(ctx, *user)
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
