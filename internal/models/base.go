package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	ID        uuid.UUID  `gorm:"primary_key;type:uuid;" json:"id"`
}

type WorkingDay struct {
	gorm.Model
	Day   string `json:"day" gorm:"type:varchar(10);not null"`
	Value uint   `json:"value" gorm:"type:integer;not null;unique"`
}
