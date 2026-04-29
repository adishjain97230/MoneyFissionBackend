package logging

import (
	"log/slog"
	"MoneyFissionBackend/config"
	"os"
	"fmt"
	"path/filepath"
)

var Logger *slog.Logger

func New() (*slog.Logger, func() error, error) {
	if err := os.MkdirAll(config.ConfigData.Logs.LogPath, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	f, err := os.OpenFile(filepath.Join(config.ConfigData.Logs.LogPath, config.ConfigData.Logs.LogFile), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file: %w", err)
	}

	handler := slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.Level(toSlogLevel(config.ConfigData.Logs.LogLevel)),
	})

	logger := slog.New(handler)
	return logger, func() error { return f.Close() }, nil
}

func SetupLogger() (func() error, error){
	logger, closeLog, err := New()
	if err != nil {
		return nil, err
	}
	Logger = logger
	return closeLog, nil
}