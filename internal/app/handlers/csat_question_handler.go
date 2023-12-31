package handlers

import (
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
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
	funcName := "CSATQuestionHandler.GetQuestions"
	errorMessage := "Getting all CSAT questions failed with error: "
	failBorder := "---------------------------------- Get all CSAT questions FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Get all CSAT questions ----------------------------------")

	questions, err := qh.qs.GetAll(rCtx)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Got questions", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"questions": questions,
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

	logger.Info("---------------------------------- Getting all CSAT questions SUCCESS ----------------------------------")
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
	funcName := "CSATQuestionHandler.Create"
	errorMessage := "Creating CSAT question failed with error: "
	failBorder := "---------------------------------- Create CSAT question FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Create CSAT question ----------------------------------")

	var newQuestionInfo dto.NewCSATQuestionInfo
	err := easyjson.UnmarshalFromReader(r.Body, &newQuestionInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	question, err := qh.qs.Create(rCtx, newQuestionInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Task created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"question": question,
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

	logger.Info("---------------------------------- Create CSAT question SUCCESS ----------------------------------")
}

// @Summary Обновить вопрос CSAT
// @Description Обновить вопрос CSAT
// @Tags csat
//
// @Accept  json
// @Produce  json
//
// @Param updatedQuestionInfo body dto.UpdatedCSATQuestionInfo true "данные обновленного вопроса CSAT"
//
// @Success 204  {string} string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /csat/question/update/ [post]
func (qh CSATQuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "CSATQuestionHandler.Update"
	errorMessage := "Updating CSAT question failed with error: "
	failBorder := "---------------------------------- Update CSAT question FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Update CSAT question ----------------------------------")

	var updatedQuestionInfo dto.UpdatedCSATQuestionInfo
	err := easyjson.UnmarshalFromReader(r.Body, &updatedQuestionInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	err = qh.qs.Update(rCtx, updatedQuestionInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("CSAT question updated", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Update CSAT question SUCCESS ----------------------------------")
}

// @Summary Статистика ответов CSAT
// @Description Получить количество и среднее ответов на опросы CSAT
// @Tags csat
//
// @Accept  json
// @Produce  json
//
// @Success 200  {object}  doc_structs.StatsResponse "вопрос CSAT"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /csat/stats [get]
func (qh CSATQuestionHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "CSATQuestionHandler.GetStats"
	errorMessage := "Getting CSAT question stats failed with error: "
	failBorder := "---------------------------------- Getting CSAT question stats FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Getting CSAT question stats ----------------------------------")

	questions, err := qh.qs.GetStats(rCtx)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Stats received", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"questions": questions,
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

	logger.Info("---------------------------------- Getting CSAT question stats SUCCESS ----------------------------------")
}
