package main

import (
	"log"
	"log/slog"
	"os"
	"sladkoezhevo-api/internal/config"
	"sladkoezhevo-api/internal/handlers"
	"sladkoezhevo-api/internal/services"
	"sladkoezhevo-api/internal/storage/pg"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

// @title SladkoEzhevoAPI
// @version 0.1
// @description Swagger doc for SladkoEzhevoAPI
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email haxrelt@gmail.com
// @license.name MIT
// @host localhost:7000
// @BasePath /api/v1
func main() {
	config := config.NewConfig("config/.yaml")
	logger := setupLogger(config.Env)
	logger.Debug("Logger initialized")

	repository := pg.New(config.DatabaseConfig)
	services := services.NewServices(repository, logger)
	server := handlers.New(services, logger)

	log.Fatal(server.Start(config.Port))
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
