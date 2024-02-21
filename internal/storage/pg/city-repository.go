package pg

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/storage"
)

type cityRepository struct {
	*Storage
}

func (r *cityRepository) Create(city *models.City) error {
	err := r.db.QueryRow(
		"INSERT INTO city (name) VALUES ($1) RETURNING id",
		city.Name,
	).Scan(&city.Id)

	if err != nil {
		var pqError *pq.Error
		ok := errors.As(err, &pqError)
		if ok {
			if pqError.Code == "23505" {
				return storage.ErrRecordAlreadyExists
			}
		}

		return err
	}

	return nil
}

func (r *cityRepository) Get() ([]*models.City, error) {

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

func (r *cityRepository) GetOne(id int) (*models.City, error) {

	row := r.db.QueryRow("SELECT * FROM city WHERE id=$1;", id)

	var city models.City
	err := row.Scan(&city.Id, &city.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return &city, nil
}

func (r *cityRepository) Update(city *models.City) error {
	_, err := r.db.Query("UPDATE city SET name=$1 WHERE id=$2;", city.Name, city.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *cityRepository) Delete(id int) error {
	_, err := r.db.Query("REMOVE FROM city WHERE id=$1;", id)
	return err
}
