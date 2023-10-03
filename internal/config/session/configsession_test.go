package config

import (
	"server/internal/apperrors"
	"server/internal/config"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_SessionConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		config        *SessionServerConfig
		serviceType   string
		expectedError error
	}{
		{
			name: "Valid config",
			config: &SessionServerConfig{
				Base: config.BaseServerConfig{
					SessionDuration: time.Duration(30 * 24 * time.Hour),
				},
				SessionIDLength: 32,
			},
			expectedError: nil,
		},
		{
			name: "Invalid config (Session duration is 0)",
			config: &SessionServerConfig{
				Base: config.BaseServerConfig{
					SessionDuration: time.Duration(0 * time.Second),
				},
				SessionIDLength: 32,
			},
			expectedError: apperrors.ErrSessionNullDuration,
		},
		{
			name: "Invalid config (Session ID length is 0)",
			config: &SessionServerConfig{
				Base: config.BaseServerConfig{
					SessionDuration: time.Duration(30 * 24 * time.Hour),
				},
				SessionIDLength: 0,
			},
			expectedError: apperrors.ErrSessionNullIDLength,
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
