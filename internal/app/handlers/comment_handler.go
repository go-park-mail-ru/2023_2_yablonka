package handlers

import (
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	logger "server/internal/logging"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
)

type CommentHandler struct {
	cs service.ICommentService
}

// @Summary Создать комментарий
// @Description Создать комментарий в каком-то задании
// @Tags comments
//
// @Accept  json
// @Produce  json
//
// @Param newCommentInfo body dto.NewCommentInfo true "данные нового комментария"
//
// @Success 200  {object}  doc_structs.CommentResponse "объект комментария"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /comment/create/ [post]
func (ch CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "CommentHandler.Create"
	nodeName := "handler"
	errorMessage := "Creating comment failed with error: "
	failBorder := "---------------------------------- Create comment FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Create comment ----------------------------------")

	var newCommentInfo dto.NewCommentInfo
	err := easyjson.UnmarshalFromReader(r.Body, &newCommentInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	comment, err := ch.cs.Create(rCtx, newCommentInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Comment created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"comment": comment,
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

	logger.Info("---------------------------------- Create comment SUCCESS ----------------------------------")
}
