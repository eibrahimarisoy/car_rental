package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ReservationStatus string

const (
	ReservationStatusPending   ReservationStatus = "pending"
	ReservationStatusConfirmed ReservationStatus = "confirmed"
	ReservationStatusCancelled ReservationStatus = "cancelled"
)
const dateFormat = "02-07-2006"

type JsonDate time.Time

func (t JsonDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(dateFormat) + `"`), nil
}

func (t *JsonDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t2, err := time.Parse(dateFormat, s)
	if err != nil {
		return err
	}
	*t = JsonDate(t2)
	return nil
}

func (t JsonDate) String() string {
	return time.Time(t).Format(dateFormat)
}

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
