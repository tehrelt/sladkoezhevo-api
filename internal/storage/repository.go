package storage

import "sladkoezhevo-api/internal/models"

type CityRepository interface {
	Create(*models.City) error
	Get() ([]*models.City, error)
	GetOne(int) (*models.City, error)
	Update(*models.City) error
	Delete(id int) error
}

type DistrictRepository interface {
	Create(*models.District) error
	Get() ([]*models.District, error)
	GetOne(int) (*models.District, error)
	Update(*models.District) error
	Delete(id int) error
}

type PackagingRepository interface {
	Create(*models.Packaging) error
	Get() ([]*models.Packaging, error)
	GetOne(int) (*models.Packaging, error)
	Update(*models.Packaging) error
	Delete(id int) error
}

type UnitsRepository interface {
	Create(*models.Units) error
	Get() ([]*models.Units, error)
	GetOne(int) (*models.Units, error)
	Update(*models.Units) error
	Delete(id int) error
}

type PropertyTypeRepository interface {
	Create(*models.PropertyType) error
	Get() ([]*models.PropertyType, error)
	GetOne(int) (*models.PropertyType, error)
	Update(*models.PropertyType) error
	Delete(id int) error
}

type ConfectionaryTypeRepository interface {
	Create(*models.ConfectionaryType) error
	Get() ([]*models.ConfectionaryType, error)
	GetOne(int) (*models.ConfectionaryType, error)
	Update(*models.ConfectionaryType) error
	Delete(id int) error
}

type FactoryRepository interface {
	Create(*models.Factory) error
	Get() ([]*models.Factory, error)
	GetOne(int) (*models.Factory, error)
	Update(*models.Factory) error
	Delete(id int) error
}

type ProductRepository interface {
	Create(*models.Product) error
	Get() ([]*models.Product, error)
	GetOne(int) (*models.Product, error)
	Update(*models.Product) error
	Delete(id int) error
}

type CatalogueRepository interface {
	Create(*models.ProductEntry) error
	Get() ([]*models.ProductEntry, error)
	GetOne(int) (*models.ProductEntry, error)
	Update(*models.ProductEntry) error
	Delete(id int) error
}
