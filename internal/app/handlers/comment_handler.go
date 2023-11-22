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
	funcName := "Create"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("Creating a new comment")

	var newCommentInfo dto.NewCommentInfo
	err := json.NewDecoder(r.Body).Decode(&newCommentInfo)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new task failed with error "+err.Error())
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	// _, err = govalidator.ValidateStruct(newCommentInfo)
	// if err != nil {
	// 	logger.Error("Creating a new board failed")
	// 	handlerDebugLog(logger, funcName, "Creating a new board failed with error "+err.Error())
	// 	apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
	// 	return
	// }
	// handlerDebugLog(logger, funcName, "New task data validated")

	comment, err := ch.cs.Create(rCtx, newCommentInfo)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new comment failed with error "+err.Error())
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Comment created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"comment": comment,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new comment failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new comment failed with error "+err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("Finished creating comment")
}
