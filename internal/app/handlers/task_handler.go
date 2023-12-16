package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/google/uuid"
)

type TaskHandler struct {
	ts service.ITaskService
}

// @Summary Создать задание
// @Description Создать задание
// @Tags tasks
//
// @Accept  json
// @Produce  json
//
// @Param newTaskInfo body dto.NewTaskInfo true "данные нового задания"
//
// @Success 200  {object}  doc_structs.TaskResponse "объект задания"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /task/create/ [post]
func (th TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TaskHandler.Create"
	nodeName := "handler"
	errorMessage := "Creating a new task failed with error: "
	failBorder := "---------------------------------- Creating a new task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Creating a new task ----------------------------------")

	var newTaskInfo dto.NewTaskInfo
	err := json.NewDecoder(r.Body).Decode(&newTaskInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	task, err := th.ts.Create(rCtx, newTaskInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("Task created", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"task": task,
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

	logger.Info("---------------------------------- Creating a new task SUCCESS ----------------------------------")
}

// @Summary Получить задание
// @Description Получить задание
// @Tags tasks
//
// @Accept  json
// @Produce  json
//
// @Param taskID body dto.TaskID true "id задания"
//
// @Success 200  {object}  doc_structs.TaskResponse "объект задания"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /task/ [post]
func (th TaskHandler) Read(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TaskHandler.Read"
	nodeName := "handler"
	errorMessage := "Getting task failed with error: "
	failBorder := "---------------------------------- Getting task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Getting task ----------------------------------")

	var taskID dto.TaskID
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	task, err := th.ts.Read(rCtx, taskID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("task read", requestID.String(), funcName, nodeName)

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"task": task,
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

	logger.Info("---------------------------------- Getting task SUCCESS ----------------------------------")
}

// @Summary Обновить задание
// @Description Обновить задание
// @Tags tasks
//
// @Accept  json
// @Produce  json
//
// @Param taskInfo body dto.UpdatedTaskInfo true "обновленные данные задания"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /task/update/ [post]
func (th TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TaskHandler.Update"
	nodeName := "handler"
	errorMessage := "Updating task failed with error: "
	failBorder := "---------------------------------- Updating task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Updating task ----------------------------------")

	var taskInfo dto.UpdatedTaskInfo
	err := json.NewDecoder(r.Body).Decode(&taskInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	err = th.ts.Update(rCtx, taskInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("task updated", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Updating task SUCCESS ----------------------------------")
}

// @Summary Удалить задание
// @Description Удалить задание
// @Tags tasks
//
// @Accept  json
// @Produce  json
//
// @Param taskID body dto.TaskID true "id задания"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /task/delete/ [delete]
func (th TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TaskHandler.Delete"
	nodeName := "handler"
	errorMessage := "Deleting task failed with error: "
	failBorder := "---------------------------------- Deleting task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Deleting task ----------------------------------")

	var taskID dto.TaskID
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	err = th.ts.Delete(rCtx, taskID)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("task deleted", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- Deleting task SUCCESS ----------------------------------")
}

// @Summary Добавить пользователя на карточку
// @Description Добавить пользователя на карточку
// @Tags tasks
//
// @Accept  json
// @Produce  json
//
// @Param taskInfo body dto.AddTaskUserInfo true "ID пользователя и ID карточки, на которую добавляют пользователя"
//
// @Success 200  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /task/user/add/ [post]
func (th TaskHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TaskHandler.AddUser"
	nodeName := "handler"
	errorMessage := "Adding user to task failed with error: "
	failBorder := "---------------------------------- Adding user to task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Adding user to task ----------------------------------")

	var info dto.AddTaskUserInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	err = th.ts.AddUser(rCtx, info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User added", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("Response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- Adding user to task SUCCESS ----------------------------------")
}

// @Summary Удалить пользователя из карточки
// @Description Удалить пользователя из карточки
// @Tags tasks
//
// @Accept  json
// @Produce  json
//
// @Param taskInfo body dto.RemoveTaskUserInfo true "ID пользователя и ID карточки, из которой удаляют пользователя"
//
// @Success 200  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /task/user/remove/ [post]
func (th TaskHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TaskHandler.RemoveUser"
	nodeName := "handler"
	errorMessage := "Removing user from task failed with error: "
	failBorder := "---------------------------------- Removing user from task FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- Removing user from task ----------------------------------")

	var info dto.RemoveTaskUserInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("JSON Decoded", requestID.String(), funcName, nodeName)

	err = th.ts.RemoveUser(rCtx, info)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("User removed", requestID.String(), funcName, nodeName)

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
	logger.DebugFmt("Response written", requestID.String(), funcName, nodeName)

	logger.Info("---------------------------------- Removing user from task SUCCESS ----------------------------------")
}

// @Summary Перенести задание в другой список
// @Description Меняет порядок у заданий в старом и новом списках и меняет связь задания со списком
// @Tags tasks
//
// @Accept  json
// @Produce  json
//
// @Param taskMoveInfo body dto.TaskMoveInfo true "id заданий из обоих списков и id нового списка"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /task/move/ [post]
func (th TaskHandler) Move(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()
	funcName := "TaskHandler.UpdateOrder"
	errorMessage := "Updating task order failed with error: "
	failBorder := "---------------------------------- TaskHandler.UpdateOrder FAIL ----------------------------------"

	logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Info("---------------------------------- TaskHandler.UpdateOrder ----------------------------------")

	var taskMoveInfo dto.TaskMoveInfo
	err := json.NewDecoder(r.Body).Decode(&taskMoveInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("request struct decoded", requestID.String(), funcName, nodeName)

	err = th.ts.Move(rCtx, taskMoveInfo)
	if err != nil {
		logger.Error(errorMessage + err.Error())
		logger.Info(failBorder)
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	logger.DebugFmt("list order changed", requestID.String(), funcName, nodeName)

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

	logger.Info("---------------------------------- TaskHandler.UpdateOrder SUCCESS ----------------------------------")
}
