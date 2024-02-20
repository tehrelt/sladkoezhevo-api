package services

import (
	"sladkoezhevo-api/internal/models"
)

type CityRepository interface {
	Create(*models.City) error
	Get() ([]*models.City, error)
	GetOne(int) (*models.City, error)
	Update(*models.City) error
	Delete(id int) error
}

type CityService struct {
	r CityRepository
}

func NewCityService(repo CityRepository) *CityService {
	return &CityService{
		r: repo,
	}
}

func (s *CityService) Get() ([]*models.City, error) {
	return s.r.Get()
}
