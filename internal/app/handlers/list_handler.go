package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"
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
	nodeName := "handler"
	errorMessage := "Creating list failed with error: "
	failBorder := "----------------- Creating list FAIL -----------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- Creating list -----------------")

	var newListInfo dto.NewListInfo
	err := json.NewDecoder(r.Body).Decode(&newListInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	logger.Debug("request struct decoded", funcName, nodeName)

	list, err := lh.ls.Create(rCtx, newListInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	logger.Debug("list created", funcName, nodeName)

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
	logger.Debug("response written", funcName, nodeName)

	logger.Info("----------------- Creating list FAIL -----------------")
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
	nodeName := "handler"
	errorMessage := "Updating failed with error: "
	failBorder := "----------------- Updating list FAIL -----------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- Updating list -----------------")

	var listInfo dto.UpdatedListInfo
	err := json.NewDecoder(r.Body).Decode(&listInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	logger.Debug("request struct decoded", funcName, nodeName)

	err = lh.ls.Update(rCtx, listInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	logger.Debug("list updated", funcName, nodeName)

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
	logger.Debug("response written", funcName, nodeName)

	logger.Info("----------------- Updating list SUCCESS -----------------")
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
	nodeName := "handler"
	errorMessage := "Deleting failed with error: "
	failBorder := "----------------- Deleting list FAIL -----------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- Deleting list -----------------")

	var listID uint64
	err := json.NewDecoder(r.Body).Decode(&listID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	logger.Debug("request struct decoded", funcName, nodeName)

	listIDObj := dto.ListID{
		Value: listID,
	}
	err = lh.ls.Delete(rCtx, listIDObj)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	logger.Debug("list deleted", funcName, nodeName)

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
	logger.Debug("response written", funcName, nodeName)

	logger.Info("----------------- Deleting list SUCCESS -----------------")
}
