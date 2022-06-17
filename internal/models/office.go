package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

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
