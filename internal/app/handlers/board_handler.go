package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service"

	logger "server/internal/logging"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type BoardHandler struct {
	as service.IAuthService
	bs service.IBoardService
}

// @Summary Получить доску
// @Description Получить доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param boardID body dto.BoardID true "id доски"
//
// @Success 200  {object}  doc_structs.BoardResponse "объект доски"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/ [post]
func (bh BoardHandler) GetFullBoard(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "GetFullBoard"
	errorMessage := "Getting full board failed with error: "
	failBorder := "---------------------------------- Get board FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Get board ----------------------------------")

	var boardID dto.BoardID
	err := json.NewDecoder(r.Body).Decode(&boardID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", requestID.String(), funcName, nodeName)

	boardRequest := dto.IndividualBoardRequest{
		UserID:  user.ID,
		BoardID: boardID.Value,
	}
	board, err := bh.bs.GetFullBoard(rCtx, boardRequest)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Got board", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: board,
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- Get board SUCCESS ----------------------------------")
}

// @Summary Создать доску
// @Description Создать доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param newBoardRequest body dto.NewBoardRequest true "данные новой доски"
//
// @Success 200  {object}  doc_structs.BoardResponse "объект доски"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/create/ [post]
func (bh BoardHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "Create"
	errorMessage := "Getting full board failed with error: "
	failBorder := "---------------------------------- Create board FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)
	logger.Info("---------------------------------- Creating board ----------------------------------")

	var newBoardRequest dto.NewBoardRequest
	err := json.NewDecoder(r.Body).Decode(&newBoardRequest)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", requestID.String(), funcName, nodeName)

	_, err = govalidator.ValidateStruct(newBoardRequest)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("New board data validated", requestID.String(), funcName, nodeName)

	newBoardInfo := dto.NewBoardInfo{
		Name:        newBoardRequest.Name,
		WorkspaceID: newBoardRequest.WorkspaceID,
		OwnerID:     user.ID,
		Thumbnail:   newBoardRequest.Thumbnail,
	}
	board, err := bh.bs.Create(rCtx, newBoardInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Board created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"board": board,
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

	logger.Info("---------------------------------- Create board SUCCESS ----------------------------------")
}

// @Summary Обновить доску
// @Description Обновить доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param boardInfo body dto.UpdatedBoardInfo true "обновленные данные доски"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/update/ [post]
func (bh BoardHandler) UpdateData(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "UpdateData"
	errorMessage := "Updating full board failed with error: "
	failBorder := "---------------------------------- Updating board FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)
	logger.Info("---------------------------------- Updating board ----------------------------------")

	var boardInfo dto.UpdatedBoardInfo
	err := json.NewDecoder(r.Body).Decode(&boardInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	_, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", requestID.String(), funcName, nodeName)

	err = bh.bs.UpdateData(rCtx, boardInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("board data updated", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Updating board SUCCESS ----------------------------------")
}

// @Summary Обновить картинку доски
// @Description Обновить картинку доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param boardInfo body dto.UpdatedBoardThumbnailInfo true "обновленные данные задания"
//
// @Success 200  {object}  doc_structs.ThumbnailUploadResponse "Ссылка на новую картинку"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/update/change_thumbnail/ [post]
func (bh BoardHandler) UpdateThumbnail(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "UpdateThumbnail"
	errorMessage := "Updating board thumbnail failed with error: "
	failBorder := "---------------------------------- Updating board thumbnail FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)
	logger.Info("---------------------------------- Updating board ----------------------------------")

	var boardInfo dto.UpdatedBoardThumbnailInfo
	err := json.NewDecoder(r.Body).Decode(&boardInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	_, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", requestID.String(), funcName, nodeName)

	urlObj, err := bh.bs.UpdateThumbnail(rCtx, boardInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("board thumbnail updated", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"url": urlObj,
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

	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)
	logger.Info("---------------------------------- Updating board thumbnail SUCCESS ----------------------------------")
}

// @Summary Удалить доску
// @Description Удалить доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param boardID body dto.BoardID true "id доски"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/delete/ [delete]
func (bh BoardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "Delete"
	errorMessage := "Deleting board failed with error: "
	failBorder := "---------------------------------- Deleting board FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)
	logger.Info("---------------------------------- Deleting board ----------------------------------")

	var boardID dto.BoardID
	err := json.NewDecoder(r.Body).Decode(&boardID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	_, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", requestID.String(), funcName, nodeName)

	err = bh.bs.Delete(rCtx, boardID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("board deleted", requestID.String(), funcName, nodeName)

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

	logger.DebugFmt("response written", requestID.String(), funcName, nodeName)
	logger.Info("---------------------------------- Deleting board SUCCESS ----------------------------------")
}

// @Summary Добавить пользователя в доску
// @Description Добавить пользователя в доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param info body dto.AddBoardUserRequest true "мэйл пользователя, id доски и воркспейса"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/user/add/ [post]
func (bh BoardHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "AddUser"
	errorMessage := "Adding user to board with error: "
	failBorder := "---------------------------------- Adding user to board FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)
	logger.Info("---------------------------------- Adding user to board ----------------------------------")

	var info dto.AddBoardUserRequest
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	_, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", requestID.String(), funcName, nodeName)

	err = bh.bs.AddUser(rCtx, info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User added", requestID.String(), funcName, nodeName)

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

	logger.DebugFmt("Response written", requestID.String(), funcName, nodeName)
	logger.Info("---------------------------------- Add user to board SUCCESS ----------------------------------")
}

// @Summary Удалить пользователя из доски
// @Description Удалить пользователя из доски
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param info body dto.RemoveBoardUserInfo true "id пользователя, доски и воркспейса"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/user/remove/ [delete]
func (bh BoardHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "RemoveUser"
	errorMessage := "Removing user from board failed with error: "
	failBorder := "---------------------------------- Removing user from board FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)
	logger.Info("---------------------------------- Removing user from board ----------------------------------")

	var info dto.RemoveBoardUserInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + "User not found")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("User object acquired from context", requestID.String(), funcName, nodeName)

	if user.ID == info.UserID {
		logger.Error(errorMessage + "user cannot remove himself from the board")
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}

	err = bh.bs.RemoveUser(rCtx, info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User removed", requestID.String(), funcName, nodeName)

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

	logger.DebugFmt("Response written", requestID.String(), funcName, nodeName)
	logger.Info("---------------------------------- Removing user from board SUCCESS ----------------------------------")
}
