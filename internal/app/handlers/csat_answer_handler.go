package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service"

	"github.com/sirupsen/logrus"
)

type CSATAnswerHandler struct {
	as service.ICSATSAnswerService
	qs service.ICSATQuestionService
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
func (ah CSATAnswerHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "CreateCSATAnswer"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)

	var CSATAnswerInfo dto.NewCSATAnswerInfo
	err := json.NewDecoder(r.Body).Decode(&CSATAnswerInfo)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
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

	err = ah.qs.CheckRating(rCtx, CSATAnswerInfo)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Rating checked")

	CSATAnswer := dto.NewCSATAnswer{
		UserID:     user.ID,
		QuestionID: CSATAnswerInfo.QuestionID,
		Rating:     CSATAnswerInfo.Rating,
	}
	err = ah.as.Create(rCtx, CSATAnswer)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Answer created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a CSAT answer failed -- "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()
	handlerDebugLog(logger, funcName, "Response written")
}
