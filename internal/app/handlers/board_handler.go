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
// @Param boardRequest body dto.IndividualBoardRequest true "id доски"
//
// @Success 200  {object}  doc_structs.BoardResponse "объект доски"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/ [get]
func (bh BoardHandler) GetFullBoard(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var boardRequest dto.IndividualBoardRequest
	err := json.NewDecoder(r.Body).Decode(&boardRequest)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(boardRequest)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	board, err := bh.bs.GetFullBoard(rCtx, boardRequest)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"board": board,
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

// @Summary Создать доску
// @Description Создать доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param newBoardInfo body dto.NewBoardInfo true "данные новой доски"
//
// @Success 200  {object}  doc_structs.BoardResponse "объект доски"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/create/ [post]
func (bh BoardHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var newBoardInfo dto.NewBoardInfo
	err := json.NewDecoder(r.Body).Decode(&newBoardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(newBoardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	board, err := bh.bs.Create(rCtx, newBoardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"board": board,
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

	var boardInfo dto.UpdatedBoardInfo
	err := json.NewDecoder(r.Body).Decode(&boardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(boardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	err = bh.bs.UpdateData(rCtx, boardInfo)
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
// @Router /board/update/change_thumbnail [post]
func (bh BoardHandler) UpdateThumbnail(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var boardInfo dto.UpdatedBoardThumbnailInfo
	err := json.NewDecoder(r.Body).Decode(&boardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(boardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	urlObj, err := bh.bs.UpdateThumbnail(rCtx, boardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"url": urlObj,
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

	var boardID dto.BoardID
	err := json.NewDecoder(r.Body).Decode(&boardID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(boardID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	err = bh.bs.Delete(rCtx, boardID)
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
