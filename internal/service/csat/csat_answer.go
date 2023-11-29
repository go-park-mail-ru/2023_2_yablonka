package csat

import (
	embedded "server/internal/service/csat/embedded"
	micro "server/internal/service/csat/microservice"
	"server/internal/storage"

	"google.golang.org/grpc"
)

func NewEmbeddedCSATAnswerService(answerStorage storage.ICSATAnswerStorage) *embedded.CSATAnswerService {
	return embedded.NewCSATAnswerService(answerStorage)
}

func NewMicroCSATAnswerService(answerStorage storage.ICSATAnswerStorage, connection *grpc.ClientConn) *micro.CSATAnswerService {
	return micro.NewCSATAnswerService(answerStorage, connection)
}
