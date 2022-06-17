package models

type Location struct {
	Base
	Name     *string `json:"name" gorm:"type:varchar(255);not null;unique"`
	IsActive bool    `json:"is_active" gorm:"type:boolean;"`
}
