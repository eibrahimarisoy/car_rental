package reservation

import (
	"fmt"

	"github.com/eibrahimarisoy/car_rental/internal/models"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReservationRepositoryInterface interface {
	GetReservations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateReservation(reservation *models.Reservation) (*models.Reservation, error)
}

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

func (r *ReservationRepository) Migration() {
	r.db.AutoMigrate(&models.Reservation{})
}

// GetReservations gets reservations from database with pagination
func (r *ReservationRepository) GetReservations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	var reservations []*models.Reservation
	var totalRows int64

	query := r.db.Model(&models.Reservation{}).Preload(clause.Associations).Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&reservations)

	if query.Error != nil {
		return nil, query.Error
	}

	pg.Rows = &reservations
	return pg, nil
}

// CreateReservation creates a new reservation in database
func (r *ReservationRepository) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {
	// getting car from database to check if it is available
	tx := r.db.Begin()

	car := models.Car{}

	query := tx.Model(&models.Car{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("status = ? AND id = ?", models.CarStatusAvailable, reservation.CarID).First(&car)
	if query.Error == gorm.ErrRecordNotFound {
		return nil, errorHandler.CarNotFoundError
	} else if query.Error != nil {
		return nil, query.Error
	}

	// check drop location office is available in drop date and time
	dropoffLocation := models.Location{}
	query = tx.Model(&models.Location{}).Where("is_active = ? AND id = ?", true, reservation.DropoffLocationID).First(&dropoffLocation)

	if query.Error == gorm.ErrRecordNotFound {
		return nil, errorHandler.LocationNotFoundError
	} else if query.Error != nil {
		return nil, query.Error
	}

	dropoffOffice := models.Office{}
	query = tx.Model(&models.Office{}).Preload("WorkingDays").Where("id = ? AND vendor_id = ?", car.OfficeID, car.VendorID).First(&dropoffOffice)

	if query.Error == gorm.ErrRecordNotFound {
		return nil, errorHandler.OfficeNotFoundError
	} else if query.Error != nil {
		return nil, query.Error
	}

	if !dropoffOffice.IsAvaliable(reservation.PickupDate.ToTime(), reservation.PickupTime.ToTime(), reservation.DropoffDate.ToTime(), reservation.DropoffTime.ToTime()) {
		return nil, errorHandler.DropoffOfficeNotAvailableError
	}

	// check pick up location is available in pick up date and time
	pickupLocation := models.Location{}
	query = tx.Model(&models.Location{}).Where("is_active = ? AND id = ?", true, reservation.PickupLocationID).First(&pickupLocation)

	if query.Error == gorm.ErrRecordNotFound {
		return nil, errorHandler.LocationNotFoundError
	} else if query.Error != nil {
		return nil, query.Error
	}

	pickupOffice := models.Office{}
	query = tx.Model(&models.Office{}).Preload("WorkingDays").Where("id = ? AND vendor_id = ?", car.OfficeID, car.VendorID).First(&pickupOffice)

	if query.Error == gorm.ErrRecordNotFound {
		return nil, errorHandler.OfficeNotFoundError
	} else if query.Error != nil {
		return nil, query.Error
	}

	if !pickupOffice.IsAvaliable(reservation.PickupDate.ToTime(), reservation.PickupTime.ToTime(), reservation.DropoffDate.ToTime(), reservation.DropoffTime.ToTime()) {
		return nil, errorHandler.DropoffOfficeNotAvailableError
	}

	fmt.Println("CreateReservation", *reservation)
	if err := r.db.Create(reservation).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Model(&models.Car{}).Where("id = ?", reservation.CarID).Update("status", models.CarStatusRented)

	tx.Commit()
	return reservation, nil
}
