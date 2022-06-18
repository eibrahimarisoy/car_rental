package reservation

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"

	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReservationRepositoryInterface interface {
	GetReservations(pg *pgHelper.Pagination) (*[]models.Reservation, error)
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
func (r *ReservationRepository) GetReservations(pg *pgHelper.Pagination) (*[]models.Reservation, error) {
	var reservations []models.Reservation
	var totalRows int64

	query := r.db.Model(&models.Reservation{}).Preload(clause.Associations).Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&reservations)

	if query.Error != nil {
		return nil, query.Error
	}

	return &reservations, nil
}

// CreateReservation creates a new reservation in database
func (r *ReservationRepository) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {

	if err := r.db.Create(reservation).Error; err != nil {
		return nil, err
	}
	return reservation, nil
}
