package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/storage"
)

type unitsRepository struct {
	*Storage
}

func (r *unitsRepository) Create(item *models.Units) error {
	err := r.db.QueryRow(
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", unitsTable),
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

func (r *unitsRepository) Get() ([]*models.Units, error) {

	var items []*models.Units

	query, err := r.db.Query(fmt.Sprintf("SELECT * FROM %s;", unitsTable))
	if err != nil {
		return nil, err
	}

	for query.Next() {
		var record models.Units
		err := query.Scan(&record.Id, &record.Name)
		if err != nil {
			return nil, err
		}

		items = append(items, &record)
	}

	return items, nil
}

func (r *unitsRepository) GetOne(id int) (*models.Units, error) {

	row := r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", unitsTable), id)

	var record models.Units
	err := row.Scan(&record.Id, &record.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return &record, nil
}

func (r *unitsRepository) Update(record *models.Units) error {
	_, err := r.db.Query(fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2;", unitsTable), record.Name, record.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *unitsRepository) Delete(id int) error {
	_, err := r.db.Query(fmt.Sprintf("REMOVE FROM %s WHERE id=$1;", unitsTable), id)
	return err
}
