package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища комментариев
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type ICommentStorage interface {
	// Create
	// создает новый комментарий в бд
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCommentInfo) (*entities.Comment, error)
	// GetFromTask
	// возвращает все комментарии у задания в БД
	// или возвращает ошибки ...
	GetFromTask(context.Context, dto.TaskID) (*[]dto.CommentInfo, error)
	// ReadMany
	// возвращает все комментарии по многим ID в БД
	// или возвращает ошибки ...
	ReadMany(context.Context, dto.CommentIDs) (*[]dto.CommentInfo, error)
}
