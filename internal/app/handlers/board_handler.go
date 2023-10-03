package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
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
		log.Println("User is missing the cookie")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, http.StatusText(http.StatusUnauthorized))),
		)
		return
	}

	ctx := context.Background()
	token := cookie.Value

	user, err := bh.as.VerifyAuth(ctx, token)

	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message)),
		)
		return
	}

	ownedBoards, err := bh.bs.GetUserOwnedBoards(ctx, *user)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message)),
		)
		return
	}
	guestBoards, err := bh.bs.GetUserGuestBoards(ctx, *user)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message)),
		)
		return
	}

	userBoards := dto.UserTotalBoardInfo{
		OwnedBoards: ownedBoards,
		GuestBoards: guestBoards,
	}

	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message)),
		)
		return
	}

	userObj := r.Context().Value("userObj").(*entities.User)
	userJson, _ := json.Marshal(&userObj)
	boardJson, _ := json.Marshal(userBoards)
	response := fmt.Sprintf(`{"body": "user": %s, "boards": %s}`, string(userJson), string(boardJson))

	w.Write([]byte(response))
}
