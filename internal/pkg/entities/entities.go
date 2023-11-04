package entities

import (
	"server/internal/pkg/dto"
	"time"
)

// Session
// структура для хранения сессии
type Session struct {
	Token      string
	UserID     uint64
	ExpiryDate time.Time
}

// User
// структура для хранения пользователя
type User struct {
	ID           uint64 `json:"user_id"`
	Email        string `json:"email" valid:"type(string),email"`
	PasswordHash string `json:"-"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	AvatarURL    string `json:"avatar_url"`
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
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======

// Board
// структура для хранения доски
type Workspace struct {
	ID           uint64 `json:"board_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DateCreated  string `json:"date_created"`
	ThumbnailURL string `json:"thumbnail_url"`
	Users        string `json:"users"`
	Boards       string `json:"boards"`
}

func (u *User) TableName() string {
	return "user"
}
>>>>>>> Stashed changes
>>>>>>> Stashed changes
