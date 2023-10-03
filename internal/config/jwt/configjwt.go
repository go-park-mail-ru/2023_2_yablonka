package config

import (
	"os"
	"server/internal/apperrors"
	"server/internal/config"
)

// ServerConfig
// структура для хранения параметров сервера
type JWTServerConfig struct {
	JWTSecret string
	Base      config.BaseServerConfig
}

func (config *JWTServerConfig) Validate() error {
	err := config.Base.Validate()
	if err != nil {
		return err
	}
	if config.JWTSecret == "" {
		return apperrors.ErrJWTSecretMissing
	}
	return nil
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
