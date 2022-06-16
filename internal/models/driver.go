package models

type Driver struct {
	Base
	FirstName            *string   `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName             *string   `json:"last_name" gorm:"type:varchar(255);not null"`
	Email                *string   `json:"email" gorm:"type:varchar(255);not null;"`
	Phone                *string   `json:"phone" gorm:"type:varchar(255);not null"`
	IdentificationNumber *string   `json:"identification_number" gorm:"type:varchar(11);not null"`
	Birthday             *JsonDate `json:"birthday" gorm:"not null"`
}
