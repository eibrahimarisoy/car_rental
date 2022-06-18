package models

import (
	"encoding/json"
	"time"
)

type ReservationStatus string

const (
	ReservationStatusPending   ReservationStatus = "pending"
	ReservationStatusConfirmed ReservationStatus = "confirmed"
	ReservationStatusCancelled ReservationStatus = "cancelled"
)

const dateFormat = "02-01-2006"

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

// ToTime converts JsonDate to time.Time
func (t JsonDate) ToTime() time.Time {
	return time.Time(t)
}

// FromTime converts time.Time to JsonDate
func (t JsonDate) FromTime(t2 time.Time) JsonDate {
	return JsonDate(t2)
}

// FromString converts string to JsonDate
func (t JsonDate) FromString(s string) JsonDate {
	t2, _ := time.Parse(dateFormat, s)
	return JsonDate(t2)
}

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

const timeFormat = "15:04"

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(timeFormat) + `"`), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t2, err := time.Parse(timeFormat, s)
	if err != nil {
		return err
	}
	*t = JsonTime(t2)
	return nil
}

func (t JsonTime) String() string {
	return time.Time(t).Format(timeFormat)
}

// FromString converts string to JsonTime
func (t JsonTime) FromString(s string) JsonTime {
	t2, _ := time.Parse(timeFormat, s)
	t = JsonTime(t2)
	return JsonTime(t2)
}

// ToTime converts JsonTime to time.Time
func (t JsonTime) ToTime() time.Time {
	return time.Time(t)
}
