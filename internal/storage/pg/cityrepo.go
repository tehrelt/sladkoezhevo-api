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

func (r *CityRepository) GetOne(id int) (*models.City, error) {

	row := r.db.QueryRow("SELECT * FROM city WHERE id=$1;", id)

	var city models.City

	err := row.Scan(&city.Id, &city.Name)
	if err != nil {
		return nil, err
	}

	return &city, nil
}

func (r *CityRepository) Update(city *models.City) error {
	_, err := r.db.Query("UPDATE city SET name=$1 WHERE id=$2;", city.Name, city.Id)
	return err
}

func (r *CityRepository) Delete(id int) error {
	_, err := r.db.Query("REMOVE FROM city WHERE id=$1;", id)
	return err
}
