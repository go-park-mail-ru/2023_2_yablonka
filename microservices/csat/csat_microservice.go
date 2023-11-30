package csat_microservice

import (
	logging "server/internal/logging"
	"server/internal/storage"
	answer "server/microservices/csat/csat_answer"
	question "server/microservices/csat/csat_question"

	"google.golang.org/grpc"
)

const nodeName = "microservice"

func RegisterServices(storages *storage.Storages, server *grpc.Server, logger *logging.LogrusLogger) {
	funcName := "CSAT.RegisterService"
	answerServer := answer.NewCSATAnswerService(storages.CSATAnswer)
	logger.Debug("CSAT answer GRPC service created", funcName, nodeName)
	questionServer := question.NewCSATQuestionService(storages.CSATQuestion)
	logger.Debug("CSAT question GRPC service created", funcName, nodeName)

	answer.RegisterCSATAnswerServiceServer(server, answerServer)
	logger.Debug("CSAT answer GRPC service registered", funcName, nodeName)
	question.RegisterCSATQuestionServiceServer(server, questionServer)
	logger.Debug("CSAT question GRPC service registered", funcName, nodeName)
}
