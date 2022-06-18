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

type CarServiceInterface interface {
	GetCars(pg *pgHelper.Pagination, filter *filters.CarFilter) (*pgHelper.Pagination, error)
	CreateCar(car *models.Car) (*models.Car, error)
}

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

func (s *CarService) GetCars(pg *pgHelper.Pagination, filter *filters.CarFilter) (*pgHelper.Pagination, error) {
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

func (s *CarService) CreateCar(car *models.Car) (*models.Car, error) {
	return s.carRepo.CreateCar(car)
}
