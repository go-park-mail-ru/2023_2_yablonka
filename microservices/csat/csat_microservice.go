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
	answerServer := answer.NewCSATAnswerService(storages.CSATAnswer, logger)
	logger.DebugRequestlessFmt("CSAT answer GRPC service created", funcName, nodeName)
	questionServer := question.NewCSATQuestionService(storages.CSATQuestion, logger)
	logger.DebugRequestlessFmt("CSAT question GRPC service created", funcName, nodeName)

	answer.RegisterCSATAnswerServiceServer(server, answerServer)
	logger.DebugRequestlessFmt("CSAT answer GRPC service registered", funcName, nodeName)
	question.RegisterCSATQuestionServiceServer(server, questionServer)
	logger.DebugRequestlessFmt("CSAT question GRPC service registered", funcName, nodeName)
}
