package handlers

import (
	"encoding/json"
	"fmt"
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

	ctx := r.Context()

	user, ok := ctx.Value("userObj").(*entities.User)
	if !ok {
		w.WriteHeader(apperrors.GenericUnauthorizedResponse.Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.GenericUnauthorizedResponse.Message),
		))
		return
	}

	userInfo := dto.VerifiedAuthInfo{
		UserID: user.ID,
	}

	ownedBoards, err := bh.bs.GetUserOwnedBoards(ctx, userInfo)
	if err != nil {
		w.WriteHeader(apperrors.ErrorMap[err].Code)
		w.Write([]byte(
			fmt.Sprintf(`{"body": {
				"error_response": "%s"
			}}`, apperrors.ErrorMap[err].Message)),
		)
		return
	}
	guestBoards, err := bh.bs.GetUserGuestBoards(ctx, userInfo)
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
	response := fmt.Sprintf(`{"body": {"user": %s, "boards": %s}}`, string(userJson), string(boardJson))

	w.Write([]byte(response))
}
