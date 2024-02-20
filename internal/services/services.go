package services

import (
	"log/slog"
	"sladkoezhevo-api/internal/storage"
)

type Services struct {
	storage storage.Storage
	logger  *slog.Logger
}

func NewServices(storage storage.Storage, logger *slog.Logger) *Services {
	return &Services{
		storage: storage,
		logger:  logger,
	}
}
