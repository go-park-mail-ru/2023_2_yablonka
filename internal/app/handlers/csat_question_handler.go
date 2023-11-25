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
// @Success 200  {object}  doc_structs.AllQuestionsResponse "все вопросы"
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

// @Summary Создать вопрос CSAT
// @Description Создать вопрос CSAT
// @Tags csat
//
// @Accept  json
// @Produce  json
//
// @Param newQuestionInfo body dto.NewCSATQuestionInfo true "данные нового вопроса CSAT"
//
// @Success 200  {object}  doc_structs.QuestionResponse "вопрос CSAT"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /csat/question/create/ [post]
func (qh CSATQuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "Create"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Creating a new CSAT question")

	var newQuestionInfo dto.NewCSATQuestionInfo
	err := json.NewDecoder(r.Body).Decode(&newQuestionInfo)
	if err != nil {
		logger.Error("Creating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Creating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	question, err := qh.qs.Create(rCtx, newQuestionInfo)
	if err != nil {
		logger.Error("Creating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Creating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Task created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"question": question,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Creating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Creating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Creating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Creating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished creating CSAT question")
}

func (qh CSATQuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "Update"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Updating a new CSAT question")

	var updatedQuestionInfo dto.UpdatedCSATQuestionInfo
	err := json.NewDecoder(r.Body).Decode(&updatedQuestionInfo)
	if err != nil {
		logger.Error("Updating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Updating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	err = qh.qs.Update(rCtx, updatedQuestionInfo)
	if err != nil {
		logger.Error("Updating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Updating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Task created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("Updating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Updating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Updating a new CSAT question failed")
		handlerDebugLog(logger, funcName, "Updating a new CSAT question failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished updating CSAT question")
}

func (qh CSATQuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
