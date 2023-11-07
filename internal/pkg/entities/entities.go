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

// Workspace
// структура для хранения доски
type Workspace struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	DateCreated  string  `json:"date_created"`
	ThumbnailURL string  `json:"thumbnail_url"`
	Users        string  `json:"users"`
	Boards       []Board `json:"boards"`
}

// Board
// структура для хранения доски
type Board struct {
	ID           uint64         `json:"board_id"`
	Name         string         `json:"name"`
	Owner        dto.UserInfo   `json:"owner"`
	Description  string         `json:"description"`
	ThumbnailURL string         `json:"thumbnail_url"`
	Users        []dto.UserInfo `json:"user"`
	Lists        []List         `json:"lists"`
}

// List
// структура для хранения списка
type List struct {
	ID           uint64 `json:"id"`
	BoardID      uint64 `json:"board_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ListPosition uint64 `json:"list_position"`
	Tasks        []Task `json:"tasks"`
}

// Role
// структура для хранения роли пользователя
type Role struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Workspace
// структура для хранения доски
type Task struct {
	ID           uint64    `json:"id"`
	ListID       uint64    `json:"list_id"`
	DateCreated  string    `json:"date_created"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ListPosition uint64    `json:"list_position"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
}

func (u *User) TableName() string {
	return "user"
}
