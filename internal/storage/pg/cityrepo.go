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

func (r *CityRepository) Get() ([]*models.City, error) {

	var cities []*models.City

	rows, err := r.db.Query("SELECT * FROM city;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var city models.City
		err := rows.Scan(&city.Id, &city.Name)
		if err != nil {
			return nil, err
		}

		cities = append(cities, &city)
	}

	return cities, nil
}
