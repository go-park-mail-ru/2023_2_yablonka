package config

import (
	"server/internal/apperrors"
	"server/internal/pkg/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		config        *entities.ServerConfig
		serviceType   string
		successful    bool
		expectedError error
	}{
		{
			name: "Valid config",
			config: &entities.ServerConfig{
				SessionDuration: time.Duration(30 * 24 * time.Hour),
				SessionIDLength: 32,
				JWTSecret:       "VALID JWT SECRET",
			},
			successful:    true,
			expectedError: nil,
		},
		{
			name: "Invalid config (Session duration is 0)",
			config: &entities.ServerConfig{
				SessionDuration: time.Duration(0 * time.Second),
				SessionIDLength: 32,
				JWTSecret:       "VALID JWT SECRET",
			},
			successful:    false,
			expectedError: apperrors.ErrSessionNullDuration,
		},
		{
			name: "Invalid config (Session ID length is 0)",
			config: &entities.ServerConfig{
				SessionDuration: time.Duration(30 * 24 * time.Hour),
				SessionIDLength: 0,
				JWTSecret:       "VALID JWT SECRET",
			},
			successful:    false,
			expectedError: apperrors.ErrSessionNullIDLength,
		},
		{
			name: "Invalid config (JWT secret is missing)",
			config: &entities.ServerConfig{
				SessionDuration: time.Duration(30 * 24 * time.Hour),
				SessionIDLength: 32,
				JWTSecret:       "",
			},
			successful:    false,
			expectedError: apperrors.ErrJWTSecretMissing,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			config := test.config

			ok, err := ValidateConfig(config)
			require.Equal(t, test.successful, ok)
			require.ErrorIs(t, test.expectedError, err)
		})
	}
}
