package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service/comment"

	logger "server/internal/logging"
)

type CommentHandler struct {
	cs comment.ICommentService
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
	failBorder := "----------------- Create comment FAIL -----------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("----------------- Create comment -----------------")

	var newCommentInfo dto.NewCommentInfo
	err := json.NewDecoder(r.Body).Decode(&newCommentInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("JSON Decoded", funcName, nodeName)

	comment, err := ch.cs.Create(rCtx, newCommentInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Comment created", funcName, nodeName)

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
	logger.Debug("response written", funcName, nodeName)

	logger.Info("----------------- Create comment SUCCESS -----------------")
}
