package comment

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"

	embedded "server/internal/service/comment/embedded"
	micro "server/internal/service/comment/microservice"

	"google.golang.org/grpc"
)

// Интерфейс для сервиса комментариев
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type ICommentService interface {
	// Create
	// создает новый комментарий
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCommentInfo) (*entities.Comment, error)
	// GetFromTask
	// возвращает все комментарии у задания
	// или возвращает ошибки ...
	GetFromTask(context.Context, dto.TaskID) (*[]dto.CommentInfo, error)
}

func NewEmbeddedCommentService(commentStorage storage.ICommentStorage) *embedded.CommentService {
	return embedded.NewCommentService(commentStorage)
}

// TODO: Board microservice
func NewMicroCommentService(commentStorage storage.ICommentStorage, connection *grpc.ClientConn) *micro.CommentService {
	return micro.NewCommentService(commentStorage)
}
