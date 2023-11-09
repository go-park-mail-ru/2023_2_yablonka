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

type WorkspaceHandler struct {
	ws service.IWorkspaceService
}

// @Summary Вывести все рабочие пространства и доски текущего пользователя
// @Description Выводит и созданные им рабочие пространства и те, в которых он гость. Работает только для авторизированного пользователя.
// @Tags workspaces
//
// @Accept  json
// @Produce  json
//
// @Success 200  {object}  doc_structs.AllWorkspacesResponse "Пользователь и его рабочие пространства"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /user/workspaces/ [get]
func (wh WorkspaceHandler) GetUserWorkspaces(w http.ResponseWriter, r *http.Request) {
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

	workspaces, err := wh.ws.GetUserWorkspaces(rCtx, userId)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"user":       user,
			"workspaces": workspaces,
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

// @Summary Создать рабочее пространство
// @Description Создать рабочее пространство
// @Tags workspaces
//
// @Accept  json
// @Produce  json
//
// @Param newWorkspaceInfo body dto.NewWorkspaceInfo true "данные нового рабочего пространства"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /workspace/create/ [post]
func (wh WorkspaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var newWorkspaceInfo dto.NewWorkspaceInfo
	err := json.NewDecoder(r.Body).Decode(&newWorkspaceInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(newWorkspaceInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	workspace, err := wh.ws.Create(rCtx, newWorkspaceInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
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

// @Summary Обновить рабочее пространство
// @Description Обновить рабочее пространство
// @Tags workspaces
//
// @Accept  json
// @Produce  json
//
// @Param workspaceInfo body dto.UpdatedWorkspaceInfo true "обновленные данные рабочего пространства"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /workspace/update/ [post]
func (wh WorkspaceHandler) UpdateData(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var workspaceInfo dto.UpdatedWorkspaceInfo
	err := json.NewDecoder(r.Body).Decode(&workspaceInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(workspaceInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	err = wh.ws.UpdateData(rCtx, workspaceInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
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

// @Summary Обновить гостей рабочего пространства
// @Description Обновить гостей рабочего пространства
// @Tags workspaces
//
// @Accept  json
// @Produce  json
//
// @Param authData body dto.UpdatedWorkspaceInfo true "обновленный список пользователей"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /workspace/update/change_users/ [post]
func (wh WorkspaceHandler) ChangeGuests(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var guestsInfo dto.ChangeWorkspaceGuestsInfo
	err := json.NewDecoder(r.Body).Decode(&guestsInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(guestsInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	err = wh.ws.UpdateUsers(rCtx, guestsInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
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

// @Summary Удалить рабочее пространство
// @Description Удалить рабочее пространство
// @Tags workspaces
//
// @Accept  json
// @Produce  json
//
// @Param workspaceID body dto.WorkspaceID true "id рабочего пространства"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /workspace/delete/ [delete]
func (wh WorkspaceHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = wh.ws.Delete(rCtx, workspaceID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
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
