package handlers

import (
	"net/http"
	_ "server/internal/pkg/doc_structs"
	"server/internal/service"
)

type CSATQuestionHandler struct {
	qs service.ICSATQuestionService
}

// @Summary Получить вопросы CSAT
// @Description Получить вопросы CSAT
// @Tags csat
//
// @Accept  json
// @Produce  json
//
// @Success 204  {string} string "no content"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /csat/question/all [get]
func (ah CSATQuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
