package services

import (
	"log/slog"
	"sladkoezhevo-api/internal/storage"
)

type Services struct {
	logger *slog.Logger

	Cities *cityService
}

func NewServices(repo storage.Repository, logger *slog.Logger) *Services {
	return &Services{
		logger: logger,
		Cities: newCityService(repo.City()),
	}
}
