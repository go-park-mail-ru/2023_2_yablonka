package utils

import (
	"fmt"
	"server/internal/pkg/datatypes"

	"crypto/sha256"
)

func HashFromAuthInfo(info datatypes.AuthInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(info.Email + info.Password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
