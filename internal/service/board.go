package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IBoardService interface {
	// GetFullBoard
	// возвращает доску со связанными пользователями, списками и заданиями
	GetFullBoard(context.Context, dto.IndividualBoardRequest) (*dto.FullBoardResult, error)
	// Create
	// создаёт доску и связь с пользователем-создателем
	Create(context.Context, dto.NewBoardInfo) (*entities.Board, error)
	// UpdateData
	// возвращает доску со связанными пользователями, списками и заданиями
	UpdateData(context.Context, dto.UpdatedBoardInfo) error
	// UpdateThumbnail
	// сохраняет картинку доски в папку images/board_thumbnails с названием id доски и сохраняет ссылку на изображение в БД
	UpdateThumbnail(context.Context, dto.UpdatedBoardThumbnailInfo) (*dto.UrlObj, error)
	// Delete
	// удаляет доску
	Delete(context.Context, dto.BoardDeleteRequest) error
	// AddUser
	// добавляет пользователя на доску
	AddUser(context.Context, dto.AddBoardUserRequest) (dto.UserPublicInfo, error)
	// AddUser
	// добавляет пользователя на доску
	RemoveUser(context.Context, dto.RemoveBoardUserInfo) error
	// GetHistory
	// возвращает историю изменения доски
	GetHistory(context.Context, dto.BoardID) (*[]dto.BoardHistoryEntry, error)
	// SubmitEdit
	// записывает изменение доски в историю
	SubmitEdit(context.Context, dto.NewHistoryEntry) error
}
