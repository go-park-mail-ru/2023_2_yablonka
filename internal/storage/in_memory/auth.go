package in_memory

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"server/internal/apperrors"
	"server/internal/pkg/entities"
	"sync"
)

// LocalAuthStorage
// локальное хранилище сессий
type LocalAuthStorage struct {
	authData map[string]*entities.Session
	mu       *sync.RWMutex
}

// func getSidLength() (uint, error) {
// 	sidLengthStr, ok := os.LookupEnv("SESSION_ID_LENGTH")
// 	if !ok {
// 		return 0, apperrors.ErrSessionIDLengthMissing
// 	}
// 	sidLength, _ := strconv.Atoi(sidLengthStr)
// 	if sidLength < 1 {
// 		return 0, apperrors.ErrSessionNullIDLength
// 	}
// 	return uint(sidLength), nil
// }

// NewLocalAuthStorage
// возвращает локальное хранилище сессий
func NewAuthStorage() *LocalAuthStorage {
	return &LocalAuthStorage{
		authData: map[string]*entities.Session{},
		mu:       &sync.RWMutex{},
	}
}

// CreateSession
// сохраняет сессию в хранилище, возвращает ID сесссии для куки
func (as LocalAuthStorage) CreateSession(ctx context.Context, session *entities.Session) (string, error) {
	// for sessionID, storedSession := range as.authData {
	// 	if storedSession.UserID == session.UserID {
	// 		return "", apperrors.ErrSessionExists
	// 	}
	// }
	sessionIDLength := ctx.Value("sessionIDLength").(uint)
	sessionID, err := generateSessionID(sessionIDLength)
	if err != nil {
		return "", err
	}
	as.mu.Lock()
	as.authData[sessionID] = session
	as.mu.Unlock()
	return sessionID, nil
}

// CreateSession
// сохраняет сессию в хранилище, возвращает ID сесссии для куки
func (as LocalAuthStorage) GetSession(ctx context.Context, sid string) (*entities.Session, error) {
	as.mu.RLock()
	session, ok := as.authData[sid]
	as.mu.RUnlock()
	if !ok {
		fmt.Println("Session not found")
		return nil, apperrors.ErrSessionNotFound
	}
	fmt.Println("Session found")
	return session, nil
}

// DeleteSession
// удаляет сессию по ID из хранилища, если она существует
func (as LocalAuthStorage) DeleteSession(ctx context.Context, sid string) error {
	as.mu.RLock()
	_, ok := as.authData[sid]
	as.mu.RUnlock()
	if !ok {
		return apperrors.ErrSessionNotFound
	}
	as.mu.Lock()
	as.authData[sid] = nil
	delete(as.authData, sid)
	as.mu.Unlock()
	return nil
}

// GenerateSessionID
// возвращает alphanumeric строку, собранную криптографически безопасным PRNG
func generateSessionID(n uint) (string, error) {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	buf := make([]rune, n)
	for i := range buf {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			return "", err
		}
		buf[i] = letterRunes[j.Int64()]
	}
	return string(buf), nil
}
