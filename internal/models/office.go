package models

import (
	"time"

	"github.com/google/uuid"
)

type Office struct {
	Base
	OpeningHours JsonTime `json:"opening_hours"`
	ClosingHours JsonTime `json:"closing_hours"`

	VendorID uuid.UUID `json:"vendor_id" gorm:"not null"`
	Vendor   Vendor    `json:"vendor" gorm:"foreignkey:VendorID"`

	LocationID uuid.UUID `json:"location_id" gorm:"not null"`
	Location   Location  `json:"location" gorm:"foreignkey:LocationID"`

	WorkingDays []WorkingDay `json:"working_days" gorm:"many2many:office_working_days"`
}

func (o *Office) GetWorkingDays() []WorkingDay {
	return o.WorkingDays
}

// IsAvaliable returns true if the office is available at the given time.
func (o *Office) IsAvaliable(pickupDate, pickupTime, dropoffDate, dropoffTime time.Time) bool {
	for _, wd := range o.WorkingDays {
		if wd.Value == uint(pickupDate.Weekday()) && o.OpeningHours.ToTime().Before(pickupTime) && o.ClosingHours.ToTime().After(pickupTime) &&
			wd.Value == uint(dropoffDate.Weekday()) && o.OpeningHours.ToTime().Before(dropoffTime) && o.ClosingHours.ToTime().After(dropoffTime) {
			return true
		}
	}
	return false
}
