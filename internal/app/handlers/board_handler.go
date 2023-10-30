package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service"
)

type BoardHandler struct {
	as service.IAuthService
	bs service.IBoardService
}

// @Summary Вывести все доски текущего пользователя
// @Description Выводит и созданные им доски и те, в которых он гость. Работает только для авторизированного пользователя.
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Success 200  {object}  doc_structs.UserBoardsResponse "Пользователь и его доски"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /api/v2/user/boards/ [get]
func (bh BoardHandler) GetUserBoards(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	var jsonResponse []byte

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	userInfo := dto.VerifiedAuthInfo{
		UserID: user.ID,
	}

	ownedBoards, err := bh.bs.GetUserOwnedBoards(rCtx, userInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	guestBoards, err := bh.bs.GetUserGuestBoards(rCtx, userInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	userBoards := dto.UserTotalBoardInfo{
		OwnedBoards: ownedBoards,
		GuestBoards: guestBoards,
	}
	userObj := rCtx.Value(dto.UserObjKey).(*entities.User)
	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user":   userObj,
			"boards": userBoards,
		},
	}
	jsonResponse, err = json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
}
