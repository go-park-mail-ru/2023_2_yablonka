package csat

import (
	micro "server/internal/service/csat/microservice"
	"server/internal/storage"

	"google.golang.org/grpc"
)

func NewMicroCSATQuestionService(quuestionStorage storage.ICSATQuestionStorage, connection *grpc.ClientConn) *micro.CSATQuestionService {
	return micro.NewCSATQuestionService(quuestionStorage, connection)
}
