package storage

type Repository interface {
	City() CityRepository
	Packaging() PackagingRepository
}
