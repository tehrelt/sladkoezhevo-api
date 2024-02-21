package pg

import (
	"fmt"
	"sladkoezhevo-api/internal/config"
	"sladkoezhevo-api/internal/storage"

	"github.com/go-sqlx/sqlx"

	_ "github.com/lib/pq"
)

const (
	CityTable = "city"
)

type Storage struct {
	db   *sqlx.DB
	city *cityRepository
}

func (s *Storage) District() storage.DistrictRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Packaging() storage.PackagingRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Units() storage.UnitsRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) PropertyType() storage.PropertyTypeRepository {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) ConfectionaryType() storage.ConfectionaryTypeRepository {
	//TODO implement me
	panic("implement me")
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
