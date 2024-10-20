package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomToken(size int) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err // Handle error as needed.
	}
	return hex.EncodeToString(bytes), err
}
