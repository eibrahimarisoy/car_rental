package errorHandler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	_type "github.com/eibrahimarisoy/car_rental/pkg/errorHandler/type"

	"gorm.io/gorm"
)

var (
	InternalServerError            = errors.New("Internal Server Error")
	NotFound                       = errors.New("Not Found")
	RequestTimeoutError            = errors.New("Request Timeout")
	CannotBindGivenData            = errors.New("Could not bind given data")
	ValidationError                = errors.New("Validation failed for given payload")
	UniqueError                    = errors.New("Item should be unique on database")
	Unauthorized                   = errors.New("Unauthorized")
	UnauthorizedError              = errors.New("Unauthorized")
	GivenAssociationNotFound       = errors.New("Given association not found")
	RequiredFieldError             = errors.New("Required field is missing")
	InvalidUUIDFormat              = errors.New("Invalid UUID format")
	InvalidIdentificationNumber    = errors.New("Invalid identification number")
	InvalidPhoneNumber             = errors.New("Invalid phone number")
	LocationNoAvailable            = errors.New("Location is not available")
	InvalidEnumsValue              = errors.New("Invalid enums value")
	InvalidDateTime                = errors.New("Invalid date time")
	InvalidDropOffDate             = errors.New("Drop off date must be after pick up date")
	InvalidEmail                   = errors.New("Invalid email")
	DriverAgeNotValid              = errors.New("Driver age is not valid")
	CarNotFoundError               = errors.New("Car not found")
	LocationIsNotAvailable         = errors.New("Location is not available")
	LocationNotFoundError          = errors.New("Location not found")
	OfficeNotFoundError            = errors.New("Office not found")
	DropoffOfficeNotAvailableError = errors.New("Drop off Office is not available")
	PickupOfficeNotAvailableError  = errors.New("Pick up Office is not available")
)

type RestError _type.APIErrorResponse

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.Code, e.Message, e.Details)
}

func (e RestError) Status() int {
	return int(e.Code)
}

func (e RestError) Causes() interface{} {
	return e.Details
}

func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		Code:    int64(status),
		Message: err,
		Details: causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		Code:    http.StatusInternalServerError,
		Message: InternalServerError.Error(),
		Details: causes,
	}
	return result
}

// ParseErrors Parser of error string messages returns RestError
func ParseErrors(err error) RestErr {
	fmt.Println(err)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, NotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
	case errors.Is(err, CannotBindGivenData) || strings.Contains(err.Error(), "EOF"):
		return NewRestError(http.StatusBadRequest, CannotBindGivenData.Error(), err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NewRestError(http.StatusNotFound, gorm.ErrRecordNotFound.Error(), err)
	case strings.Contains(err.Error(), "validation"):
		return NewRestError(http.StatusBadRequest, ValidationError.Error(), err)
	case strings.Contains(err.Error(), "required"):
		return NewRestError(http.StatusBadRequest, RequiredFieldError.Error(), err)
	case errors.Is(err, InvalidIdentificationNumber):
		return NewRestError(http.StatusBadRequest, InvalidIdentificationNumber.Error(), err)
	case strings.Contains(err.Error(), "cannot unmarshal"):
		return NewRestError(http.StatusBadRequest, CannotBindGivenData.Error(), err)
	case strings.Contains(err.Error(), "invalid UUID length"):
		return NewRestError(http.StatusBadRequest, InvalidUUIDFormat.Error(), err)
	case errors.Is(err, LocationIsNotAvailable):
		return NewRestError(http.StatusBadRequest, LocationIsNotAvailable.Error(), err)
	case errors.Is(err, InvalidEnumsValue):
		return NewRestError(http.StatusBadRequest, InvalidEnumsValue.Error(), err)
	case errors.Is(err, InvalidPhoneNumber):
		return NewRestError(http.StatusBadRequest, InvalidPhoneNumber.Error(), err)
	case strings.Contains(err.Error(), "parsing time"):
		return NewRestError(http.StatusBadRequest, InvalidDateTime.Error(), err)
	case strings.Contains(err.Error(), "drop_off_date must be after pick_up_date"):
		return NewRestError(http.StatusBadRequest, InvalidDropOffDate.Error(), err)
	case errors.Is(err, InvalidEmail):
		return NewRestError(http.StatusBadRequest, InvalidEmail.Error(), err)
	case errors.Is(err, DriverAgeNotValid):
		return NewRestError(http.StatusBadRequest, DriverAgeNotValid.Error(), err)
	case errors.Is(err, CarNotFoundError):
		return NewRestError(http.StatusNotFound, CarNotFoundError.Error(), err)
	case errors.Is(err, LocationNotFoundError):
		return NewRestError(http.StatusNotFound, LocationNotFoundError.Error(), err)
	case errors.Is(err, OfficeNotFoundError):
		return NewRestError(http.StatusNotFound, OfficeNotFoundError.Error(), err)
	case errors.Is(err, DropoffOfficeNotAvailableError):
		return NewRestError(http.StatusBadRequest, DropoffOfficeNotAvailableError.Error(), err)
	case errors.Is(err, PickupOfficeNotAvailableError):
		return NewRestError(http.StatusBadRequest, PickupOfficeNotAvailableError.Error(), err)
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
