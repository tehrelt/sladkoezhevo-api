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
	repository CityRepository
}

func NewCityService(repo CityRepository) *CityService {
	return &CityService{
		repository: repo,
	}
}

func (s *CityService) Get() ([]*models.City, error) {
	return s.repository.Get()
}

func (s *CityService) GetOne(id int) (*models.City, error) {
	return s.repository.GetOne(id)
}
