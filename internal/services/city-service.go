package services

import (
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/storage"
)

type cityService struct {
	repository storage.CityRepository
}

func newCityService(repo storage.CityRepository) *cityService {
	return &cityService{
		repository: repo,
	}
}

func (s *cityService) Create(city *models.City) error {
	return s.repository.Create(city)
}
func (s *cityService) Get() ([]*models.City, error) {
	return s.repository.Get()
}
func (s *cityService) GetOne(id int) (*models.City, error) {
	return s.repository.GetOne(id)
}
func (s *cityService) Update(city *models.City) error {
	return s.repository.Update(city)
}
func (s *cityService) Delete(id int) error {
	return s.repository.Delete(id)
}
