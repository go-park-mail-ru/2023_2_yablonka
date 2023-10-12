package config

import (
	"log"
	"os"
	"server/internal/apperrors"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	yaml "gopkg.in/yaml.v2"
)

// ServerConfig
// структура для хранения параметров сервера
type BaseServerConfig struct {
	SessionDuration time.Duration `yaml:"-"`
	Server          struct {
		AllowedMethods   []string `yaml:"allowed_methods"`
		AllowedHosts     []string `yaml:"allowed_hosts"`
		AllowedHeaders   []string `yaml:"allowed_headers"`
		AllowCredentials bool     `yaml:"allow_credentials"`
	} `yaml:"server"`
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
func NewBaseEnvConfig(envPath string, configPath string) (*BaseServerConfig, error) {
	var err error
	if envPath == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(envPath)
	}

	if err != nil {
		return nil, apperrors.ErrEnvNotFound
	}

	sessionDuration, err := buildSessionDurationEnv()
	if err != nil {
		return nil, err
	}

	config := &BaseServerConfig{
		SessionDuration: sessionDuration,
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
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
