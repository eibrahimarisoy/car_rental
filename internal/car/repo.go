package car

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
)

type CarRepositoryInterface interface {
	CreateCar(car *models.Car) (*models.Car, error)
	GetCarByID(id uuid.UUID) (*models.Car, error)
	UpdateCarStatus(car *models.Car) (*models.Car, error)
	GetCarsByOfficeIDs(pg *pgHelper.Pagination, officeIDs []uuid.UUID) (*[]models.Car, error)
}

type CarRepository struct {
	db *gorm.DB
}

// NewCarRepository creates a new car repository
func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

// Migrations for car table
func (r *CarRepository) Migration() {
	r.db.AutoMigrate(&models.Car{})
}

// CreateCar creates a car and returns it
func (r *CarRepository) CreateCar(car *models.Car) (*models.Car, error) {

	if err := r.db.Create(car).Error; err != nil {
		return nil, err
	}
	return car, nil
}

// GetCarByID returns a car by id
func (r *CarRepository) GetCarByID(id uuid.UUID) (*models.Car, error) {
	car := models.Car{}
	res := r.db.Model(&models.Car{}).Where("status = ? AND id = ?", models.CarStatusAvailable, id).First(&car)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, errorHandler.CarNotFoundError
	} else if res.Error != nil {
		return nil, res.Error
	}
	return &car, nil
}

// UpdateCarStatus updates a car status
func (r *CarRepository) UpdateCarStatus(car *models.Car) (*models.Car, error) {
	res := r.db.Model(&models.Car{}).Where("id = ?", car.ID).Update("status", car.Status)
	if res.Error != nil {
		return nil, res.Error
	}
	return car, nil
}

// GetCarsByOfficeIDs returns a list of cars by office ids
func (r *CarRepository) GetCarsByOfficeIDs(pg *pgHelper.Pagination, officeIDs []uuid.UUID) (*[]models.Car, error) {
	var cars []models.Car
	var totalRows int64

	query := r.db.Model(&models.Car{}).Preload("Office").Preload("Vendor").
		Where(
			"status = ? AND office_id IN ?", models.CarStatusAvailable, officeIDs,
		).Find(&cars).Count(&totalRows)

	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&cars)

	if query.Error != nil {
		return nil, query.Error
	}

	return &cars, nil
}
