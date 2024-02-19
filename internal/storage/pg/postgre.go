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
	city *CityRepository
}

func New(config *config.DatabaseConfig) (*Storage, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) City() storage.CityRepository {
	if s.city != nil {
		return s.city
	}
	s.city = &CityRepository{s}
	return s.city
}

