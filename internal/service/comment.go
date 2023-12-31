package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
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
