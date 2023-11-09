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
type Config struct {
	Session  *SessionConfig  `yaml:"-"`
	CORS     *CORSConfig     `yaml:"cors"`
	Database *DatabaseConfig `yaml:"db"`
}

type CORSConfig struct {
	AllowedMethods   []string `yaml:"allowed_methods"`
	AllowedHosts     []string `yaml:"allowed_hosts"`
	AllowedHeaders   []string `yaml:"allowed_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	Debug            bool     `yaml:"debug"`
}

type SessionConfig struct {
	Duration time.Duration `yaml:"-"`
	IDLength uint          `yaml:"-"`
}

type DatabaseConfig struct {
	User              string `yaml:"user"`
	Password          string `yaml:"-"`
	Host              string `yaml:"-"`
	Port              uint64 `yaml:"port"`
	DBName            string `yaml:"db_name"`
	AppName           string `yaml:"app_name"`
	Schema            string `yaml:"schema"`
	ConnectionTimeout uint64 `yaml:"connection_timeout"`
}

// LoadConfig
// создаёт конфиг из .env файла, находящегося по полученному пути
func LoadConfig(envPath string, configPath string) (*Config, error) {
	var (
		config Config
		err    error
	)

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	if envPath == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(envPath)
	}

	if err != nil {
		return nil, apperrors.ErrEnvNotFound
	}

	config.Session, err = NewSessionConfig()
	if err != nil {
		return nil, err
	}

	config.Database.Password, err = getDBPassword()
	if err != nil {
		return nil, err
	}

	config.Database.Host = getDBConnectionHost()

	return &config, nil
}

// NewSessionConfig
// создаёт конфиг сессии
func NewSessionConfig() (*SessionConfig, error) {
	var (
		config SessionConfig
		err    error
	)

	config.Duration, err = getSessionDurationEnv()
	if err != nil {
		return nil, err
	} else if config.Duration < time.Duration(1*time.Second) {
		return nil, apperrors.ErrSessionNullDuration
	}

	config.IDLength, err = getSessionIDLength()
	if err != nil {
		return nil, err
	} else if config.IDLength < 1 {
		return nil, apperrors.ErrSessionNullIDLength
	}

	return &config, nil
}

// getDBConnectionHost
// возвращает имя хоста из env для соединения с БД (по умолчанию localhost)
func getDBConnectionHost() string {
	host, hOk := os.LookupEnv("POSTGRES_HOST")
	if !hOk {
		return "localhost"
	}
	return host
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func getDBPassword() (string, error) {
	pwd, pOk := os.LookupEnv("POSTGRES_PASSWORD")
	if !pOk {
		return "", apperrors.ErrDatabasePWMissing
	}

	return pwd, nil
}

// getSessionDurationEnv
// возвращает время жизни сессии на основе параметров в .env (по умолчанию 14 дней)
func getSessionDurationEnv() (time.Duration, error) {
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

// getSessionIDLength
// возвращает время жизни сессии на основе параметров в .env (по умолчанию 14 дней)
func getSessionIDLength() (uint, error) {
	sidLengthString, ok := os.LookupEnv("SESSION_ID_LENGTH")
	if !ok {
		log.Println("WARNING: session ID length is not set, defaulting to 32")
		return uint(32), nil
	}

	sidLength64, err := strconv.ParseUint(sidLengthString, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(sidLength64), nil
}
