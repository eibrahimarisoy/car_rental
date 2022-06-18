package models

import "github.com/google/uuid"

type FuelEnums string
type TransmissionEnums string

func (f *FuelEnums) String() string {
	return string(*f)
}

func (t *TransmissionEnums) String() string {
	return string(*t)
}

const (
	Gas      FuelEnums = "gas"
	Diesel   FuelEnums = "diesel"
	Electric FuelEnums = "electric"

	Automatic TransmissionEnums = "automatic"
	Manual    TransmissionEnums = "manual"
)

type Car struct {
	Base
	Name         *string           `json:"name" gorm:"type:varchar(255);not null"`
	Fuel         FuelEnums         `json:"fuel" gorm:"not null"`
	Transmission TransmissionEnums `json:"transmission" gorm:"not null"`

	VendorID uuid.UUID `json:"vendor_id" gorm:"not null"`
	Vendor   Vendor    `json:"vendor" gorm:"foreignkey:VendorID"`

	OfficeID uuid.UUID `json:"office_id" gorm:"not null"`
	Office   Office    `json:"office" gorm:"foreignkey:OfficeID"`
}
