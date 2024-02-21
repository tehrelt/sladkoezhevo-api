package cache

import (
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/storage"
)

type Cache struct {
	Cities    []*models.City
	Packaging []*models.Packaging
}

func newCache() *Cache {
	return &Cache{}
}

type Storage struct {
	cache               *Cache
	cityRepository      *cityRepository
	packagingRepository *packagingRepository
}

func (s Storage) City() storage.CityRepository {
	if s.cityRepository != nil {
		return s.cityRepository
	}

	s.cityRepository = newCityRepo(s.cache.Cities)
	return s.cityRepository
}

func (s Storage) Packaging() storage.PackagingRepository {
	if s.packagingRepository != nil {
		return s.packagingRepository
	}

	s.packagingRepository = newPackagingRepo(s.cache.Packaging)
	return s.packagingRepository
}

func New() *Storage {
	cache := newCache()
	return &Storage{
		cache: cache,
	}
}
