package location

import "github.com/eibrahimarisoy/car_rental/internal/models"

type LocationService struct {
	locationRepo LocationRepositoryInterface
}

type LocationServiceInterface interface {
	GetAllLocations() ([]*models.Location, error)
}

func NewLocationService(locationRepo LocationRepositoryInterface) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

func (s *LocationService) GetAllLocations() ([]*models.Location, error) {
	return s.locationRepo.GetAllLocations()
}
