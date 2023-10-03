package handlers

import (
	"context"
	"encoding/json"
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
	rCtx := r.Context()
	var jsonResponse []byte

	user, ok := rCtx.Value("userObj").(*entities.User)
	if !ok {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.GenericUnauthorizedResponse))
		return
	}

	userInfo := dto.VerifiedAuthInfo{
		UserID: user.ID,
	}

	ownedBoards, err := bh.bs.GetUserOwnedBoards(rCtx, userInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}
	guestBoards, err := bh.bs.GetUserGuestBoards(rCtx, userInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.ErrorMap[err]))
		return
	}

	userBoards := dto.UserTotalBoardInfo{
		OwnedBoards: ownedBoards,
		GuestBoards: guestBoards,
	}
	userObj := rCtx.Value("userObj").(*entities.User)
	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user":   userObj,
			"boards": userBoards,
		},
	}
	jsonResponse, err = json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, "errorResponse", apperrors.InternalServerErrorResponse))
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}
