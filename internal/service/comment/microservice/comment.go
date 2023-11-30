package microservice

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
)

type CommentService struct {
	storage storage.ICommentStorage
}

// NewBoardService
// возвращает BoardService с инициализированным хранилищем
func NewCommentService(storage storage.ICommentStorage) *CommentService {
	return &CommentService{
		storage: storage,
	}
}

// GetFromTask
// возвращает все комментарии у задания
// или возвращает ошибки ...
func (cs CommentService) GetFromTask(ctx context.Context, id dto.TaskID) (*[]dto.CommentInfo, error) {
	return cs.storage.GetFromTask(ctx, id)
}

// Create
// создает новый комментарий
// или возвращает ошибки ...
func (cs CommentService) Create(ctx context.Context, info dto.NewCommentInfo) (*entities.Comment, error) {
	return cs.storage.Create(ctx, info)
}
