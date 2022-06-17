package car

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"
)

type CarFilter struct {
	PickupDate models.JsonDate `json:"pickup_date"`
	PickupTime models.JsonTime `json:"pickup_time"`

	DropoffDate models.JsonDate `json:"dropoff_date"`
	DropoffTime models.JsonTime `json:"dropoff_time"`

	Location uuid.UUID `json:"location_id"`
}

func (f *CarFilter) Validate() error {
	// if f.PickupDate.After(f.DropoffDate) {
	// 	return errors.New("pickup date must be before dropoff date")
	// }
	return nil
}
