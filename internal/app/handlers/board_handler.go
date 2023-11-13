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
	"server/internal/service"
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
	log.Println("--------------BoardHandler.GetFullBoard Endpoint START--------------")

	rCtx := r.Context()

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		log.Println("user not found")
		log.Println("--------------BoardHandler.GetFullBoard Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}

	var boardID dto.BoardID
	err := json.NewDecoder(r.Body).Decode(&boardID)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.GetFullBoard Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	boardRequest := dto.IndividualBoardRequest{
		UserID:  user.ID,
		BoardID: boardID.Value,
	}

	board, err := bh.bs.GetFullBoard(rCtx, boardRequest)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.GetFullBoard Endpoint FAIL--------------")
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
		log.Println(err)
		log.Println("--------------BoardHandler.GetFullBoard Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalleed")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.GetFullBoard Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------BoardHandler.GetFullBoard Endpoint SUCCESS--------------")
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
	log.Println("--------------BoardHandler.Create Endpoint START--------------")

	rCtx := r.Context()

	var newBoardRequest dto.NewBoardRequest
	err := json.NewDecoder(r.Body).Decode(&newBoardRequest)
	if err != nil {
		log.Println("Failed to decode incoming JSON")
		log.Println(err)
		log.Println("--------------BoardHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		log.Println("user not found")
		log.Println("--------------BoardHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
		return
	}
	log.Println("user found")

	// // _, err = govalidator.ValidateStruct(newBoardInfo)
	// // if err != nil {
	// // 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// // 	return
	// // }

	newBoardInfo := dto.NewBoardInfo{
		Name:        newBoardRequest.Name,
		Description: newBoardRequest.Description,
		Thumbnail:   newBoardRequest.Thumbnail,
		WorkspaceID: newBoardRequest.WorkspaceID,
		OwnerID:     user.ID,
	}

	board, err := bh.bs.Create(rCtx, newBoardInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("board created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"board": board,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------BoardHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------BoardHandler.Create Endpoint SUCCESS--------------")
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
