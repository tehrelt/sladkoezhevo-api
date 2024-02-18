package main

import (
	"log/slog"
	"os"
	"sladkoezhevo-api/internal/config"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	config := config.NewConfig("config/.yaml")
	logger := setupLogger(config.Env)
	logger.Debug("Logger initialized")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
