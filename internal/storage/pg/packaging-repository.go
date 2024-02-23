package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"sladkoezhevo-api/internal/models"
	"sladkoezhevo-api/internal/storage"
)

type packagingRepository struct {
	*Storage
}

func (r *packagingRepository) Create(item *models.Packaging) error {
	err := r.db.QueryRow(
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", packagingTable),
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

func (r *packagingRepository) Get() ([]*models.Packaging, error) {

	var items []*models.Packaging

	query, err := r.db.Query(fmt.Sprintf("SELECT * FROM %s;", packagingTable))
	if err != nil {
		return nil, err
	}

	for query.Next() {
		var record models.Packaging
		err := query.Scan(&record.Id, &record.Name)
		if err != nil {
			return nil, err
		}

		items = append(items, &record)
	}

	return items, nil
}

func (r *packagingRepository) GetOne(id int) (*models.Packaging, error) {

	row := r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", packagingTable), id)

	var record models.Packaging
	err := row.Scan(&record.Id, &record.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return &record, nil
}

func (r *packagingRepository) Update(record *models.Packaging) error {
	_, err := r.db.Query(fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2;", packagingTable), record.Name, record.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *packagingRepository) Delete(id int) error {
	_, err := r.db.Query(fmt.Sprintf("REMOVE FROM %s WHERE id=$1;", packagingTable), id)
	return err
}
