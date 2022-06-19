package reservation

import (
	"github.com/eibrahimarisoy/car_rental/internal/car"
	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/office"
	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type ReservationService struct {
	reservationRepo ReservationRepositoryInterface
	carRepo         car.CarRepositoryInterface
	officeRepo      office.OfficeRepositoryInterface
	locationRepo    location.LocationRepositoryInterface
}

//go:generate mockgen -destination=../../mocks/reservation/reservation_service_interface.go -package=reservation github.com/eibrahimarisoy/car_rental/internal/reservation ReservationServiceInterface
type ReservationServiceInterface interface {
	GetReservations(pg *pgHelper.Pagination) (*[]models.Reservation, error)
	CreateReservation(reservation *models.Reservation) error
}

// NewReservationService creates a new reservation service
func NewReservationService(
	reservationRepo ReservationRepositoryInterface,
	carRepo car.CarRepositoryInterface,
	officeRepo office.OfficeRepositoryInterface,
	locationRepo location.LocationRepositoryInterface,
) *ReservationService {
	return &ReservationService{
		reservationRepo: reservationRepo,
		carRepo:         carRepo,
		officeRepo:      officeRepo,
		locationRepo:    locationRepo,
	}
}

// GetReservations returns all reservations
func (s *ReservationService) GetReservations(pg *pgHelper.Pagination) (*[]models.Reservation, error) {
	return s.reservationRepo.GetReservations(pg)
}

// CreateReservation creates a reservation and returns it
func (s *ReservationService) CreateReservation(reservation *models.Reservation) error {
	car, err := s.carRepo.GetCarByID(reservation.CarID)
	if err != nil {
		return err
	}

	pickupLocation, err := s.locationRepo.GetLocationByID(reservation.PickupLocationID)
	if err != nil {
		return err
	}
	pickupOffice, err := s.officeRepo.FindByOfficeAndVendorID(car.OfficeID, car.VendorID)
	if err != nil {
		return err
	}
	if !pickupOffice.IsAvaliable(reservation.PickupDate.ToTime(), reservation.PickupTime.ToTime()) {
		return errorHandler.PickupOfficeNotAvailableError
	}

	dropoffLocation, err := s.locationRepo.GetLocationByID(reservation.DropoffLocationID)
	if err != nil {
		return err
	}
	dropoffOffice, err := s.officeRepo.FindByOfficeAndVendorID(car.OfficeID, car.VendorID)
	if err != nil {
		return err
	}
	if !dropoffOffice.IsAvaliable(reservation.DropoffDate.ToTime(), reservation.DropoffTime.ToTime()) {
		return errorHandler.DropoffOfficeNotAvailableError
	}

	reservation, err = s.reservationRepo.CreateReservation(reservation)
	if err != nil {
		return err
	}
	status := models.CarStatusRented
	car.Status = &status
	car, err = s.carRepo.UpdateCarStatus(car)
	if err != nil {
		return err
	}
	reservation.Car = *car
	reservation.PickupLocation = *pickupLocation
	reservation.DropoffLocation = *dropoffLocation

	return nil
}
