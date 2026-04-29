package config

import (
	"os"
	"fmt"
	"encoding/json"
)

// var ConfigData = LoadConfig(ConfigPath)
var ConfigData Config

type Config struct {
	Server ServerConfig `json:"server"`
	Logs LogsConfig `json:"logs"`
}

type LogsConfig struct {
	LogPath string `json:"logPath"`
	LogFile string `json:"logFile"`
	LogLevel string `json:"logLevel"`
}

type ServerConfig struct {
	Port int `json:"port"`
	ReadTimeout int `json:"readTimeout"`
	WriteTimeout int `json:"writeTimeout"`
	IdleTimeout int `json:"idleTimeout"`
}

func LoadConfig(path string) error {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("Failed to unmarshall config: %w", err)
	}

	ConfigData = cfg

	return nil
}