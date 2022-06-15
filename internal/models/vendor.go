package models

type Vendor struct {
	Base
	Name *string `json:"name" gorm:"type:varchar(255);not null;unique"`
}
