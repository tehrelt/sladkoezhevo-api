package pg

import (
	"fmt"
	"sladkoezhevo-api/internal/config"
	"sladkoezhevo-api/internal/storage"

	"github.com/go-sqlx/sqlx"

	_ "github.com/lib/pq"
)

const (
	cityTable              = "city"
	packagingTable         = "packaging"
	unitsTable             = "units"
	propertyTypeTable      = "property_type"
	confectionaryTypeTable = "confectionary_type"
	factoryTable           = "factory"
	productTable           = "product"
	catalogueTable         = "catalogue"
	districtTable          = "district"
	shipmentsTable         = "shipments"
)

type Storage struct {
	db            *sqlx.DB
	city          *cityRepository
	packaging     *packagingRepository
	units         *unitsRepository
	property      *propertyRepository
	confectionary *confectionaryRepository
}

func (s *Storage) District() storage.DistrictRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Packaging() storage.PackagingRepository {
	if s.packaging != nil {
		return s.packaging
	}
	s.packaging = &packagingRepository{s}
	return s.packaging
}

func (s *Storage) Units() storage.UnitsRepository {
	if s.units != nil {
		return s.units
	}
	s.units = &unitsRepository{s}
	return s.units
}

func (s *Storage) PropertyType() storage.PropertyTypeRepository {
	if s.property != nil {
		return s.property
	}
	s.property = &propertyRepository{s}
	return s.property
}

func (s *Storage) ConfectionaryType() storage.ConfectionaryTypeRepository {
	if s.confectionary != nil {
		return s.confectionary
	}
	s.confectionary = &confectionaryRepository{s}
	return s.confectionary
}

func (s *Storage) Factory() storage.FactoryRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Product() storage.ProductRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Catalogue() storage.CatalogueRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Shop() storage.ShopRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Shipments() storage.ShipmentsRepository {
	//TODO implement me
	panic("implement me")
}

func New(config *config.DatabaseConfig) *Storage {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))
	if err != nil {
		panic(err)
	}

	return &Storage{db: db}
}

func (s *Storage) City() storage.CityRepository {
	if s.city != nil {
		return s.city
	}
	s.city = &cityRepository{s}
	return s.city
}
