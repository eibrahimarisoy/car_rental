package driver

import (
	"errors"
	"regexp"
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	"github.com/google/uuid"
)

type DriverRequest struct {
	FirstName            string          `json:"first_name"`
	LastName             string          `json:"last_name"`
	Email                string          `json:"email"`
	Phone                string          `json:"phone"`
	IdentificationNumber string          `json:"identification_number"`
	Birthday             models.JsonDate `json:"birthday"`
}

// Validate validates the driver request.
func (r *DriverRequest) Validate() error {
	if r.FirstName == "" {
		return errors.New("required data")
	}
	if r.LastName == "" {
		return errors.New("required data")
	}
	if r.Email == "" {
		return errors.New("required data")
	}
	if r.Phone == "" {
		return errors.New("required data")
	}
	if r.IdentificationNumber == "" {
		return errors.New("required data")
	}
	if match, _ := regexp.MatchString("^[1-9]{1}[0-9]{9}[02468]{1}$", r.IdentificationNumber); !match {
		return errorHandler.InvalidIdentificationNumber
	}
	if r.Birthday == models.JsonDate(time.Time{}) {
		return errors.New("required data")
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
