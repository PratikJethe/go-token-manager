package server

import (
    "os"
)

type ServerConfig struct {
    Host string
    Port string
}

// NewServerConfig initializes a new ServerConfig by reading environment variables.
func NewServerConfig() *ServerConfig {
    host := os.Getenv("SERVER_HOST")
    if host == "" {
        host = "0.0.0.0" // Default host if not specified
    }

    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }

    return &ServerConfig{
        Host: host,
        Port: port,
    }
}
