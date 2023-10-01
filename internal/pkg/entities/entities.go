package entities

import (
	"server/internal/pkg/dto"
	"time"
)

// TODO Split into DTO and Entity

type Session struct {
	UserID     uint64
	ExpiryDate time.Time
}

type User struct {
	ID           uint64 `json:"user_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
}

type Board struct {
	ID           uint64         `json:"board_id"`
	Name         string         `json:"name"`
	Owner        dto.UserInfo   `json:"owner"`
	ThumbnailURL string         `json:"thumbnail_url"`
	Guests       []dto.UserInfo `json:"guests"`
}
