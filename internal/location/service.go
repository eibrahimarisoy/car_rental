package location

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

//go:generate mockgen -destination=../../mocks/location/location_service_interface.go -package=location github.com/eibrahimarisoy/car_rental/internal/location LocationServiceInterface
type LocationService struct {
	locationRepo LocationRepositoryInterface
}

type LocationServiceInterface interface {
	GetAllActiveLocations(pg *pgHelper.Pagination) (*[]models.Location, error)
	CreateLocation(location *models.Location) (*models.Location, error)
}

// NewLocationService creates a new location service
func NewLocationService(locationRepo LocationRepositoryInterface) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

// GetAllActiveLocations returns all active locations
func (s *LocationService) GetAllActiveLocations(pg *pgHelper.Pagination) (*[]models.Location, error) {
	return s.locationRepo.GetAllActiveLocations(pg)
}

// CreateLocation creates a location and returns it
func (s *LocationService) CreateLocation(location *models.Location) (*models.Location, error) {
	return s.locationRepo.CreateLocation(location)
}
