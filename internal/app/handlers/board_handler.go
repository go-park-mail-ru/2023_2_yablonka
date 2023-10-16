package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"
)

type BoardHandler struct {
	bs service.IBoardService
}

// @Summary Создает новую доску
// @Description Добавляет доску в БД и возвращает её данные. Работает только для авторизированного пользователя.
// @Tags boards
//
// @Param boardInfo body dto.NewBoardInfo true "Данные новой доски"
//
// @Accept  json
// @Produce  json
//
// @Success 200  {object}  doc_structs.NewBoardResponse "Новая доска"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /api/v2/boards/create/ [post]
func (bh BoardHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var boardInfo dto.NewBoardInfo

	err := json.NewDecoder(r.Body).Decode(&boardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	newBoard, err := bh.bs.CreateBoard(rCtx, boardInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"board": newBoard,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}
