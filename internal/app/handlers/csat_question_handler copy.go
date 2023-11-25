package handlers

import (
	"net/http"
	_ "server/internal/pkg/doc_structs"
	"server/internal/service"
)

type CSATQuestionHandler struct {
	qs service.ICSATQuestionService
}

func (ah CSATQuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (ah CSATQuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
