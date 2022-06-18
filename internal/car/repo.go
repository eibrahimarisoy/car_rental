package car

import (
	"fmt"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"

	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
)

type CarRepositoryInterface interface {
	GetCars(pg *pgHelper.Pagination, filter *CarFilter) (*pgHelper.Pagination, error)
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

func (r *CarRepository) GetCars(pg *pgHelper.Pagination, filter *CarFilter) (*pgHelper.Pagination, error) {
	location := models.Location{}
	locationId := filter.Location

	res := r.db.Model(&models.Location{}).Where("id = ? AND is_active = ?", locationId, true).First(&location)
	if res.Error != nil || res.Error == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Location is Not Active")
	}

	pickupWeekDay := int(filter.PickupDate.ToTime().Weekday())
	dropoffWeekDay := int(filter.DropoffDate.ToTime().Weekday())

	pickupTime := filter.PickupTime.ToTime()
	dropoffTime := filter.DropoffTime.ToTime()

	officeIDs := []uuid.UUID{}

	res = r.db.Model(&models.Office{}).Select("id").Where(
		"opening_hours <= ? AND opening_hours <=  ? AND closing_hours >= ? AND closing_hours >= ? AND location_id = ?",
		pickupTime, dropoffTime, pickupTime, dropoffTime, locationId,
	).Joins(
		"JOIN office_working_days ON office_working_days.office_id = offices.id",
	).Where(
		"office_working_days.working_day_id IN ?", []int{pickupWeekDay, dropoffWeekDay},
	).Find(&officeIDs)

	var cars []*models.Car
	var totalRows int64

	query := r.db.Model(&models.Car{}).Preload("Office").Preload("Vendor").
		Where(
			"status = ? AND office_id IN ?", models.CarStatusAvailable, officeIDs,
		).Find(&cars).Scopes().Count(&totalRows)

	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&cars)

	if query.Error != nil {
		return nil, query.Error
	}

	pg.Rows = &cars
	return pg, nil
}

func (r *CarRepository) CreateCar(car *models.Car) (*models.Car, error) {

	if err := r.db.Create(car).Error; err != nil {
		return nil, err
	}
	return car, nil
}
