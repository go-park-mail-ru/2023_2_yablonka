package config

import (
	"log"
	"os"
	"server/internal/apperrors"
	"server/internal/config"
	"strconv"
)

type SessionServerConfig struct {
	SessionIDLength uint
	Base            config.BaseServerConfig
}

func (config *SessionServerConfig) Validate() error {
	err := config.Base.Validate()
	if err != nil {
		return err
	}
	if config.SessionIDLength < 1 {
		return apperrors.ErrSessionNullIDLength
	}
	return nil
}

func NewSessionEnvConfig(filepath string) (*SessionServerConfig, error) {
	baseConfig, err := config.NewBaseEnvConfig(filepath)

	if err != nil {
		return nil, err
	}

	var sidLength uint
	sidLengthString, ok := os.LookupEnv("SESSION_ID_LENGTH")
	if !ok {
		sidLength = uint(32)
		log.Println("WARNING: session ID length is not set, defaulting to 32")
	} else {
		sidLength64, err := strconv.ParseUint(sidLengthString, 10, 32)
		sidLength = uint(sidLength64)
		if err != nil {
			return nil, err
		}
	}

	return &SessionServerConfig{
		Base:            *baseConfig,
		SessionIDLength: sidLength,
	}, nil
}
