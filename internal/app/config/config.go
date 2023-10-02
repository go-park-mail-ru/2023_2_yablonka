package config

import (
	"log"
	"net/http"
	"os"
	"server/internal/app/handlers"
	"server/internal/app/utils"
	"server/internal/apperrors"
	"server/internal/pkg/entities"
	authservice "server/internal/service/auth"
	boardservice "server/internal/service/board"
	userservice "server/internal/service/user"
	"server/internal/storage/in_memory"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func NewEnvConfig(filepath string) (*entities.ServerConfig, error) {
	var err error
	if filepath == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(filepath)
	}
	if err != nil {
		return nil, apperrors.ErrEnvNotFound
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

	sessionDuration, err := utils.BuildSessionDuration()
	if err != nil {
		return nil, err
	}

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok || jwtSecret == "" {
		log.Fatal("JWT secret is missing from config or empty")
		return nil, apperrors.ErrJWTSecretMissing
	}

	return &entities.ServerConfig{
		SessionIDLength: sidLength,
		SessionDuration: sessionDuration,
		JWTSecret:       jwtSecret,
	}, nil
}

func ValidateConfig(config *entities.ServerConfig) (bool, error) {
	if config.SessionDuration < time.Duration(1*time.Second) {
		return false, apperrors.ErrSessionNullDuration
	}
	if config.SessionIDLength < 1 {
		return false, apperrors.ErrSessionNullIDLength
	}
	if config.JWTSecret == "" {
		return false, apperrors.ErrJWTSecretMissing
	}
	return true, nil
}

func ConfigMux(config *entities.ServerConfig, mux *http.ServeMux) error {
	ok, err := ValidateConfig(config)
	if !ok {
		return err
	}

	userStorage := in_memory.NewUserStorage()
	authStorage := in_memory.NewAuthStorage()
	boardStorage := in_memory.NewBoardStorage()

	authService := authservice.NewAuthSessionService(config, authStorage)
	userAuthService := userservice.NewAuthUserService(userStorage)
	boardService := boardservice.NewBoardService(boardStorage)
	// userService := userservice.NewUserService(userStorage)
	authHandler := handlers.NewAuthHandler(authService, userAuthService)
	boardHandler := handlers.NewBoardHandler(authService, boardService)

	mux.HandleFunc("/api/v1/login/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodPost:
			authHandler.LogIn(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/signup/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodPost:
			authHandler.SignUp(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/getuserboards/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log.Println(r.URL.Path)

		switch r.Method {
		case http.MethodGet:
			boardHandler.GetUserBoards(w, r)
		default:
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})
	return nil
}
