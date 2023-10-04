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

//	@Summary Получить все доски пользователя
//	@Description И те, которые он создал и те, у которых он гость
//
//	@Accept  json
//	@Produce  json

//	@Success 200  body object{} true "Список объектов досок"
//	@Failure 400  {object}  error
//	@Failure 404  {object}  error
//	@Failure 500  {object}  error
//
// @Router /api/v1/user/boards [get]
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

	w.Write(jsonResponse)
	r.Body.Close()
}
