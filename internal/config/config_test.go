package config_test

import (
	"server/internal/apperrors"
	"server/internal/config"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_NewSessionConfig(t *testing.T) {
	tests := []struct {
		name          string
		configObj     *config.SessionConfig
		envVaribles   map[string]string
		expectedError error
	}{
		{
			name: "Valid config",
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "14",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
				"SESSION_ID_LENGTH":        "32",
			},
			configObj: &config.SessionConfig{
				IDLength: 32,
				Duration: time.Duration(14 * 24 * time.Hour),
			},
			expectedError: nil,
		},
		{
			name: "Invalid config (Session duration is 0)",
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "0",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
				"SESSION_ID_LENGTH":        "32",
			},
			configObj:     nil,
			expectedError: apperrors.ErrSessionNullDuration,
		},
		{
			name: "Invalid config (Session ID length is 0)",
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "14",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
				"SESSION_ID_LENGTH":        "0",
			},
			configObj:     nil,
			expectedError: apperrors.ErrSessionNullIDLength,
		},
		{
			name: "Session duration not set",
			envVaribles: map[string]string{
				"SESSION_ID_LENGTH": "32",
			},
			configObj: &config.SessionConfig{
				IDLength: 32,
				Duration: time.Duration(14 * 24 * time.Hour),
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.envVaribles {
				t.Setenv(key, value)
			}
			config, err := config.NewSessionConfig()

			require.Equalf(t, test.configObj, config, test.name)
			require.ErrorIs(t, err, test.expectedError)
		})
	}
}

func Test_GetDBConnectionHost(t *testing.T) {
	tests := []struct {
		name           string
		envVaribles    map[string]string
		expectedResult string
	}{
		{
			name: "Host set",
			envVaribles: map[string]string{
				"POSTGRES_HOST": "coolhost",
			},
			expectedResult: "coolhost",
		},
		{
			name:           "Host not set",
			expectedResult: "localhost",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.envVaribles {
				t.Setenv(key, value)
			}
			host := config.GetDBConnectionHost()

			require.Equalf(t, test.expectedResult, host, test.name)
		})
	}
}

func Test_GetDBPassword(t *testing.T) {
	tests := []struct {
		name           string
		envVaribles    map[string]string
		expectedResult string
		expectedError  error
	}{
		{
			name: "Password set",
			envVaribles: map[string]string{
				"POSTGRES_PASSWORD": "pass",
			},
			expectedResult: "pass",
			expectedError:  nil,
		},
		{
			name:           "Password not set",
			expectedResult: "",
			expectedError:  apperrors.ErrDatabasePWMissing,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.envVaribles {
				t.Setenv(key, value)
			}
			password, err := config.GetDBPassword()

			require.Equalf(t, test.expectedResult, password, test.name)
			require.ErrorIs(t, err, test.expectedError)
		})
	}
}

func Test_GetSessionDurationEnv(t *testing.T) {
	tests := []struct {
		name           string
		envVaribles    map[string]string
		expectedResult time.Duration
		expectedError  error
	}{
		{
			name: "Duration set",
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "15",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
			},
			expectedResult: time.Duration(15 * 24 * time.Hour),
			expectedError:  nil,
		},
		{
			name:           "Duration not set",
			expectedResult: time.Duration(14 * 24 * time.Hour),
			expectedError:  nil,
		},
		{
			name: "Duration less than 1 second",
			envVaribles: map[string]string{
				"SESSION_DURATION_DAYS":    "0",
				"SESSION_DURATION_HOURS":   "0",
				"SESSION_DURATION_MINUTES": "0",
				"SESSION_DURATION_SECONDS": "0",
			},
			expectedResult: 0,
			expectedError:  apperrors.ErrSessionNullDuration,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.envVaribles {
				t.Setenv(key, value)
			}
			password, err := config.GetSessionDurationEnv()

			require.Equalf(t, test.expectedResult, password, test.name)
			require.ErrorIs(t, err, test.expectedError)
		})
	}
}

func Test_GetSessionIDLength(t *testing.T) {
	tests := []struct {
		name           string
		envVaribles    map[string]string
		expectedResult uint
		hasError       bool
	}{
		{
			name: "Session ID length set",
			envVaribles: map[string]string{
				"SESSION_ID_LENGTH": "33",
			},
			expectedResult: 33,
			hasError:       false,
		},
		{
			name:           "Session ID length not set",
			expectedResult: 32,
			hasError:       false,
		},
		{
			name: "Bad string",
			envVaribles: map[string]string{
				"SESSION_ID_LENGTH": "lol",
			},
			expectedResult: 0,
			hasError:       true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.envVaribles {
				t.Setenv(key, value)
			}
			password, err := config.GetSessionIDLength()

			require.Equalf(t, test.expectedResult, password, test.name)
			require.Equalf(t, test.hasError, err != nil, test.name)
		})
	}
}
