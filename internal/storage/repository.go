package storage

import "sladkoezhevo-api/internal/models"

type CityRepository interface {
	Create(*models.City) error
	Get() ([]*models.City, error)
	GetOne(int) (*models.City, error)
	Update(*models.City) error
	Delete(id int) error
}
