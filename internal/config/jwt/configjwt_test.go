package config

import (
	"server/internal/apperrors"
	"server/internal/config"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_JWTConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		config        *JWTServerConfig
		serviceType   string
		expectedError error
	}{
		{
			name: "Invalid config (JWT secret is missing)",
			config: &JWTServerConfig{
				Base: config.BaseServerConfig{
					SessionDuration: time.Duration(30 * 24 * time.Hour),
				},
				JWTSecret: "",
			},
			expectedError: apperrors.ErrJWTSecretMissing,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			config := test.config

			err := config.Validate()
			require.ErrorIs(t, test.expectedError, err)
		})
	}
}
