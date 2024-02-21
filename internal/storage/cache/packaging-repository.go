package cache

import "sladkoezhevo-api/internal/models"

type packagingRepository struct {
	cache []*models.Packaging
}

func (p packagingRepository) Create(packaging *models.Packaging) error {
	p.cache = append(p.cache, packaging)
	return nil
}

func (p packagingRepository) Get() ([]*models.Packaging, error) {
	//TODO implement me
	panic("implement me")
}

func (p packagingRepository) GetOne(i int) (*models.Packaging, error) {
	//TODO implement me
	panic("implement me")
}

func (p packagingRepository) Update(packaging *models.Packaging) error {
	//TODO implement me
	panic("implement me")
}

func (p packagingRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func newPackagingRepo(cache []*models.Packaging) *packagingRepository {
	return &packagingRepository{
		cache: cache,
	}
}
