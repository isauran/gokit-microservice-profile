package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	serverPortEnvName = "APP.SERVER.PORT"
)

type ServerConfig interface {
	ServerPort() string
}

type serverConfig struct {
	port string
}

func NewServerConfig() (ServerConfig, error) {
	port := os.Getenv(serverPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("server port not found")
	}

	return &serverConfig{
		port: port,
	}, nil
}

func (cfg *serverConfig) ServerPort() string {
	return fmt.Sprintf(":%s", cfg.port)
}
