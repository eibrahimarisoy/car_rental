package reservation

import (
	"encoding/json"
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/car"
	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"
)

var firstName, lastName, email, phone, identificationNumber = "John", "Doe", "john@doe.com", "+905331234567", "37540829408"
var dateJson = &models.JsonDate{}
var timeJson = &models.JsonTime{}
var birthday = dateJson.FromString("01-01-2000")

var FakeDriver = models.Driver{
	Base:                 models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
	FirstName:            &firstName,
	LastName:             &lastName,
	Email:                &email,
	Phone:                &phone,
	IdentificationNumber: &identificationNumber,
	Birthday:             &birthday,
}

var (
	FakeReservation1 = &models.Reservation{
		Base:           models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		PickupLocation: location.FakeLocation_1,
		PickupDate:     dateJson.FromString("01-01-2020"),
		PickupTime:     timeJson.FromString("12:00"),

		DropoffLocation: location.FakeLocation_1,
		DropoffDate:     dateJson.FromString("01-01-2020"),
		DropoffTime:     timeJson.FromString("13:00"),

		Car:    car.FakeCar_1,
		Driver: FakeDriver,
	}

	FakeReservation2 = models.Reservation{
		Base:           models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		PickupLocation: location.FakeLocation_1,
		PickupDate:     dateJson.FromString("01-01-2020"),
		PickupTime:     timeJson.FromString("12:00"),

		DropoffLocation: location.FakeLocation_1,
		DropoffDate:     dateJson.FromString("01-01-2020"),
		DropoffTime:     timeJson.FromString("13:00"),

		Car:    car.FakeCar_1,
		Driver: FakeDriver,
	}

	FakeReservation3 = models.Reservation{
		Base:           models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		PickupLocation: location.FakeLocation_1,
		PickupDate:     dateJson.FromString("01-01-2020"),
		PickupTime:     timeJson.FromString("12:00"),

		DropoffLocation: location.FakeLocation_1,
		DropoffDate:     dateJson.FromString("01-01-2020"),
		DropoffTime:     timeJson.FromString("13:00"),

		Car:    car.FakeCar_1,
		Driver: FakeDriver,
	}

	FakeReservationList = []models.Reservation{*FakeReservation1, FakeReservation2, FakeReservation3}

	FakeReservation4 = models.Reservation{
		PickupLocationID: location.FakeLocation_1.ID,
		PickupDate:       dateJson.FromString("01-01-2020"),
		PickupTime:       timeJson.FromString("12:00"),

		DropoffLocationID: location.FakeLocation_1.ID,
		DropoffDate:       dateJson.FromString("01-01-2020"),
		DropoffTime:       timeJson.FromString("13:00"),

		CarID:  car.FakeCar_1.ID,
		Driver: FakeDriver,
	}

	FakeReservationRequest, _ = json.Marshal(FakeReservation4)
)
