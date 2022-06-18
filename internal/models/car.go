package models

import "github.com/google/uuid"

type Car struct {
	Base
	Status       *CarStatusEnums   `json:"status" gorm:"type:varchar(10);not null"`
	Name         *string           `json:"name" gorm:"type:varchar(255);not null"`
	Fuel         FuelEnums         `json:"fuel" gorm:"not null"`
	Transmission TransmissionEnums `json:"transmission" gorm:"not null"`

	VendorID uuid.UUID `json:"vendor_id" gorm:"not null"`
	Vendor   Vendor    `json:"vendor" gorm:"foreignkey:VendorID"`

	OfficeID uuid.UUID `json:"office_id" gorm:"not null"`
	Office   Office    `json:"office" gorm:"foreignkey:OfficeID"`
}
