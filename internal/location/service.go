package location

import (
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type LocationService struct {
	locationRepo LocationRepositoryInterface
}

type LocationServiceInterface interface {
	GetAllActiveLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
}

func NewLocationService(locationRepo LocationRepositoryInterface) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

func (s *LocationService) GetAllActiveLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	return s.locationRepo.GetAllActiveLocations(pg)
}
