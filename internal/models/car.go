package models

import "github.com/google/uuid"

type CarStatusEnums string

const (
	CarStatusAvailable CarStatusEnums = "available"
	CarStatusRented    CarStatusEnums = "rented"
)

func (c *CarStatusEnums) IsValid() bool {
	switch *c {
	case CarStatusAvailable, CarStatusRented:
		return true
	}
	return false
}

type FuelEnums string

const (
	Gas      FuelEnums = "gas"
	Diesel   FuelEnums = "diesel"
	Electric FuelEnums = "electric"
)

func (f *FuelEnums) String() string {
	return string(*f)
}

func (f *FuelEnums) IsValid() bool {
	switch *f {
	case Gas, Diesel, Electric:
		return true
	}
	return false
}

type TransmissionEnums string

const (
	Automatic TransmissionEnums = "automatic"
	Manual    TransmissionEnums = "manual"
)

func (t *TransmissionEnums) String() string {
	return string(*t)
}

func (t *TransmissionEnums) IsValid() bool {
	switch *t {
	case Automatic, Manual:
		return true
	}
	return false
}

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
