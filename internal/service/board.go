package service

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
)

type IBoardService interface {
	// GetFullBoard
	// возвращает доску со связанными пользователями, списками и заданиями
	GetFullBoard(context.Context, dto.IndividualBoardRequest) (*entities.Board, error)
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
	Delete(context.Context, dto.BoardID) error
}
