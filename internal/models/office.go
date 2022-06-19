package models

import (
	"fmt"
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

// GetWorkingDays returns the working days of the office.
func (o *Office) GetWorkingDays() []WorkingDay {
	return o.WorkingDays
}

// IsAvaliable returns true if the office is available at the given time.
func (o *Office) IsAvaliable(date, time time.Time) bool {
	dateOK := false
	for _, wd := range o.WorkingDays {
		if wd.Value == uint(date.Weekday()) {
			dateOK = true
			break
		}
	}
	if dateOK && o.OpeningHours.ToTime().Before(time) && o.ClosingHours.ToTime().After(time) {
		fmt.Println("office is available")
		return true
	}
	return false
}
