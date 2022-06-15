package models

import (
	"time"

	"github.com/google/uuid"
)

type ReservationStatus string

const (
	ReservationStatusPending   ReservationStatus = "pending"
	ReservationStatusConfirmed ReservationStatus = "confirmed"
	ReservationStatusCancelled ReservationStatus = "cancelled"
)

type Reservation struct {
	Base
	Status ReservationStatus `json:"status"`

	PickupLocationID uuid.UUID `json:"pickup_location_id"`
	PickupLocation   Location  `json:"pickup_location" gorm:"foreignkey:PickupLocationID"`

	PickupDate time.Time `json:"pickup_date"`
	PickupTime time.Time `json:"pickup_time"`

	DropoffLocationID uuid.UUID `json:"dropoff_location_id"`
	DropoffLocation   Location  `json:"dropoff_location" gorm:"foreignkey:DropoffLocationID"`

	DropoffDate time.Time `json:"dropoff_date"`
	DropoffTime time.Time `json:"dropoff_time"`

	VendorID uuid.UUID `json:"vendor_id"`
	Vendor   Vendor    `json:"vendor" gorm:"foreignkey:VendorID"`

	OfficeID uuid.UUID `json:"office_id"`
	Office   Office    `json:"office" gorm:"foreignkey:OfficeID"`

	DriverID uuid.UUID `json:"driver_id"`
	Driver   Driver    `json:"driver" gorm:"foreignkey:DriverID"`

	CarID uuid.UUID `json:"car_id"`
	Car   Car       `json:"car" gorm:"foreignkey:CarID"`
}
