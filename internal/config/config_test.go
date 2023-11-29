package config_test

import (
	"os"
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
		configObj     config.SessionConfig
		envVaribles   map[string]string
		expectedError error
	}{
		{
			name: "Valid config",
			configObj: config.SessionConfig{
				Duration: time.Duration(30 * 24 * time.Hour),
				IDLength: 32,
			},
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "14",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
				"SESSION_ID_LENGTH":        "32",
			},
			expectedError: nil,
		},
		{
			name: "Invalid config (Session duration is 0)",
			configObj: config.SessionConfig{
				Duration: time.Duration(0 * time.Second),
				IDLength: 32,
			},
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "0",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
				"SESSION_ID_LENGTH":        "32",
			},
			expectedError: apperrors.ErrSessionNullDuration,
		},
		{
			name: "Invalid config (Session ID length is 0)",
			configObj: config.SessionConfig{
				Duration: time.Duration(30 * 24 * time.Hour),
				IDLength: 0,
			},
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "14",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
				"SESSION_ID_LENGTH":        "0",
			},
			expectedError: apperrors.ErrSessionNullIDLength,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			for key, value := range test.envVaribles {
				os.Setenv(key, value)
			}
			_, err := config.NewSessionConfig()

			require.ErrorIs(t, err, test.expectedError)
		})
	}
}
