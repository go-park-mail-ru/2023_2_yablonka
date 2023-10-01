package in_memory

import (
	"context"
	"os"
	"server/internal/app/utils"
	"server/internal/apperrors"
	"server/internal/pkg/entities"
	"strconv"
	"sync"
)

// LocalAuthStorage
// Локальное хранилище сессий
type LocalAuthStorage struct {
	authData  map[string]*entities.Session
	sidLength uint
	mu        *sync.RWMutex
}

func getSidLength() (uint, error) {
	sidLengthStr, ok := os.LookupEnv("SESSION_ID_LENGTH")
	if !ok {
		return 0, apperrors.ErrSessionIDLengthMissing
	}
	sidLength, _ := strconv.Atoi(sidLengthStr)
	if sidLength < 1 {
		return 0, apperrors.ErrSessionNullIDLength
	}
	return uint(sidLength), nil
}

// NewLocalAuthStorage
// Возвращает локальное хранилище данных с инициализированным параметром длины ID сессии
func NewAuthStorage() (*LocalAuthStorage, error) {
	sidLength, err := getSidLength()
	if err != nil {
		return nil, err
	}
	return &LocalAuthStorage{
		authData:  map[string]*entities.Session{},
		sidLength: sidLength,
		mu:        &sync.RWMutex{},
	}, nil
}

func (as LocalAuthStorage) CreateSession(ctx context.Context, session *entities.Session) (string, error) {
	sessionID, err := utils.GenerateSessionID(as.sidLength)
	if err != nil {
		return "", err
	}
	as.authData[sessionID] = session
	return sessionID, nil
}

func (as LocalAuthStorage) GetSession(ctx context.Context, sid string) (*entities.Session, error) {
	session, ok := as.authData[sid]
	if !ok {
		return nil, apperrors.ErrSessionNotFound
	}
	return session, nil
}

func (as LocalAuthStorage) DeleteSession(ctx context.Context, user *entities.User) error {
	as.authData[user.Email] = nil
	return nil
}
