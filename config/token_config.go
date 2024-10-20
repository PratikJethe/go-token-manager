package config

import (
	"os"
	"strconv"
)

type TokenConfig struct {
	TokenActiveDuration int // in seconds
	TokenExpireDuration int // in seconds
	TokenLength         int // length of the token
}

// GetNew initializes a new TokenConfig with values from environment variables
func NewTokenConfig() *TokenConfig {
	tokenActiveDurationStr := os.Getenv("TOKEN_ACTIVE_DURATION")
	tokenExpireDurationStr := os.Getenv("TOKEN_EXPIRATION_DURATION")
	tokenLengthStr := os.Getenv("TOKEN_LENGTH")

	// Default values
	defaultActiveDuration := 60  // default to 60 seconds
	defaultExpireDuration := 300 // default to 300 seconds
	defaultTokenLength := 16     // default to 16 characters

	tokenActiveDuration, err := strconv.Atoi(tokenActiveDurationStr)
	if err != nil || tokenActiveDuration <= 0 {
		tokenActiveDuration = defaultActiveDuration
	}

	tokenExpireDuration, err := strconv.Atoi(tokenExpireDurationStr)
	if err != nil || tokenExpireDuration <= 0 {
		tokenExpireDuration = defaultExpireDuration
	}

	tokenLength, err := strconv.Atoi(tokenLengthStr)
	if err != nil || tokenLength <= 0 {
		tokenLength = defaultTokenLength
	}

	return &TokenConfig{
		TokenActiveDuration: tokenActiveDuration,
		TokenExpireDuration: tokenExpireDuration,
		TokenLength:         tokenLength,
	}
}
