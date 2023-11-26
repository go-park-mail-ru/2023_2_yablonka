package entities

import (
	"server/internal/pkg/dto"
	"time"
)

// Session
// структура для хранения сессии
type Session struct {
	SessionID  string
	UserID     uint64
	ExpiryDate time.Time
}

// CSRF
// структура для хранения сессии CSRF
type CSRF struct {
	Token          string
	UserID         uint64
	ExpirationDate time.Time
}

// User
// структура для хранения пользователя
type User struct {
	ID           uint64  `json:"user_id"`
	Email        string  `json:"email" valid:"type(string),email"`
	PasswordHash string  `json:"password_hash"`
	Name         *string `json:"name"`
	Surname      *string `json:"surname"`
	AvatarURL    *string `json:"avatar_url"`
	Description  *string `json:"description"`
}

// Workspace
// структура для хранения доски
type Workspace struct {
	ID          uint64               `json:"id"`
	Name        string               `json:"name"`
	Description *string              `json:"description"`
	DateCreated time.Time            `json:"date_created"`
	Users       []dto.UserPublicInfo `json:"users"`
	Boards      []Board              `json:"boards"`
}

// Board
// структура для хранения доски
type Board struct {
	ID           uint64               `json:"board_id"`
	Name         string               `json:"name"`
	Owner        dto.UserID           `json:"owner"`
	ThumbnailURL *string              `json:"thumbnail_url"`
	DateCreated  time.Time            `json:"date_created"`
	Users        []dto.UserPublicInfo `json:"users"`
	Lists        []List               `json:"lists"`
}

// List
// структура для хранения списка
type List struct {
	ID           uint64  `json:"id"`
	BoardID      uint64  `json:"board_id"`
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	ListPosition uint64  `json:"list_position"`
	Tasks        []Task  `json:"tasks"`
}

// QuestionType
// структура для хранения типа вопроса
type QuestionType struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	MaxRating uint64 `json:"max_rating"`
}

// CSATQuestion
// структура для CSAT ответа
type CSATQuestion struct {
	ID      uint64 `json:"id"`
	TypeID  string `json:"type"`
	Content string `json:"content"`
}

// Role
// структура для хранения роли пользователя
type Role struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// Workspace
// структура для хранения доски
type Task struct {
	ID           uint64      `json:"id"`
	ListID       uint64      `json:"list_id"`
	DateCreated  time.Time   `json:"date_created"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	ListPosition uint64      `json:"list_position"`
	Start        *time.Time  `json:"start"`
	End          *time.Time  `json:"end"`
	Users        []uint64    `json:"users"`
	Checklists   []Checklist `json:"checklists"`
	Comments     []Comment   `json:"comments"`
}

type Comment struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	TaskID      uint64    `json:"task_id"`
	Text        string    `json:"text"`
	DateCreated time.Time `json:"date_created"`
}

type Checklist struct {
	ID           uint64          `json:"id"`
	TaskID       uint64          `json:"task_id"`
	Name         string          `json:"name"`
	ListPosition uint64          `json:"list_position"`
	Items        []ChecklistItem `json:"items"`
}

type ChecklistItem struct {
	ID           uint64 `json:"id"`
	ChecklistID  uint64 `json:"checklist_id"`
	Name         string `json:"name"`
	Done         bool   `json:"done"`
	ListPosition uint64 `json:"list_position"`
}

func (u *User) TableName() string {
	return "user"
}
