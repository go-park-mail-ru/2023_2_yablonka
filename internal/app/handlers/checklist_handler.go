package handlers

import (
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
)

type ChecklistHandler struct {
	cls service.IChecklistService
}

// @Summary Создать чеклист
// @Description Создать чеклист
// @Tags checklists
//
// @Accept  json
// @Produce  json
//
// @Param newChecklistInfo body dto.NewChecklistInfo true "данные нового чеклиста"
//
// @Success 200  {object}  doc_structs.ChecklistResponse "объект чеклиста"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /checklist/create/ [post]
func (clh ChecklistHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChecklistHandler.Create"
	nodeName := "handler"
	errorMessage := "Creating Checklist failed with error: "
	failBorder := "---------------------------------- Creating Checklist FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Creating Checklist ----------------------------------")

	var newChecklistInfo dto.NewChecklistInfo
	err := easyjson.UnmarshalFromReader(r.Body, &newChecklistInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	Checklist, err := clh.cls.Create(rCtx, newChecklistInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Checklist created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"checklist": Checklist,
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

	logger.Info("---------------------------------- Creating Checklist FAIL ----------------------------------")
}

// @Summary Обновить чеклист
// @Description Обновить чеклист
// @Tags checklists
//
// @Accept  json
// @Produce  json
//
// @Param ChecklistInfo body dto.UpdatedChecklistInfo true "обновленные данные чеклиста"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /checklist/update/ [post]
func (clh ChecklistHandler) Update(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChecklistHandler.Update"
	nodeName := "handler"
	errorMessage := "Updating failed with error: "
	failBorder := "---------------------------------- Updating Checklist FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Updating Checklist ----------------------------------")

	var ChecklistInfo dto.UpdatedChecklistInfo
	err := easyjson.UnmarshalFromReader(r.Body, &ChecklistInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = clh.cls.Update(rCtx, ChecklistInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Checklist updated", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Updating Checklist SUCCESS ----------------------------------")
}

// @Summary Удалить чеклист
// @Description Удалить чеклист
// @Tags checklists
//
// @Accept  json
// @Produce  json
//
// @Param ChecklistID body dto.ChecklistID true "id чеклиста"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /checklist/delete/ [delete]
func (clh ChecklistHandler) Delete(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "ChecklistHandler.Delete"
	nodeName := "handler"
	errorMessage := "Deleting failed with error: "
	failBorder := "---------------------------------- Deleting Checklist FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Deleting Checklist ----------------------------------")

	var checklistID dto.ChecklistID
	err := easyjson.UnmarshalFromReader(r.Body, &checklistID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = clh.cls.Delete(rCtx, checklistID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Checklist deleted", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Deleting Checklist SUCCESS ----------------------------------")
}
