package handlers

import (
	_ "server/internal/pkg/doc_structs"
	"server/internal/service"
)

type CSATQuestionHandler struct {
	qs service.ICSATQuestionService
}
