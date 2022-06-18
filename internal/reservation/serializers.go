package reservation

import (
	"errors"

	"github.com/eibrahimarisoy/car_rental/internal/car"
	"github.com/eibrahimarisoy/car_rental/internal/driver"
	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/office"
	"github.com/eibrahimarisoy/car_rental/internal/vendors"
	"github.com/google/uuid"
)

type ReservationRequest struct {
	PickupLocationID uuid.UUID       `json:"pickup_location_id" validate:"required" binding:"required" format:"UUID"`
	PickUpDate       models.JsonDate `json:"pick_up_date" validate:"required" binding:"required" format:"02-01-2006"`
	PickUpTime       models.JsonTime `json:"pick_up_time" validate:"required" binding:"required" format:"15:04"`

	DropoffLocationID uuid.UUID       `json:"dropoff_location_id" validate:"required" binding:"required" format:"UUID"`
	DropOffDate       models.JsonDate `json:"drop_off_date" validate:"required" binding:"required" format:"02-01-2006"`
	DropOffTime       models.JsonTime `json:"drop_off_time" validate:"required" binding:"required" format:"15:04"`

	VendorID uuid.UUID `json:"vendor_id" validate:"required" binding:"required" format:"UUID"`
	OfficeID uuid.UUID `json:"office_id" validate:"required" binding:"required" format:"UUID"`
	CarID    uuid.UUID `json:"car_id" validate:"required" binding:"required" format:"UUID"`

	Driver driver.DriverRequest `json:"driver_request"`
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
		VendorID:          r.VendorID,
		OfficeID:          r.OfficeID,
		CarID:             r.CarID,
		Driver:            *r.Driver.ToDriver(),
	}
}

type ReservationResponse struct {
	ID             uuid.UUID                 `json:"id"`
	PickupLocation location.LocationResponse `json:"pickup_location"`
	PickupDate     models.JsonDate           `json:"pick_up_date"`
	PickupTime     models.JsonTime           `json:"pick_up_time"`

	DropoffLocation location.LocationResponse `json:"dropoff_location"`
	DropoffDate     models.JsonDate           `json:"drop_off_date"`
	DropoffTime     models.JsonTime           `json:"drop_off_time"`

	Vendor vendors.VendorResponse `json:"vendor"`
	Office office.OfficeResponse  `json:"office"`
	Car    car.CarResponse        `json:"car"`
	Driver driver.DriverResponse  `json:"driver_response"`
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
	r.Vendor = vendors.VendorResponse(r.Vendor)
	r.Office = office.OfficeResponse(r.Office)
	r.Car = car.CarResponse(r.Car)
	r.Driver = driver.DriverResponse(r.Driver)
}
