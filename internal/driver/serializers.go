package driver

import (
	"regexp"
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	"github.com/google/uuid"
)

var (
	PhoneRegex                = regexp.MustCompile(`^0[.\(\/]5[0-9][0-9][.\)\/][1-9]([0-9]){2}([0-9]){4}$`)
	IdentificationNumberRegex = regexp.MustCompile(`^[1-9]{1}[0-9]{9}[02468]{1}$`)
	EmailRegex                = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
)

type DriverRequest struct {
	FirstName            string          `json:"first_name" validate:"required" example:"John" binding:"required"`
	LastName             string          `json:"last_name" validate:"required" example:"Doe" binding:"required"`
	Email                string          `json:"email" validate:"required" binding:"required"`
	Phone                string          `json:"phone" validate:"required" binding:"required" format:"05012345678"`
	Birthday             models.JsonDate `json:"birthday" validate:"required" binding:"required,overyearsold" example:"02-08-2022" format:"02-01-2006"`
	IdentificationNumber string          `json:"identification_number" binding:"required" validate:"required" format:"12345678901"`
}

// Validate validates the driver request.
func (r *DriverRequest) Validate() error {
	if match := PhoneRegex.MatchString(r.Phone); !match {
		return errorHandler.InvalidPhoneNumber
	}
	if match := IdentificationNumberRegex.MatchString(r.IdentificationNumber); !match {
		return errorHandler.InvalidIdentificationNumber
	}
	if match := EmailRegex.MatchString(r.Email); !match {
		return errorHandler.InvalidEmail
	}

	birthday := r.Birthday.ToTime()

	if !(time.Now().Sub(birthday) > 18*365*24*time.Hour) {
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
