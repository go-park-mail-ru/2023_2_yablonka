package storage

import (
	// apperrors "server/internal/apperrors"

	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardStorage interface {
	// GetById
	// находит доску и связанные с ней списки и задания по id
	// или возвращает ошибки ...
	GetById(context.Context, dto.BoardID) (*entities.Board, error)
	// UpdateData
	// обновляет доску
	// или возвращает ошибки ...
	UpdateData(context.Context, dto.UpdatedBoardInfo) error
	// Update
	// обновляет картинку доски
	// или возвращает ошибки ...
	UpdateThumbnailUrl(context.Context, dto.ImageUrlInfo) error
	// Create
	// создает доску
	// или возвращает ошибки ...
	Create(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// Delete
	// удаляет доску
	// или возвращает ошибки ...
	Delete(context.Context, dto.BoardID) error
}
