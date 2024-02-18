package pg

import (
	"sladkoezhevo-api/internal/storage/models"
)

type CityRepository struct {
	*Storage
}

func (r *CityRepository) Create(city *models.City) error {
	return r.db.QueryRow(
		"INSERT INTO city (name) VALUES ($1) RETURNING id",
		city.Name,
	).Scan(&city.Id)
}
