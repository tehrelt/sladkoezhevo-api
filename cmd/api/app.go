package main

import (
	"log/slog"
	"os"
	"sladkoezhevo-api/internal/config"
	"sladkoezhevo-api/internal/storage/pg"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	config := config.NewConfig("config/.yaml")
	logger := setupLogger(config.Env)
	logger.Debug("Logger initialized")

	repo, err := pg.New(config.DatabaseConfig)
	if err != nil {
		logger.Error("CANNOT INIT REPOSITORY", slog.String("err", err.Error()))
		return
	}

	_ = repo
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	log = log.With(slog.String("env", env))

	return log
}
