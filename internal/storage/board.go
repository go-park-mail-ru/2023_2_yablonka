package storage

import (
	// apperrors "server/internal/apperrors"

	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardStorage interface {
	// GetUsers
	// находит пользователей, у которых есть доступ к доске
	GetUsers(context.Context, dto.BoardID) (*[]dto.UserPublicInfo, error)
	// GetById
	// находит доску и связанные с ней списки и задания по id
	GetById(context.Context, dto.BoardID) (*dto.FullBoardResult, error)
	// UpdateData
	// обновляет доску
	UpdateData(context.Context, dto.UpdatedBoardInfo) error
	// Update
	// обновляет картинку доски
	UpdateThumbnailUrl(context.Context, dto.ImageUrlInfo) error
	// Create
	// создает доску
	Create(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// Delete
	// удаляет доску
	Delete(context.Context, dto.BoardID) error
	// AddUser
	// добавляет пользователя на доску
	AddUser(context.Context, dto.AddBoardUserInfo) error
	// RemoveUser
	// удаляет пользователя с доски
	RemoveUser(context.Context, dto.RemoveBoardUserInfo) error
}
