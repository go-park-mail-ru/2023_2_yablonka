package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/sirupsen/logrus"
)

type CSATQuestionHandler struct {
	qs service.ICSATQuestionService
}

// @Summary Получить вопросы CSAT
// @Description Получить вопросы CSAT
// @Tags csat
//
// @Accept  json
// @Produce  json
//
// @Success 200  {object}  doc_structs.QuestionsResponse "все вопросы"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /csat/question/all [get]
func (qh CSATQuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "GetQuestions"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Getting all questions")

	questions, err := qh.qs.GetAll(rCtx)
	if err != nil {
		handlerDebugLog(logger, funcName, "Getting questions failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Got questions")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"questions": questions,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handlerDebugLog(logger, funcName, "Getting questions failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		handlerDebugLog(logger, funcName, "Getting questions failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
}

func (ah CSATQuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
