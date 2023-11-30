package csat

import (
	embedded "server/internal/service/csat/embedded"
	micro "server/internal/service/csat/microservice"
	"server/internal/storage"

	"google.golang.org/grpc"
)

func NewEmbeddedCSATQuestionService(quuestionStorage storage.ICSATQuestionStorage) *embedded.CSATQuestionService {
	return embedded.NewCSATQuestionService(quuestionStorage)
}

func NewMicroCSATQuestionService(quuestionStorage storage.ICSATQuestionStorage, connection *grpc.ClientConn) *micro.CSATQuestionService {
	return micro.NewCSATQuestionService(quuestionStorage, connection)
}
