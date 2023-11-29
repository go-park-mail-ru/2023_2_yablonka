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
)

type FileHandler struct {
	fs service.IFileService
}

// @Summary Загрузить изображение
// @Description Получить доску
// @Tags boards
//
// @Accept  json
// @Produce  json
//
// @Param boardID body dto.BoardID true "id доски"
//
// @Success 200  {object}  doc_structs.BoardResponse "объект доски"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /board/ [post]
func (bh BoardHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "UploadImage"
	errorMessage := "Uploading image failed with error: "
	failBorder := "---------------------------------- Uploading image FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Uploading image ----------------------------------")

	var boardID dto.BoardID
	err := json.NewDecoder(r.Body).Decode(&boardID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("JSON Decoded", funcName, nodeName)

	user, ok := rCtx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
		return
	}
	logger.Debug("User object acquired from context", funcName, nodeName)

	boardRequest := dto.IndividualBoardRequest{
		UserID:  user.ID,
		BoardID: boardID.Value,
	}
	board, err := bh.bs.GetFullBoard(rCtx, boardRequest)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Got board", funcName, nodeName)

	response := dto.JSONResponse{
		Body: board,
	}
	err = WriteResponse(response, w, r)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.Debug("response written", funcName, nodeName)

	logger.Info("---------------------------------- Uploading image SUCCESS ----------------------------------")
}
