package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/google/uuid"
)

type ListHandler struct {
	ls service.IListService
}

// @Summary Создать список
// @Description Создать список
// @Tags lists
//
// @Accept  json
// @Produce  json
//
// @Param newListInfo body dto.NewListInfo true "данные нового списка"
//
// @Success 200  {object}  doc_structs.ListResponse "объект списка"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /list/create/ [post]
func (lh ListHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ListHandler.Create"
	errorMessage := "Creating list failed with error: "
	failBorder := "---------------------------------- Creating list FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Creating list ----------------------------------")

	var newListInfo dto.NewListInfo
	err := json.NewDecoder(r.Body).Decode(&newListInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	list, err := lh.ls.Create(rCtx, newListInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("list created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"list": list,
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

	logger.Info("---------------------------------- Creating list FAIL ----------------------------------")
}

// @Summary Обновить список
// @Description Обновить список
// @Tags lists
//
// @Accept  json
// @Produce  json
//
// @Param listInfo body dto.UpdatedListInfo true "обновленные данные списка"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /list/edit/ [post]
func (lh ListHandler) Update(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ListHandler.Update"
	errorMessage := "Updating failed with error: "
	failBorder := "---------------------------------- Updating list FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Updating list ----------------------------------")

	var listInfo dto.UpdatedListInfo
	err := json.NewDecoder(r.Body).Decode(&listInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = lh.ls.Update(rCtx, listInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("list updated", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Updating list SUCCESS ----------------------------------")
}

// @Summary Удалить список
// @Description Удалить список
// @Tags lists
//
// @Accept  json
// @Produce  json
//
// @Param listID body dto.ListID true "id списка"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /list/delete/ [delete]
func (lh ListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ListHandler.Delete"
	errorMessage := "Deleting failed with error: "
	failBorder := "---------------------------------- Deleting list FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Deleting list ----------------------------------")

	var listID dto.ListID
	err := json.NewDecoder(r.Body).Decode(&listID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = lh.ls.Delete(rCtx, listID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("list deleted", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Deleting list SUCCESS ----------------------------------")
}

// @Summary Обновить порядок списков
// @Description Поставить списки по порядку в запросе
// @Tags lists
//
// @Accept  json
// @Produce  json
//
// @Param listID body dto.ListIDs true "id списков"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /list/reorder/ [post]
func (lh ListHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ListHandler.UpdateOrder"
	errorMessage := "Updating order failed with error: "
	failBorder := "---------------------------------- ListHandler.UpdateOrder FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- ListHandler.UpdateOrder ----------------------------------")

	var listIDs dto.ListIDs
	err := json.NewDecoder(r.Body).Decode(&listIDs)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = lh.ls.UpdateOrder(rCtx, listIDs)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("list order changed", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- ListHandler.UpdateOrder SUCCESS ----------------------------------")
}
