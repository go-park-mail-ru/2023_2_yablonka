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
)

type ListHandler struct {
	ls service.IListService
}

// @Summary Создать список
// @Description Создать список
// @Tags lists
//
// @Accept  json
// @Produce  json
//
// @Param newListInfo body dto.NewListInfo true "данные нового списка"
//
// @Success 200  {object}  doc_structs.ListResponse "объект списка"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /list/create/ [post]
func (lh ListHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("--------------ListHandler.Create Endpoint START--------------")

	rCtx := r.Context()

	var newListInfo dto.NewListInfo
	err := json.NewDecoder(r.Body).Decode(&newListInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	list, err := lh.ls.Create(rCtx, newListInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("list created")

	response := dto.JSONResponse{
		Body: dto.JSONMap{
			"list": list,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Create Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------ListHandler.Create Endpoint SUCCESS--------------")
}

// @Summary Обновить список
// @Description Обновить список
// @Tags lists
//
// @Accept  json
// @Produce  json
//
// @Param listInfo body dto.UpdatedListInfo true "обновленные данные списка"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /list/update/ [post]
func (lh ListHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("--------------ListHandler.Update Endpoint START--------------")

	rCtx := r.Context()

	var listInfo dto.UpdatedListInfo
	err := json.NewDecoder(r.Body).Decode(&listInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(listInfo)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	err = lh.ls.Update(rCtx, listInfo)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("list updated")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Update Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------ListHandler.Update Endpoint SUCCESS--------------")
}

// @Summary Удалить список
// @Description Удалить список
// @Tags lists
//
// @Accept  json
// @Produce  json
//
// @Param listID body dto.ListID true "id списка"
//
// @Success 204  {string}  string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /list/delete/ [delete]
func (lh ListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("--------------ListHandler.Delete Endpoint START--------------")

	rCtx := r.Context()

	var listID dto.ListID
	err := json.NewDecoder(r.Body).Decode(&listID)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
		return
	}
	log.Println("request struct decoded")

	// _, err = govalidator.ValidateStruct(listID)
	// if err != nil {
	// 	*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.BadRequestResponse))
	// 	return
	// }

	err = lh.ls.Delete(rCtx, listID)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
		return
	}
	log.Println("list deleted")

	response := dto.JSONResponse{
		Body: dto.JSONMap{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	log.Println("json response marshalled")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		log.Println("--------------ListHandler.Delete Endpoint FAIL--------------")
		*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.InternalServerErrorResponse))
		return
	}
	r.Body.Close()
	log.Println("response written")

	log.Println("--------------ListHandler.Delete Endpoint SUCCESS--------------")
}
