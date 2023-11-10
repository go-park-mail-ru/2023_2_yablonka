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

	"github.com/asaskevich/govalidator"
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
	log.Println("Handler -- Creating task")
	rCtx := r.Context()

	var newTaskInfo dto.NewTaskInfo
	err := json.NewDecoder(r.Body).Decode(&newTaskInfo)
	if err != nil {
		log.Println("Failed to decode JSON")
		log.Println("Error:", err.Error())
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	// _, err = govalidator.ValidateStruct(newTaskInfo)
	// if err != nil {
	// 	log.Println("Failed to validate struct")
	// 	log.Println("Error:", err.Error())
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	task, err := th.ts.Create(rCtx, newTaskInfo)
	if err != nil {
		log.Println("Failed to create task")
		log.Println("Error:", err.Error())
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"list": task,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
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
// @Router /task [get]
func (th TaskHandler) Read(w http.ResponseWriter, r *http.Request) {
	rCtx := r.Context()

	var taskID dto.TaskID
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(taskID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	task, err := th.ts.Read(rCtx, taskID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"task": task,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
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

	var taskInfo dto.UpdatedTaskInfo
	err := json.NewDecoder(r.Body).Decode(&taskInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(taskInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	err = th.ts.Update(rCtx, taskInfo)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
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

	var taskID dto.TaskID
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	_, err = govalidator.ValidateStruct(taskID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}

	err = th.ts.Delete(rCtx, taskID)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}

	r.Body.Close()
}
