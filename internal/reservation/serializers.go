package reservation

import (
	"errors"

	"github.com/eibrahimarisoy/car_rental/internal/car"
	"github.com/eibrahimarisoy/car_rental/internal/driver"
	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/google/uuid"
)

// ReservationRequest represents the reservation request.
type ReservationRequest struct {
	PickupLocationID uuid.UUID       `json:"pickup_location_id" validate:"required" binding:"required" format:"UUID"`
	PickUpDate       models.JsonDate `json:"pick_up_date" validate:"required" binding:"required" format:"02-01-2006"`
	PickUpTime       models.JsonTime `json:"pick_up_time" validate:"required" binding:"required" format:"15:04"`

	DropoffLocationID uuid.UUID       `json:"dropoff_location_id" validate:"required" binding:"required" format:"UUID"`
	DropOffDate       models.JsonDate `json:"drop_off_date" validate:"required" binding:"required" format:"02-01-2006"`
	DropOffTime       models.JsonTime `json:"drop_off_time" validate:"required" binding:"required" format:"15:04"`

	CarID uuid.UUID `json:"car_id" validate:"required" binding:"required" format:"UUID"`

	Driver *driver.DriverRequest `json:"driver_request" validate:"required" binding:"required"`
}

// Validate validates the reservation request.
func (r *ReservationRequest) Validate() error {

	if drop_off_date := r.DropOffDate.ToTime(); drop_off_date.Before(r.PickUpDate.ToTime()) {
		return errors.New("drop_off_date must be after pick_up_date")
	}

	if err := r.Driver.Validate(); err != nil {
		return err
	}
	return nil
}

// ToReservation converts the reservation request to reservation model.
func (r *ReservationRequest) ToReservation() *models.Reservation {
	return &models.Reservation{
		PickupLocationID:  r.PickupLocationID,
		PickupDate:        r.PickUpDate,
		PickupTime:        r.PickUpTime,
		DropoffLocationID: r.DropoffLocationID,
		DropoffDate:       r.DropOffDate,
		DropoffTime:       r.DropOffTime,

		CarID:  r.CarID,
		Driver: *r.Driver.ToDriver(),
	}
}

// ReservationResponse represents the reservation response.
type ReservationResponse struct {
	ID             uuid.UUID                 `json:"id"`
	PickupLocation location.LocationResponse `json:"pickup_location"`
	PickupDate     models.JsonDate           `json:"pick_up_date"`
	PickupTime     models.JsonTime           `json:"pick_up_time"`

	DropoffLocation location.LocationResponse `json:"dropoff_location"`
	DropoffDate     models.JsonDate           `json:"drop_off_date"`
	DropoffTime     models.JsonTime           `json:"drop_off_time"`

	Car    car.CarSimpleResponse `json:"car"`
	Driver driver.DriverResponse `json:"driver_response"`
}

// FromReservation converts the reservation model to reservation response.
func (r *ReservationResponse) FromReservation(reservation *models.Reservation) {
	r.ID = reservation.ID
	r.PickupLocation = location.LocationResponse(r.PickupLocation)
	r.PickupDate = reservation.PickupDate
	r.PickupTime = reservation.PickupTime
	r.DropoffLocation = location.LocationResponse(r.DropoffLocation)
	r.DropoffDate = reservation.DropoffDate
	r.DropoffTime = reservation.DropoffTime
	r.Car = *car.CarToCarSimpleResponse(&reservation.Car)
	r.Driver = driver.DriverResponse(r.Driver)
}

// ReservationToResponse converts a reservation to a response.
func ReservationToResponse(reservation *models.Reservation) *ReservationResponse {
	reservationResponse := &ReservationResponse{
		ID:              reservation.ID,
		PickupLocation:  *location.LocationToResponse(&reservation.PickupLocation),
		PickupDate:      reservation.PickupDate,
		PickupTime:      reservation.PickupTime,
		DropoffLocation: *location.LocationToResponse(&reservation.DropoffLocation),
		DropoffDate:     reservation.DropoffDate,
		DropoffTime:     reservation.DropoffTime,
		Car:             *car.CarToCarSimpleResponse(&reservation.Car),
		Driver:          *driver.DriverToResponse(&reservation.Driver),
	}
	return reservationResponse
}

type ReservationListResponse struct {
	pgHelper.Pagination
	Data []ReservationResponse `json:"cars"`
}

// ReservationsToReservationListResponse converts a list of reservations to a reservation list response
func ReservationsToReservationListResponse(reservations *[]models.Reservation, pagination *pgHelper.Pagination) *ReservationListResponse {
	response := &ReservationListResponse{}
	response.Pagination = *pagination
	for _, reservation := range *reservations {
		response.Data = append(response.Data, *ReservationToResponse(&reservation))
	}
	return response
}
