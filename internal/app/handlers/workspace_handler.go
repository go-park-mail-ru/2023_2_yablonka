package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/asaskevich/govalidator"
)

type WorkspaceHandler struct {
	ws service.IWorkspaceService
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
func (wh WorkspaceHandler) GetWorkspace(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var workspaceID dto.WorkspaceID
	err := json.NewDecoder(r.Body).Decode(&workspaceID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(workspaceID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	workspace, err := wh.ws.GetWorkspace(rCtx, workspaceID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user":      "",
			"workspace": workspace,
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
