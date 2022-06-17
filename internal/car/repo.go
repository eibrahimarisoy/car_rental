package car

import (
	"fmt"

	"github.com/eibrahimarisoy/car_rental/internal/models"

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

// TODO
func (r *CarRepository) GetCars(pg *pgHelper.Pagination, filter *CarFilter) (*pgHelper.Pagination, error) {
	fmt.Println("filter: ", filter)
	location := models.Location{}
	locationId := filter.Location

	res := r.db.Model(&models.Location{}).Where("id = ? AND is_active = ?", locationId, true).First(&location)
	if res.Error != nil || res.Error == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Location is Not Active")
	}

	pickupWeekDay := (filter.PickupDate.ToTime().Weekday())
	dropoffWeekDay := (filter.DropoffDate.ToTime().Weekday())

	pickupTime := filter.PickupTime.ToTime()
	dropoffTime := filter.DropoffTime.ToTime()

	fmt.Println("pickupWeekDay: ", pickupWeekDay)
	fmt.Println("dropoffWeekDay: ", dropoffWeekDay)

	fmt.Println("pickupTime: ", pickupTime)
	fmt.Println("dropoffTime: ", dropoffTime)

	office := models.Office{}
	res = r.db.Model(&models.Office{}).Where(
		"opening_hours <= ? AND opening_hours <=  ? AND closing_hours >= ? AND closing_hours >= ? AND location_id = ?",
		pickupTime, dropoffTime, pickupTime, dropoffTime, locationId,
	).Where(r.db.Table("office_working_days").Where("value IN ?", []int{1, 2, 3, 4, 5, 6, 7})).First(&office)

	fmt.Println("office: ", office)

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
