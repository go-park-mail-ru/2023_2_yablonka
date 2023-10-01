package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"strconv"
	"time"
)

// TODO salt
func HashFromAuthInfo(info dto.AuthInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(info.Email + info.Password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func GenerateSessionID(n uint) (string, error) {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	buf := make([]rune, n)
	for i := range buf {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			return "", err
		}
		buf[i] = letterRunes[j.Int64()]
	}
	return string(buf), nil
}

func BuildSessionDuration() (time.Duration, error) {
	durationDays, dDok := os.LookupEnv("SESSION_DURATION_DAYS")
	days, _ := strconv.Atoi(durationDays)
	durationHours, dHok := os.LookupEnv("SESSION_DURATION_HOURS")
	hours, _ := strconv.Atoi(durationHours)
	durationMinutes, dMok := os.LookupEnv("SESSION_DURATION_MINUTES")
	minutes, _ := strconv.Atoi(durationMinutes)
	durationSeconds, dSok := os.LookupEnv("SESSION_DURATION_SECONDS")
	seconds, _ := strconv.Atoi(durationSeconds)
	if !(dDok || dHok || dMok || dSok) {
		return 0, apperrors.ErrSessionDurationMissing
	}
	totalDuration := time.Duration(24*days+hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	if totalDuration < time.Second {
		return 0, apperrors.ErrSessionNullDuration
	}
	return totalDuration, nil
}
