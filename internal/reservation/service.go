package reservation

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type ReservationService struct {
	reservationRepo ReservationRepositoryInterface
}

type ReservationServiceInterface interface {
	GetReservations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateReservation(reservation *models.Reservation) (*models.Reservation, error)
}

func NewReservationService(reservationRepo ReservationRepositoryInterface) *ReservationService {
	return &ReservationService{
		reservationRepo: reservationRepo,
	}
}

func (s *ReservationService) GetReservations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	return s.reservationRepo.GetReservations(pg)
}

func (s *ReservationService) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {
	return s.reservationRepo.CreateReservation(reservation)
}
