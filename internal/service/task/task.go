package task

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	embedded "server/internal/service/task/embedded"
	micro "server/internal/service/task/microservice"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса заданий
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ITaskService interface {
	// Create
	// создает новое задание
	// или возвращает ошибки ...
	Create(context.Context, dto.NewTaskInfo) (*entities.Task, error)
	// Read
	// возвращает заание с привязанными пользователями
	// или возвращает ошибки ...
	Read(context.Context, dto.TaskID) (*entities.Task, error)
	// Update
	// обновляет задание
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedTaskInfo) error
	// Delete
	// удаляет задание
	// или возвращает ошибки ...
	Delete(context.Context, dto.TaskID) error
}

func NewEmbeddedTaskService(taskStorage storage.ITaskStorage) *embedded.TaskService {
	return embedded.NewTaskService(taskStorage)
}

// TODO: Board microservice
func NewMicroTaskService(taskStorage storage.ITaskStorage, connection *grpc.ClientConn) *micro.TaskService {
	return micro.NewTaskService(taskStorage)
}
