package car

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type CarService struct {
	carRepo CarRepositoryInterface
}

type CarServiceInterface interface {
	GetCars(pg *pgHelper.Pagination, filter *CarFilter) (*pgHelper.Pagination, error)
	CreateCar(car *models.Car) (*models.Car, error)
}

func NewCarService(carRepo CarRepositoryInterface) *CarService {
	return &CarService{
		carRepo: carRepo,
	}
}

func (s *CarService) GetCars(pg *pgHelper.Pagination, filter *CarFilter) (*pgHelper.Pagination, error) {
	return s.carRepo.GetCars(pg, filter)
}

func (s *CarService) CreateCar(car *models.Car) (*models.Car, error) {
	return s.carRepo.CreateCar(car)
}
