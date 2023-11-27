package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/internal/apperrors"
	_ "server/internal/pkg/doc_structs"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/sirupsen/logrus"
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
	funcName := "Create"

	logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)
	logger.Info("----------------- Creating a new task -----------------")

	var newTaskInfo dto.NewTaskInfo
	err := json.NewDecoder(r.Body).Decode(&newTaskInfo)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new task failed with error "+err.Error())
		logger.Error("----------------- Creating a new task FAIL -----------------")
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON Decoded")

	// _, err = govalidator.ValidateStruct(newTaskInfo)
	// if err != nil {
	// 	logger.Error("Creating a new board failed")
	// 	handlerDebugLog(logger, funcName, "Creating a new board failed with error "+err.Error())
	// 	apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
	// 	return
	// }
	// handlerDebugLog(logger, funcName, "New task data validated")

	task, err := th.ts.Create(rCtx, newTaskInfo)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new task failed with error "+err.Error())
		logger.Error("----------------- Creating a new task FAIL -----------------")
		apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
		return
	}
	handlerDebugLog(logger, funcName, "Task created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"task": task,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new task failed with error "+err.Error())
		logger.Error("----------------- Creating a new task FAIL -----------------")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	handlerDebugLog(logger, funcName, "JSON response marshaled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		handlerDebugLog(logger, funcName, "Creating a new task failed with error "+err.Error())
		logger.Error("----------------- Creating a new task FAIL -----------------")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()

	handlerDebugLog(logger, funcName, "Response written")
	logger.Info("----------------- Creating a new task SUCCESS -----------------")
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
	log.Println("--------------TaskHandler.Read Endpoint START--------------")

	rCtx := r.Context()

	log.Println("Handler -- Reading task")

	var taskID dto.TaskID
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Read Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(taskID)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	task, err := th.ts.Read(rCtx, taskID)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Read Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("task read")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"task": task,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Read Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Read Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------TaskHandler.Read Endpoint SUCCESS--------------")
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
	log.Println("--------------TaskHandler.Update Endpoint START--------------")

	rCtx := r.Context()

	var taskInfo dto.UpdatedTaskInfo
	err := json.NewDecoder(r.Body).Decode(&taskInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(taskInfo)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	err = th.ts.Update(rCtx, taskInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("task updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------TaskHandler.Update Endpoint SUCCESS--------------")
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
	log.Println("--------------TaskHandler.Delete Endpoint START--------------")

	rCtx := r.Context()

	var taskID dto.TaskID
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(taskID)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	err = th.ts.Delete(rCtx, taskID)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("task deleted")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------TaskHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------TaskHandler.Delete Endpoint SUCCESS--------------")
}
