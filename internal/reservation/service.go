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

type ReservationServiceInterface interface {
	GetReservations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateReservation(reservation *models.Reservation) (*models.Reservation, error)
}

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

func (s *ReservationService) GetReservations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	return s.reservationRepo.GetReservations(pg)
}

func (s *ReservationService) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {
	car, err := s.carRepo.GetCarByID(reservation.CarID)
	if err != nil {
		return nil, err
	}

	pickupLocation, err := s.locationRepo.GetLocationByID(reservation.PickupLocationID)
	if err != nil {
		return nil, err
	}
	pickupOffice, err := s.officeRepo.FindByOfficeAndVendorID(car.OfficeID, car.VendorID)
	if err != nil {
		return nil, err
	}
	if !pickupOffice.IsAvaliable(reservation.PickupDate.ToTime(), reservation.PickupTime.ToTime()) {
		return nil, errorHandler.PickupOfficeNotAvailableError
	}

	dropoffLocation, err := s.locationRepo.GetLocationByID(reservation.DropoffLocationID)
	if err != nil {
		return nil, err
	}
	dropoffOffice, err := s.officeRepo.FindByOfficeAndVendorID(car.OfficeID, car.VendorID)
	if err != nil {
		return nil, err
	}
	if !dropoffOffice.IsAvaliable(reservation.DropoffDate.ToTime(), reservation.DropoffTime.ToTime()) {
		return nil, errorHandler.DropoffOfficeNotAvailableError
	}

	reservation, err = s.reservationRepo.CreateReservation(reservation)
	if err != nil {
		return nil, err
	}
	status := models.CarStatusRented
	car.Status = &status
	car, err = s.carRepo.UpdateCarStatus(car)
	if err != nil {
		return nil, err
	}
	reservation.Car = *car
	reservation.PickupLocation = *pickupLocation
	reservation.DropoffLocation = *dropoffLocation

	return reservation, nil
}
