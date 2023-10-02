package config

import (
	"os"
	"server/internal/app/config"
	"server/internal/apperrors"
)

// ServerConfig
// структура для хранения параметров сервера
type JWTServerConfig struct {
	JWTSecret string
	Base      config.BaseServerConfig
}

func (config *JWTServerConfig) Validate() (bool, error) {
	base, err := config.Base.Validate()
	if err != nil {
		return base, err
	}
	if config.JWTSecret == "" {
		return false, apperrors.ErrJWTSecretMissing
	}
	return true, nil
}

func NewJWTEnvConfig(filepath string) (*JWTServerConfig, error) {
	baseConfig, err := config.NewBaseEnvConfig(filepath)

	if err != nil {
		return nil, err
	}

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok || jwtSecret == "" {
		return nil, apperrors.ErrJWTSecretMissing
	}

	return &JWTServerConfig{
		Base:      *baseConfig,
		JWTSecret: jwtSecret,
	}, nil
}
