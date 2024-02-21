package cache

import (
	"fmt"
	"sladkoezhevo-api/internal/models"
)

type cityRepository struct {
	cache []*models.City
}

func (c cityRepository) Create(city *models.City) error {
	c.cache = append(c.cache, city)
	return nil
}

func (c cityRepository) Get() ([]*models.City, error) {
	return c.cache, nil
}

func (c cityRepository) GetOne(i int) (*models.City, error) {
	if i > len(c.cache) {
		return nil, fmt.Errorf("out of bounds array")
	}

	if i < 0 {
		return nil, fmt.Errorf("negative id not allowed")
	}

	return c.cache[i], nil
}

func (c cityRepository) Update(city *models.City) error {
	//TODO implement me
	panic("implement me")
}

func (c cityRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func newCityRepo(cache []*models.City) *cityRepository {
	return &cityRepository{
		cache: cache,
	}
}
