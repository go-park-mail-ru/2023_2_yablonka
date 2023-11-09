package config_test

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
		config        config.SessionConfig
		serviceType   string
		expectedError error
	}{
		{
			name: "Valid config",
			config: config.SessionConfig{
				Duration: time.Duration(30 * 24 * time.Hour),
				IDLength: 32,
			},
			expectedError: nil,
		},
		{
			name: "Invalid config (Session duration is 0)",
			config: config.SessionConfig{
				Duration: time.Duration(0 * time.Second),
				IDLength: 32,
			},
			expectedError: apperrors.ErrSessionNullDuration,
		},
		{
			name: "Invalid config (Session ID length is 0)",
			config: config.SessionConfig{
				Duration: time.Duration(30 * 24 * time.Hour),
				IDLength: 0,
			},
			expectedError: apperrors.ErrSessionNullIDLength,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := config.NewSessionConfig()

			require.ErrorIs(t, test.expectedError, err)
		})
	}
}
