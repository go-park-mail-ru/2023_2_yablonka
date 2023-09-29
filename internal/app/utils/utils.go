package utils

import (
	"fmt"
	"server/internal/pkg/datatypes"

	"crypto/sha256"
)

func HashFromAuthInfo(info datatypes.AuthInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(info.Email + info.Password))
	fmt.Println(hasher.Sum(nil))
	fmt.Printf("Hash as %%x %x\n", hasher.Sum(nil))
	fmt.Printf("Hash as %%s %s\n", string(hasher.Sum(nil)))
	return string(hasher.Sum(nil))
}
