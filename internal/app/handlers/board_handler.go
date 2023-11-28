package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service/auth"
	"server/internal/service/board"

	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

type BoardHandler struct {
	as auth.IAuthService
	bs board.IBoardService
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

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Getting a board")

	var boardID dto.BoardID
	err := json.NewDecoder(r.Body).Decode(&boardID)
	if err != nil {
		logger.Error("Getting a board failed")
		handlerDebugLog(logger, funcName, "Getting a board failed with error: "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error("Getting a board failed")
		handlerDebugLog(logger, funcName, "Getting a board failed -- no user passed in context")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User object acquired from context")

	boardRequest := dto.IndividualBoardRequest{
		UserID:  user.ID,
		BoardID: boardID.Value,
	}

	board, err := bh.bs.GetFullBoard(rCtx, boardRequest)
	if err != nil {
		logger.Error("Getting a board failed")
		handlerDebugLog(logger, funcName, "Getting a board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Got board")

	response := dto.JSONResponse{
		Body: board,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Getting a board failed")
		handlerDebugLog(logger, funcName, "Getting a board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Getting a board failed")
		handlerDebugLog(logger, funcName, "Getting a board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished getting board")
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

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Creating a new board")

	var newBoardRequest dto.NewBoardRequest
	err := json.NewDecoder(r.Body).Decode(&newBoardRequest)
	if err != nil {
		logger.Error("Creating a new board failed")
		handlerDebugLog(logger, funcName, "Creating a new board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error("Creating a new board failed")
		handlerDebugLog(logger, funcName, "Creating a new board failed failed -- no user passed in context")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User object acquired from context")

	_, err = govalidator.ValidateStruct(newBoardRequest)
	if err != nil {
		logger.Error("Creating a new board failed")
		handlerDebugLog(logger, funcName, "Creating a new board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "New board data validated")

	newBoardInfo := dto.NewBoardInfo{
		Name:        newBoardRequest.Name,
		WorkspaceID: newBoardRequest.WorkspaceID,
		OwnerID:     user.ID,
		Thumbnail:   newBoardRequest.Thumbnail,
	}

	board, err := bh.bs.Create(rCtx, newBoardInfo)
	if err != nil {
		logger.Error("Creating a new board failed")
		handlerDebugLog(logger, funcName, "Creating a new board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Board created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"board": board,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Creating a new board failed")
		handlerDebugLog(logger, funcName, "Creating a new board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Creating a new board failed")
		handlerDebugLog(logger, funcName, "Creating a new board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished creating board")
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
	log.Println("--------------BoardHandler.UpdateData Endpoint START--------------")

	rCtx := r.Context()

	var boardInfo dto.UpdatedBoardInfo
	err := json.NewDecoder(r.Body).Decode(&boardInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateData Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(boardInfo)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	err = bh.bs.UpdateData(rCtx, boardInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateData Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("board data updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateData Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateData Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------BoardHandler.UpdateData Endpoint SUCCESS--------------")
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
	log.Println("--------------BoardHandler.UpdateThumbnail Endpoint START--------------")
	rCtx := r.Context()

	var boardInfo dto.UpdatedBoardThumbnailInfo
	err := json.NewDecoder(r.Body).Decode(&boardInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateThumbnail Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(boardInfo)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	//boardInfo.BaseURL = r.URL.Scheme + "://" + r.URL.Host

	urlObj, err := bh.bs.UpdateThumbnail(rCtx, boardInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateThumbnail Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("board thumbnail updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"url": urlObj,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateThumbnail Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marchalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.UpdateThumbnail Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------BoardHandler.UpdateThumbnail Endpoint SUCCESS--------------")
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
	log.Println("--------------BoardHandler.Delete Endpoint START--------------")

	rCtx := r.Context()

	var boardID dto.BoardID
	err := json.NewDecoder(r.Body).Decode(&boardID)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(boardID)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	err = bh.bs.Delete(rCtx, boardID)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("board deleted")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------LogIn Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------LogIn Endpoint SUCCESS--------------")
}

func (bh BoardHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "AddUser"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Adding user to board")

	var info dto.AddBoardUserRequest
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.Error("Adding user to board failed")
		handlerDebugLog(logger, funcName, "Adding user to board failed with error: "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	_, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error("Adding user to board failed")
		handlerDebugLog(logger, funcName, "Adding user to board failed -- no user passed in context")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User object acquired from context")

	err = bh.bs.AddUser(rCtx, info)
	if err != nil {
		logger.Error("Adding user to board failed")
		handlerDebugLog(logger, funcName, "Adding user to board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User added")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Adding user to board failed")
		handlerDebugLog(logger, funcName, "Adding user to board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Adding user to board failed")
		handlerDebugLog(logger, funcName, "Adding user to board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished adding user to board")
}

func (bh BoardHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "RemoveUser"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Removing user from board")

	var info dto.RemoveBoardUserInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.Error("Removing user from board failed")
		handlerDebugLog(logger, funcName, "Removing user from board failed with error: "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	_, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error("Removing user from board failed")
		handlerDebugLog(logger, funcName, "Removing user from board failed -- no user passed in context")
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User object acquired from context")

	err = bh.bs.RemoveUser(rCtx, info)
	if err != nil {
		logger.Error("Removing user from board failed")
		handlerDebugLog(logger, funcName, "Removing user from board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User removed")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Removing user from board failed")
		handlerDebugLog(logger, funcName, "Removing user from board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Removing user from board failed")
		handlerDebugLog(logger, funcName, "Removing user from board failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished Removing user from board")
}
