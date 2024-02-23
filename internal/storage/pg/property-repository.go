package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/storage"
)

type propertyRepository struct {
	*Storage
}

func (r *propertyRepository) Create(item *models.PropertyType) error {
	err := r.db.QueryRow(
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", propertyTypeTable),
		item.Name,
	).Scan(&item.Id)

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

func (r *propertyRepository) Get() ([]*models.PropertyType, error) {

	var items []*models.PropertyType

	query, err := r.db.Query(fmt.Sprintf("SELECT * FROM %s;", propertyTypeTable))
	if err != nil {
		return nil, err
	}

	for query.Next() {
		var record models.PropertyType
		err := query.Scan(&record.Id, &record.Name)
		if err != nil {
			return nil, err
		}

		items = append(items, &record)
	}

	return items, nil
}

func (r *propertyRepository) GetOne(id int) (*models.PropertyType, error) {

	row := r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", propertyTypeTable), id)

	var record models.PropertyType
	err := row.Scan(&record.Id, &record.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return &record, nil
}

func (r *propertyRepository) Update(record *models.PropertyType) error {
	_, err := r.db.Query(fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2;", propertyTypeTable), record.Name, record.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *propertyRepository) Delete(id int) error {
	_, err := r.db.Query(fmt.Sprintf("REMOVE FROM %s WHERE id=$1;", propertyTypeTable), id)
	return err
}
