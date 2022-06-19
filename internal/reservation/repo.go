package reservation

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"go.uber.org/zap"

	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockgen -destination=../../mocks/reservation/reservation_repository_interface.go -package=reservation github.com/eibrahimarisoy/car_rental/internal/reservation ReservationRepositoryInterface
type ReservationRepositoryInterface interface {
	GetReservations(pg *pgHelper.Pagination) (*[]models.Reservation, error)
	CreateReservation(reservation *models.Reservation) (*models.Reservation, error)
}

type ReservationRepository struct {
	db *gorm.DB
}

// NewReservationRepository creates a new reservation repository
func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

// Migration migrates database
func (r *ReservationRepository) Migration() {
	r.db.AutoMigrate(&models.Reservation{})
}

// GetReservations gets reservations from database with pagination
func (r *ReservationRepository) GetReservations(pg *pgHelper.Pagination) (*[]models.Reservation, error) {
	zap.L().Debug("reservation.repo.GetReservations", zap.Reflect("pg", *pg))

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
	zap.L().Debug("reservation.repo.CreateReservation", zap.Reflect("reservation", *reservation))

	if err := r.db.Create(reservation).Error; err != nil {
		return nil, err
	}
	return reservation, nil
}
