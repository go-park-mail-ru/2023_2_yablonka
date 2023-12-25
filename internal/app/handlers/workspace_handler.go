package handlers

import (
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
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
// @Router /user/workspaces [get]
func (wh WorkspaceHandler) GetUserWorkspaces(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "WorkspaceHandler.GetUserWorkspaces"
	errorMessage := "Getting user workspaces failed with error: "
	failBorder := "---------------------------------- Getting user workspaces FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Getting user workspaces ----------------------------------")

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User found", requestID.String(), funcName, nodeName)

	userId := dto.UserID{
		Value: user.ID,
	}
	workspaces, err := wh.ws.GetUserWorkspaces(rCtx, userId)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User workspaces received", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"workspaces": workspaces,
		},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- Getting user workspaces SUCCESS ----------------------------------")
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
// @Success 200  {object}  doc_structs.AllWorkspacesResponse "Пользователь и его рабочие пространства"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /workspace/ [post]
func (wh WorkspaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "WorkspaceHandler.Create"
	errorMessage := "Creating workspace failed with error: "
	failBorder := "---------------------------------- Creating workspace FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Creating workspace ----------------------------------")

	var newWorkspaceInfo dto.NewWorkspaceInfo
	err := easyjson.UnmarshalFromReader(r.Body, &newWorkspaceInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON parsed", requestID.String(), funcName, nodeName)

	workspace, err := wh.ws.Create(rCtx, newWorkspaceInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Workspace created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"workspace": workspace,
		},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- Creating workspace SUCCESS ----------------------------------")
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
// @Router /workspace/{workspaceID}/ [put]
func (wh WorkspaceHandler) UpdateData(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "WorkspaceHandler.UpdateData"
	errorMessage := "Updating workspace data failed with error: "
	failBorder := "---------------------------------- Updating workspace data FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Updating workspace data ----------------------------------")

	var workspaceInfo dto.UpdatedWorkspaceInfo
	err := easyjson.UnmarshalFromReader(r.Body, &workspaceInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON parsed", requestID.String(), funcName, nodeName)

	workspaceInfo.ID = rCtx.Value(dto.WorkspaceIDKey).(uint64)

	err = wh.ws.UpdateData(rCtx, workspaceInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("workspace data updated", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- Updating workspace data SUCCESS ----------------------------------")
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
// @Router /workspace/{workspaceID}/delete/ [delete]
func (wh WorkspaceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "WorkspaceHandler.Delete"
	errorMessage := "Deleting workspace failed with error: "
	failBorder := "---------------------------------- Deleting workspace FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Deleting workspace ----------------------------------")

	var workspaceID dto.WorkspaceID
	workspaceID.Value = rCtx.Value(dto.WorkspaceIDKey).(uint64)

	err := wh.ws.Delete(rCtx, workspaceID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Workspace deleted", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- Deleting workspace SUCCESS ----------------------------------")
}
