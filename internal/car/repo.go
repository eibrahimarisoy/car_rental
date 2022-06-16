package car

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"

	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
)

type CarRepositoryInterface interface {
	GetCars(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateCar(car *models.Car) (*models.Car, error)
}

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (r *CarRepository) Migration() {
	r.db.AutoMigrate(&models.Car{})
}

// TODO
func (r *CarRepository) GetCars(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	var cars []*models.Car
	var totalRows int64

	query := r.db.Model(&models.Car{}).Preload("Vendor").Preload("Office").Scopes().Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&cars)

	if query.Error != nil {
		return nil, query.Error
	}

	pg.Rows = &cars
	return pg, nil
}

// TODO

func (r *CarRepository) CreateCar(car *models.Car) (*models.Car, error) {

	if err := r.db.Create(car).Error; err != nil {
		return nil, err
	}
	return car, nil
}
