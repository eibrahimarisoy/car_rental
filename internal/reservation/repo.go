package reservation

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"

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

// TODO
func (r *ReservationRepository) GetReservations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	var reservations []*models.Reservation
	var totalRows int64

	query := r.db.Model(&models.Reservation{}).Preload(clause.Associations).Scopes().Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&reservations)

	if query.Error != nil {
		return nil, query.Error
	}

	pg.Rows = &reservations
	return pg, nil
}

// TODO

func (r *ReservationRepository) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {

	if err := r.db.Create(reservation).Error; err != nil {
		return nil, err
	}
	return reservation, nil
}
