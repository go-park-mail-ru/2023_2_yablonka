package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(email string, password string) (string, error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(email+password), 8)
	if err != nil {
		return "", err
	}
	return string(bhash), nil
}
