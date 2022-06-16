package location

import (
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type LocationService struct {
	locationRepo LocationRepositoryInterface
}

type LocationServiceInterface interface {
	GetAllLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
}

func NewLocationService(locationRepo LocationRepositoryInterface) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

func (s *LocationService) GetAllLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	return s.locationRepo.GetAllLocations(pg)
}
