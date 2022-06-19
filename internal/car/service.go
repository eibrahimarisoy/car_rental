package car

import (
	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/office"
	"github.com/eibrahimarisoy/car_rental/pkg/filters"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type CarService struct {
	carRepo      CarRepositoryInterface
	officeRepo   office.OfficeRepositoryInterface
	locationRepo location.LocationRepositoryInterface
}

//go:generate mockgen -destination=../../mocks/car/car_service_interface.go -package=car github.com/eibrahimarisoy/car_rental/internal/car CarServiceInterface
type CarServiceInterface interface {
	GetCars(pg *pgHelper.Pagination, filter *filters.CarFilter) (*[]models.Car, error)
	CreateCar(car *models.Car) (*models.Car, error)
}

// NewCarService creates a new car service
func NewCarService(
	carRepo CarRepositoryInterface,
	locationRepo location.LocationRepositoryInterface,
	officeRepo office.OfficeRepositoryInterface,
) *CarService {
	return &CarService{
		carRepo:      carRepo,
		officeRepo:   officeRepo,
		locationRepo: locationRepo,
	}
}

// GetCars returns cars by filter and pagination
func (s *CarService) GetCars(pg *pgHelper.Pagination, filter *filters.CarFilter) (*[]models.Car, error) {
	location, err := s.locationRepo.GetLocationByID(filter.Location)
	if err != nil {
		return nil, err
	}
	pickupWeekDay := int(filter.PickupDate.ToTime().Weekday())
	dropoffWeekDay := int(filter.DropoffDate.ToTime().Weekday())

	pickupTime := filter.PickupTime.ToTime()
	dropoffTime := filter.DropoffTime.ToTime()

	officeIDs, err := s.officeRepo.GetOfficeIDs(location.ID, pickupWeekDay, dropoffWeekDay, pickupTime, dropoffTime)
	if err != nil {
		return nil, err
	}

	return s.carRepo.GetCarsByOfficeIDs(pg, officeIDs)
}

// CreateCar creates a car and returns it
func (s *CarService) CreateCar(car *models.Car) (*models.Car, error) {
	return s.carRepo.CreateCar(car)
}
