package csat_microservice

import (
	logging "server/internal/logging"
	"server/internal/storage"
	answer "server/microservices/csat/csat_answer"
	question "server/microservices/csat/csat_question"

	"google.golang.org/grpc"
)

type Microservices struct {
	CSATQuestion question.CSATQuestionServiceServer
	CSATAnswer   answer.CSATAnswerServiceServer
}

func NewMicroServices(storages *storage.Storages) *Microservices {
	return &Microservices{
		CSATAnswer:   answer.NewCSATAnswerService(storages.CSATAnswer),
		CSATQuestion: question.NewCSATQuestionService(storages.CSATQuestion),
	}
}

func RegisterServices(storages *storage.Storages, server *grpc.Server, logger *logging.LogrusLogger) {
	answerServer := answer.NewCSATAnswerService(storages.CSATAnswer)
	questionServer := question.NewCSATQuestionService(storages.CSATQuestion)

	answer.RegisterCSATAnswerServiceServer(server, answerServer)
	question.RegisterCSATQuestionServiceServer(server, questionServer)
}
