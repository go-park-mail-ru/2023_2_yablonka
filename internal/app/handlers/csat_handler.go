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

	"github.com/sirupsen/logrus"
)

type CSATHandler struct {
	cs service.ICSATSAnswerService
}

// @Summary Ответить на опрос CSAT
// @Description Создать ответит на опрос CSAT
// @Tags csat
//
// @Accept  json
// @Produce  json
//
// @Param CSATAnswerInfo body dto.NewCSATAnswerInfo true "данные ответа CSAT"
//
// @Success 204  {string} string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /csat/answer/ [post]
func (ch CSATHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "GetFullBoard"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)

	var CSATAnswerInfo dto.NewCSATAnswerInfo
	err := json.NewDecoder(r.Body).Decode(&CSATAnswerInfo)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	handlerDebugLog(logger, funcName, "request struct decoded")

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "User object acquired from context")

	CSATAnswer := dto.NewCSATAnswer{
		UserID:     user.ID,
		QuestionID: CSATAnswerInfo.QuestionID,
		Rating:     CSATAnswerInfo.Rating,
	}

	err = ch.cs.Create(rCtx, CSATAnswer)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("list created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	handlerDebugLog(logger, funcName, "Response written")
}
