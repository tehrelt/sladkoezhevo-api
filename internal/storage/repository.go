package storage

import "sladkoezhevo-api/internal/models"

type CityRepository interface {
	Create(*models.City) error
	Get() ([]*models.City, error)
	GetOne(int) (*models.City, error)
	Update(*models.City) error
	Delete(id int) error
}

type PackagingRepository interface {
	Create(*models.Packaging) error
	Get() ([]*models.Packaging, error)
	GetOne(int) (*models.Packaging, error)
	Update(*models.Packaging) error
	Delete(id int) error
}
