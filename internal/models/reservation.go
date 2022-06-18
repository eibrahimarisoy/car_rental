package models

import (
	"github.com/google/uuid"
)

type Reservation struct {
	Base
	Status ReservationStatus `json:"status"`

	PickupLocationID uuid.UUID `json:"pickup_location_id"`
	PickupLocation   Location  `json:"pickup_location" gorm:"foreignkey:PickupLocationID"`

	PickupDate JsonDate `json:"pickup_date"`
	PickupTime JsonTime `json:"pickup_time"`

	DropoffLocationID uuid.UUID `json:"dropoff_location_id"`
	DropoffLocation   Location  `json:"dropoff_location" gorm:"foreignkey:DropoffLocationID"`

	DropoffDate JsonDate `json:"dropoff_date"`
	DropoffTime JsonTime `json:"dropoff_time"`

	VendorID uuid.UUID `json:"vendor_id"`
	Vendor   Vendor    `json:"vendor" gorm:"foreignkey:VendorID"`

	OfficeID uuid.UUID `json:"office_id"`
	Office   Office    `json:"office" gorm:"foreignkey:OfficeID"`

	DriverID uuid.UUID `json:"driver_id"`
	Driver   Driver    `json:"driver" gorm:"foreignkey:DriverID"`

	CarID uuid.UUID `json:"car_id"`
	Car   Car       `json:"car" gorm:"foreignkey:CarID"`
}
