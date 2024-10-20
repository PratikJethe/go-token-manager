package utils_test

import (
	"encoding/hex"
	"testing"

	"github.com/pratikjethe/go-token-manager/utils"
)

func TestGenerateRandomToken(t *testing.T) {
	size := 16
	token, err := utils.GenerateRandomToken(size)

	// Check for errors
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that the length of the generated token matches the expected size.
	// Since the token is hex-encoded, its length should be twice the size of the input.
	expectedLength := size * 2
	if len(token) != expectedLength {
		t.Errorf("Expected token length %d, got %d", expectedLength, len(token))
	}

	// Check if the token is a valid hex string.
	_, decodeErr := hex.DecodeString(token)
	if decodeErr != nil {
		t.Errorf("Generated token is not a valid hex string: %v", decodeErr)
	}
}
