package car

import (
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/office"
	"github.com/eibrahimarisoy/car_rental/internal/vendors"
	filters "github.com/eibrahimarisoy/car_rental/pkg/filters"
	"github.com/google/uuid"
)

var availableStatus = models.CarStatusAvailable
var carName1, carName2, carName3 = "car1", "car2", "car3"

var (
	FakeCar_1 = models.Car{
		Base:         models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Status:       &availableStatus,
		Name:         &carName1,
		Fuel:         models.Diesel,
		Transmission: models.Automatic,
		Vendor:       vendors.FakeVendor_1,
		Office:       office.FakeOffice_1,
	}
	FakeCar_2 = models.Car{
		Base:         models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Status:       &availableStatus,
		Name:         &carName2,
		Fuel:         models.Diesel,
		Transmission: models.Automatic,
		Vendor:       vendors.FakeVendor_2,
		Office:       office.FakeOffice_2,
	}
	FakeCar_3 = models.Car{
		Base:         models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Status:       &availableStatus,
		Name:         &carName3,
		Fuel:         models.Diesel,
		Transmission: models.Automatic,
		Vendor:       vendors.FakeVendor_1,
		Office:       office.FakeOffice_1,
	}

	FakeCarList = []models.Car{FakeCar_1, FakeCar_2, FakeCar_3}
)

var (
	dateJson = models.JsonDate{}
	timeJson = models.JsonTime{}

	FakeCarFilterData = filters.CarFilter{
		PickupDate:  dateJson.FromString("19-06-2022"),
		PickupTime:  timeJson.FromString("09:00"),
		DropoffDate: dateJson.FromString("21-06-2022"),
		DropoffTime: timeJson.FromString("12:00"),

		Location: location.LocationId1,
	}
)
