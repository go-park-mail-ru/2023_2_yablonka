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

	"github.com/google/uuid"
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
	funcName := "CSATAnswerHandler.Create"
	nodeName := "handler"
	errorMessage := "Creating CSAT answer failed with error: "
	failBorder := "---------------------------------- Create CSAT answer FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Create CSAT answer ----------------------------------")

	var CSATAnswerInfo dto.NewCSATAnswerInfo
	err := json.NewDecoder(r.Body).Decode(&CSATAnswerInfo)
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

	err = ah.qs.CheckRating(rCtx, CSATAnswerInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Rating checked", requestID.String(), funcName, nodeName)

	CSATAnswer := dto.NewCSATAnswer{
		UserID:     user.ID,
		QuestionID: CSATAnswerInfo.QuestionID,
		Rating:     CSATAnswerInfo.Rating,
	}
	err = ah.as.Create(rCtx, CSATAnswer)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Answer created", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Create CSAT answer SUCCESS ----------------------------------")
}
