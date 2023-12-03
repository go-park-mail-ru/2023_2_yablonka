package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"
)

type ChecklistItemHandler struct {
	clis service.IChecklistItemService
}

// @Summary Создать элемент чеклиста
// @Description Создать элемент чеклиста
// @Tags checklistItems
//
// @Accept  json
// @Produce  json
//
// @Param newChecklistItemInfo body dto.NewChecklistItemInfo true "данные нового элемента чеклистаа"
//
// @Success 200  {object}  doc_structs.ChecklistItemResponse "объект элемента чеклиста"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /checklist/item/create/ [post]
func (clh ChecklistItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChecklistItemHandler.Create"
	nodeName := "handler"
	errorMessage := "Creating ChecklistItem failed with error: "
	failBorder := "---------------------------------- Creating ChecklistItem FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Creating ChecklistItem ----------------------------------")

	var newChecklistItemInfo dto.NewChecklistItemInfo
	err := json.NewDecoder(r.Body).Decode(&newChecklistItemInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("request struct decoded", funcName, nodeName)

	ChecklistItem, err := clh.clis.Create(rCtx, newChecklistItemInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("ChecklistItem created", funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"checklistItem": ChecklistItem,
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

	logger.Info("---------------------------------- Creating ChecklistItem FAIL ----------------------------------")
}

// @Summary Обновить элемент чеклиста
// @Description Обновить элемент чеклиста
// @Tags checklistItems
//
// @Accept  json
// @Produce  json
//
// @Param ChecklistItemInfo body dto.UpdatedChecklistItemInfo true "обновленные данные элемент чеклистаа"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /checklist/item/update/ [post]
func (clh ChecklistItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChecklistItemHandler.Update"
	nodeName := "handler"
	errorMessage := "Updating failed with error: "
	failBorder := "---------------------------------- Updating ChecklistItem FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Updating ChecklistItem ----------------------------------")

	var ChecklistItemInfo dto.UpdatedChecklistItemInfo
	err := json.NewDecoder(r.Body).Decode(&ChecklistItemInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("request struct decoded", funcName, nodeName)

	err = clh.clis.Update(rCtx, ChecklistItemInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("ChecklistItem updated", funcName, nodeName)

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

	logger.Info("---------------------------------- Updating ChecklistItem SUCCESS ----------------------------------")
}

// @Summary Удалить элемент чеклиста
// @Description Удалить элемент чеклиста
// @Tags checklistItems
//
// @Accept  json
// @Produce  json
//
// @Param ChecklistItemID body dto.ChecklistItemID true "id элемента чеклиста"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /checklist/item/delete/ [delete]
func (clh ChecklistItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChecklistItemHandler.Delete"
	nodeName := "handler"
	errorMessage := "Deleting failed with error: "
	failBorder := "---------------------------------- Deleting ChecklistItem FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Deleting ChecklistItem ----------------------------------")

	var checklistItemID dto.ChecklistItemID
	err := json.NewDecoder(r.Body).Decode(&checklistItemID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("request struct decoded", funcName, nodeName)

	err = clh.clis.Delete(rCtx, checklistItemID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("ChecklistItem deleted", funcName, nodeName)

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

	logger.Info("---------------------------------- Deleting ChecklistItem SUCCESS ----------------------------------")
}
