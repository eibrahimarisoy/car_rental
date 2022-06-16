package location

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type LocationService struct {
	locationRepo LocationRepositoryInterface
}

type LocationServiceInterface interface {
	GetAllActiveLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateLocation(location *models.Location) (*models.Location, error)
}

func NewLocationService(locationRepo LocationRepositoryInterface) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

func (s *LocationService) GetAllActiveLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	return s.locationRepo.GetAllActiveLocations(pg)
}

func (s *LocationService) CreateLocation(location *models.Location) (*models.Location, error) {
	return s.locationRepo.CreateLocation(location)
}
