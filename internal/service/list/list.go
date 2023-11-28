package list

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	embedded "server/internal/service/list/embedded"
	micro "server/internal/service/list/microservice"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса списков заданий
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IListService interface {
	// Create
	// создает новый список
	// или возвращает ошибки ...
	Create(context.Context, dto.NewListInfo) (*entities.List, error)
	// Update
	// обновляет список
	// или возвращает ошибки ...
	Update(context.Context, dto.UpdatedListInfo) error
	// Delete
	// удаляет список по id
	// или возвращает ошибки ...
	Delete(context.Context, dto.ListID) error
}

func NewEmbeddedListService(listStorage storage.IListStorage) *embedded.ListService {
	return embedded.NewListService(listStorage)
}

// TODO: List microservice
func NewMicroListService(listStorage storage.IListStorage, connection *grpc.ClientConn) *micro.ListService {
	return micro.NewListService(listStorage)
}
