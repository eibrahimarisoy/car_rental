package car

import (
	"errors"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/office"
	"github.com/eibrahimarisoy/car_rental/internal/vendors"
	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	"github.com/google/uuid"
)

// CarRequest is the request for creating a car
type CarRequest struct {
	Status       models.CarStatusEnums    `json:"status" enums:"available,rented"`
	Name         *string                  `json:"name"`
	Fuel         models.FuelEnums         `json:"fuel" enums:"gas,diesel,electric"`
	Transmission models.TransmissionEnums `json:"transmission" enums:"automatic,manual"`
	VendorID     uuid.UUID                `json:"vendor_id"`
	OfficeID     uuid.UUID                `json:"office_id"`
}

// Validate validates the request
func (req *CarRequest) Validate() error {
	if req.Status == "" {
		return errors.New("status is required")
	}
	if req.Status.IsValid() == false {
		return errorHandler.InvalidEnumsValue
	}
	if req.Name == nil || *req.Name == "" {
		return errors.New("name is required")
	}
	if req.Fuel == "" {
		return errors.New("fuel is required")
	}
	if req.Fuel.IsValid() == false {
		return errorHandler.InvalidEnumsValue
	}
	if req.Transmission == "" {
		return errors.New("transmission is required")
	}
	if req.Transmission.IsValid() == false {
		return errorHandler.InvalidEnumsValue
	}
	if req.VendorID == uuid.Nil {
		return errors.New("vendor_id is required")
	}
	if req.OfficeID == uuid.Nil {
		return errors.New("office_id is required")
	}
	return nil
}

// ToCar converts the request to a car
func (req *CarRequest) ToCar() *models.Car {
	return &models.Car{
		Status:       &req.Status,
		Name:         req.Name,
		Fuel:         req.Fuel,
		Transmission: req.Transmission,
		VendorID:     req.VendorID,
		OfficeID:     req.OfficeID,
	}
}

type CarResponse struct {
	ID           uuid.UUID              `json:"id"`
	Status       models.CarStatusEnums  `json:"status"`
	Name         string                 `json:"name"`
	Fuel         string                 `json:"fuel"`
	Transmission string                 `json:"transmission"`
	Vendor       vendors.VendorResponse `json:"vendor"`
	Office       office.OfficeResponse  `json:"office"`
}

// FromCar converts a car to a response
func (res *CarResponse) FromCar(car *models.Car) {

	res.ID = car.ID
	res.Status = *car.Status
	res.Name = *car.Name
	res.Fuel = car.Fuel.String()
	res.Transmission = car.Transmission.String()
	res.Vendor = vendors.VendorResponse(res.Vendor)
	res.Office = office.OfficeResponse(res.Office)
}
