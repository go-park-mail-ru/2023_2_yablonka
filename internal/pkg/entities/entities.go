package entities

import (
	"server/internal/pkg/dto"
	"time"
)

// Session
// структура для хранения сессии
type Session struct {
	UserID     uint64
	ExpiryDate time.Time
}

// User
// структура для хранения пользователя
type User struct {
	ID           uint64 `json:"user_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	ThumbnailURL string `json:"thumbnail_url"`
}

// Board
// структура для хранения доски
type Board struct {
	ID           uint64         `json:"board_id"`
	Name         string         `json:"name"`
	Owner        dto.UserInfo   `json:"owner"`
	ThumbnailURL string         `json:"thumbnail_url"`
	Guests       []dto.UserInfo `json:"guests"`
}

// Board
// структура для хранения доски
type Workspace struct {
	ID          uint64         `json:"board_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	DateCreated string         `json:"thumbnail_url"`
	Guests      []dto.UserInfo `json:"guests"`
}

func (u *User) TableName() string {
	return "user"
}
