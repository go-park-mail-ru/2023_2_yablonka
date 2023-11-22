package storage

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type ICommentStorage interface {
	// Create
	// создает новый комментарий в бд
	// или возвращает ошибки ...
	Create(context.Context, dto.NewCommentInfo) (*entities.Comment, error)
	// GetFromTask
	// возвращает все комментарии у задания в БД
	// или возвращает ошибки ...
	GetFromTask(context.Context, dto.TaskID) (*[]dto.CommentInfo, error)
}
