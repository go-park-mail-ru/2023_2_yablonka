package service

import (
	"context"
	"server/internal/pkg/dto"
)

//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_service/$GOFILE -package=mock_service
type IFileService interface {
	// Upload
	// загружает изображение на сервер
	Upload(context.Context, dto.IndividualBoardRequest) (*dto.FullBoardResult, error)
	// Download
	// выгружает изображение с сервера
	Download(context.Context, dto.NewBoardInfo) (*dto.Image, error)
}
