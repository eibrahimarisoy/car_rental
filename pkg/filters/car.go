package filters

import (
	"errors"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"
)

type CarFilter struct {
	PickupDate models.JsonDate `json:"pickup_date" validate:"required" binding:"required"`
	PickupTime models.JsonTime `json:"pickup_time" validate:"required" binding:"required"`

	DropoffDate models.JsonDate `json:"dropoff_date" validate:"required" binding:"required"`
	DropoffTime models.JsonTime `json:"dropoff_time" validate:"required" binding:"required"`

	Location uuid.UUID `json:"location_id" validate:"required" binding:"uuid"`
}

func (f *CarFilter) Validate() error {
	if f.PickupDate.ToTime().After(f.DropoffDate.ToTime()) {
		return errors.New("pickup date must be before dropoff date")
	}
	return nil
}
