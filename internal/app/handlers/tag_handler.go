package handlers

import (
	"fmt"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
)

type TagHandler struct {
	ts service.ITagService
}

// @Summary Создать тэг
// @Description Создать тэг
// @Tags tags
//
// @Accept  json
// @Produce  json
//
// @Param newTagInfo body dto.NewTagInfo true "данные нового тэга"
//
// @Success 200  {object}  doc_structs.TagResponse "объект тэга"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /tag/create/ [post]
func (th TagHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TagHandler.Create"
	errorMessage := "Creating tag failed with error: "
	failBorder := "---------------------------------- Creating Tag FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Creating Tag ----------------------------------")

	var newTagInfo dto.NewTagInfo
	err := easyjson.UnmarshalFromReader(r.Body, &newTagInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	logger.DebugFmt(fmt.Sprintf("Adding tag pre-color %v", newTagInfo), requestID.String(), funcName, nodeName)
	newTagInfo.Color = "FFFFFF"
	logger.DebugFmt(fmt.Sprintf("Adding tag post-color %v", newTagInfo), requestID.String(), funcName, nodeName)
	tag, err := th.ts.Create(rCtx, newTagInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Tag created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"tag": tag,
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

	logger.Info("---------------------------------- Creating Tag FAIL ----------------------------------")
}

// @Summary Обновить тэг
// @Description Обновить тэг
// @Tags tags
//
// @Accept  json
// @Produce  json
//
// @Param tagInfo body dto.UpdatedTagInfo true "обновленные данные тэга"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /tag/update/ [post]
func (th TagHandler) Update(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TagHandler.Update"
	errorMessage := "Updating failed with error: "
	failBorder := "---------------------------------- Updating Tag FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Updating Tag ----------------------------------")

	var tagInfo dto.UpdatedTagInfo
	err := easyjson.UnmarshalFromReader(r.Body, &tagInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = th.ts.Update(rCtx, tagInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Tag updated", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Updating Tag SUCCESS ----------------------------------")
}

// @Summary Удалить тэг
// @Description Удалить тэг
// @Tags Tags
//
// @Accept  json
// @Produce  json
//
// @Param tagID body dto.TagID true "id тэга"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /tag/delete/ [delete]
func (th TagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TagHandler.Delete"
	errorMessage := "Deleting failed with error: "
	failBorder := "---------------------------------- Deleting Tag FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Deleting Tag ----------------------------------")

	var TagID dto.TagID
	err := easyjson.UnmarshalFromReader(r.Body, &TagID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = th.ts.Delete(rCtx, TagID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Tag deleted", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Deleting Tag SUCCESS ----------------------------------")
}

// @Summary Добавить тэг к заданию
// @Description Добавить тэг к заданию
// @Tags tags
//
// @Accept  json
// @Produce  json
//
// @Param tagAndTaskIDs body dto.TagAndTaskIDs true "id тэга и связанного задания"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /tag/add_to_task/ [post]
func (th TagHandler) AddToTask(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TagHandler.AddToTask"
	errorMessage := "Adding tag to task failed with error: "
	failBorder := "---------------------------------- Adding tag to task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Adding tag to task ----------------------------------")

	var tagAndTaskIDs dto.TagAndTaskIDs
	err := easyjson.UnmarshalFromReader(r.Body, &tagAndTaskIDs)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = th.ts.AddToTask(rCtx, tagAndTaskIDs)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Tag added to task", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Adding tag to task SUCCESS ----------------------------------")
}

// @Summary Добавить тэг к заданию
// @Description Добавить тэг к заданию
// @Tags tags
//
// @Accept  json
// @Produce  json
//
// @Param tagAndTaskIDs body dto.TagAndTaskIDs true "id тэга и связанного задания"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /tag/remove_from_task/ [post]
func (th TagHandler) RemoveFromTask(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TagHandler.RemoveFromTask"
	errorMessage := "Removing tag from task failed with error: "
	failBorder := "---------------------------------- Removing tag from task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Removing tag from task ----------------------------------")

	var tagAndTaskIDs dto.TagAndTaskIDs
	err := easyjson.UnmarshalFromReader(r.Body, &tagAndTaskIDs)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = th.ts.RemoveFromTask(rCtx, tagAndTaskIDs)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Tag removed from task", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Removing tag from task SUCCESS ----------------------------------")
}
