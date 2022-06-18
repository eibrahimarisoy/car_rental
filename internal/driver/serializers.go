package driver

import (
	"regexp"
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	"github.com/google/uuid"
)

var (
	IdentificationNumberRegex = regexp.MustCompile(`^[1-9]{1}[0-9]{9}[02468]{1}$`)
)

type DriverRequest struct {
	FirstName            string          `json:"first_name" validate:"required" binding:"required"`
	LastName             string          `json:"last_name" validate:"required" binding:"required"`
	Email                string          `json:"email" validate:"required" binding:"required,email"`
	Phone                string          `json:"phone" validate:"required" binding:"required,e164" format:"+905012345678"`
	Birthday             models.JsonDate `json:"birthday" validate:"required" binding:"required" time_format:"02-01-2006" format:"02-01-2006"`
	IdentificationNumber string          `json:"identification_number" validate:"required" binding:"required" format:"12345678901"`
}

// Validate validates the driver request.
func (r *DriverRequest) Validate() error {
	if match := IdentificationNumberRegex.MatchString(r.IdentificationNumber); !match {
		return errorHandler.InvalidIdentificationNumber
	}

	if !(time.Now().Sub(r.Birthday.ToTime()) > 18*365*24*time.Hour) {
		return errorHandler.DriverAgeNotValid
	}

	return nil
}

// ToDriver converts the driver request to driver model.
func (r *DriverRequest) ToDriver() *models.Driver {
	firstName := r.FirstName
	lastName := r.LastName
	email := r.Email
	phone := r.Phone
	identificationNumber := r.IdentificationNumber
	birthday := r.Birthday

	return &models.Driver{
		FirstName:            &firstName,
		LastName:             &lastName,
		Email:                &email,
		Phone:                &phone,
		IdentificationNumber: &identificationNumber,
		Birthday:             &birthday,
	}
}

type DriverResponse struct {
	ID                   uuid.UUID `json:"id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Email                string    `json:"email"`
	Phone                string    `json:"phone"`
	IdentificationNumber string    `json:"identification_number"`
	Birthday             string    `json:"birthday"`
}

// FromDriver converts the driver model to driver response.
func FromDriver(d *models.Driver) *DriverResponse {
	id := d.ID
	firstName := d.FirstName
	lastName := d.LastName
	email := d.Email
	phone := d.Phone
	identificationNumber := d.IdentificationNumber
	birthday := d.Birthday

	return &DriverResponse{
		ID:                   id,
		FirstName:            *firstName,
		LastName:             *lastName,
		Email:                *email,
		Phone:                *phone,
		IdentificationNumber: *identificationNumber,
		Birthday:             birthday.String(),
	}
}
