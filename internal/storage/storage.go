package storage

type Repository interface {
	City() CityRepository
	District() DistrictRepository
	Packaging() PackagingRepository
	Units() UnitsRepository
	PropertyType() PropertyTypeRepository
	ConfectionaryType() ConfectionaryTypeRepository
	Factory() FactoryRepository
	Product() ProductRepository
	Catalogue() CatalogueRepository
	Shop() ShopRepository
	Shipments() ShipmentsRepository
}
