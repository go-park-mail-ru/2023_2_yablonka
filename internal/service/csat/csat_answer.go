package csat

import (
	micro "server/internal/service/csat/microservice"
	"server/internal/storage"

	"google.golang.org/grpc"
)

func NewMicroCSATAnswerService(answerStorage storage.ICSATAnswerStorage, connection *grpc.ClientConn) *micro.CSATAnswerService {
	return micro.NewCSATAnswerService(answerStorage, connection)
}
