package storage

import (
	// apperrors "server/internal/apperrors"

	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

// Интерфейс для хранилища досок
//
//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_storage/$GOFILE -package=mock_storage
type IBoardStorage interface {
	// GetUsers
	// находит пользователей, у которых есть доступ к доске
	GetUsers(context.Context, dto.BoardID) (*[]dto.UserPublicInfo, error)
	// GetById
	// находит доску и связанные с ней списки и задания по id
	GetById(context.Context, dto.BoardID) (*dto.SingleBoardInfo, error)
	// CheckAccess
	// находит пользователя в доске
	// или возвращает ошибки ...
	CheckAccess(context.Context, dto.CheckBoardAccessInfo) (bool, error)
	// GetLists
	// находит списки в доске
	// или возвращает ошибки ...
	GetLists(context.Context, dto.BoardID) (*[]dto.SingleListInfo, error)
	// GetTags
	// находит тэги в доске
	// или возвращает ошибки ...
	GetTags(context.Context, dto.BoardID) (*[]dto.TagInfo, error)
	// UpdateData
	// обновляет доску
	UpdateData(context.Context, dto.UpdatedBoardInfo) error
	// Update
	// обновляет картинку доски
	UpdateThumbnailUrl(context.Context, dto.BoardImageUrlInfo) error
	// Create
	// создает доску
	Create(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// Delete
	// удаляет доску
	Delete(context.Context, dto.BoardID) error
	// AddUser
	// добавляет пользователя на доску
	AddUser(context.Context, dto.AddBoardUserInfo) (dto.UserPublicInfo, error)
	// RemoveUser
	// удаляет пользователя с доски
	RemoveUser(context.Context, dto.RemoveBoardUserInfo) error
	// GetHistory
	// возвращает историю изменения доски
	GetHistory(context.Context, dto.BoardID) (*[]dto.BoardHistoryEntry, error)
	// SubmitEdit
	// записывает изменение доски в историю
	SubmitEdit(context.Context, dto.NewHistoryEntry) error
}
