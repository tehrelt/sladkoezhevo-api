package services

import (
	"log/slog"
	"sladkoezhevo-api/internal/storage"
)

type Services struct {
	logger *slog.Logger

	Cities *CityService
}

func NewServices(repo storage.Repository, logger *slog.Logger) *Services {
	return &Services{
		logger: logger,
		Cities: NewCityService(repo.City()),
	}
}
