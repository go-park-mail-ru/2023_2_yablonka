package config

import (
	"log"
	"os"
	"server/internal/apperrors"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// ServerConfig
// структура для хранения параметров сервера
type BaseServerConfig struct {
	SessionDuration time.Duration
}

// ServerConfig
// структура для хранения параметров сервера
type ServerConfig interface {
	Validate() error
}

// Validate
// проверяет параметры конфига на валидность
func (config *BaseServerConfig) Validate() error {
	if config.SessionDuration < time.Duration(1*time.Second) {
		return apperrors.ErrSessionNullDuration
	}
	return nil
}

// NewJWTEnvConfig
// создаёт конфиг из .env файла, находящегося по полученному пути
func NewBaseEnvConfig(filepath string) (*BaseServerConfig, error) {
	var err error
	if filepath == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(filepath)
	}

	if err != nil {
		return nil, apperrors.ErrEnvNotFound
	}

	sessionDuration, err := buildSessionDurationEnv()
	if err != nil {
		return nil, err
	}

	return &BaseServerConfig{
		SessionDuration: sessionDuration,
	}, nil
}

// buildSessionDurationEnv
// возвращает время жизни сессии на основе параметров в .env (по умолчанию 14 дней)
func buildSessionDurationEnv() (time.Duration, error) {
	durationDays, dDok := os.LookupEnv("SESSION_DURATION_DAYS")
	days, _ := strconv.Atoi(durationDays)
	durationHours, dHok := os.LookupEnv("SESSION_DURATION_HOURS")
	hours, _ := strconv.Atoi(durationHours)
	durationMinutes, dMok := os.LookupEnv("SESSION_DURATION_MINUTES")
	minutes, _ := strconv.Atoi(durationMinutes)
	durationSeconds, dSok := os.LookupEnv("SESSION_DURATION_SECONDS")
	seconds, _ := strconv.Atoi(durationSeconds)
	if !(dDok || dHok || dMok || dSok) {
		log.Println("WARNING: session duration is not set, defaulting to 14 days")
		return time.Duration(14 * 24 * time.Hour), nil
	}
	totalDuration := time.Duration(24*days+hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	if totalDuration < time.Second {
		return 0, apperrors.ErrSessionNullDuration
	}
	return totalDuration, nil
}
