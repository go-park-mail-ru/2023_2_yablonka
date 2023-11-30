package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	logger "server/internal/logging"
)

type FileHandler struct {
	fs service.IFileService
}

// @Summary Загрузить изображение
// @Description Загрузить изображение
// @Tags images
//
// @Accept  json
// @Produce  json
//
// @Param image body dto.Image true "байты изоьражения"
//
// @Success 200  {object}  doc_structs.URLResponse "ссылка на изображение"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /image/upload/ [post]
func (fh FileHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "UploadImage"
	errorMessage := "Uploading image failed with error: "
	failBorder := "---------------------------------- Uploading image FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

	logger.Info("---------------------------------- Uploading image ----------------------------------")

	var image dto.Image
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.Debug("JSON Decoded", funcName, nodeName)

	url, err := fh.fs.Upload(rCtx, image)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.Debug("Uploaded image", funcName, nodeName)

	response := dto.JSONResponse{
		Body: url,
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
