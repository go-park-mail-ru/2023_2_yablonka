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

	"github.com/asaskevich/govalidator"
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

	userId := dto.UserID{
		Value: user.ID,
	}

	ownedBoards, err := bh.bs.GetUserOwnedBoards(rCtx, userId)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	guestBoards, err := bh.bs.GetUserGuestBoards(rCtx, userId)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user": rCtx.Value(dto.UserObjKey).(*entities.User),
			"boards": dto.UserTotalBoardInfo{
				OwnedBoards: ownedBoards,
				GuestBoards: guestBoards,
			},
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

func (bh BoardHandler) GetFullBoard(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var boardRequest dto.IndividualBoardRequest
	err := json.NewDecoder(r.Body).Decode(&boardRequest)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(boardRequest)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	board, err := bh.bs.GetBoardWithListsAndTasks(rCtx, dto.BoardID{Value: boardRequest.BoardID})
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"board": board,
		},
	}

	jsonResponse, err := json.Marshal(response)
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
