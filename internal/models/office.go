package models

import (
	"time"

	"github.com/google/uuid"
)

type Office struct {
	Base
	OpeningHours time.Time `json:"opening_hours"`
	ClosingHours time.Time `json:"closing_hours"`

	VendorID uuid.UUID `json:"vendor_id" gorm:"not null"`
	Vendor   Vendor    `json:"vendor" gorm:"foreignkey:VendorID"`

	LocationID uuid.UUID `json:"location_id" gorm:"not null"`
	Location   Location  `json:"location" gorm:"foreignkey:LocationID"`

	WorkingDays []WorkingDay `json:"working_days" gorm:"many2many:office_working_days"`
}
