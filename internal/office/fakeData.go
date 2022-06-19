package office

import (
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/vendors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	Monday      = models.WorkingDay{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Day: "Monday", Value: 1}
	Tuesday     = models.WorkingDay{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Day: "Tuesday", Value: 2}
	Wednesday   = models.WorkingDay{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Day: "Wednesday", Value: 3}
	Thursday    = models.WorkingDay{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Day: "Thursday", Value: 4}
	Friday      = models.WorkingDay{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Day: "Friday", Value: 5}
	Saturday    = models.WorkingDay{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Day: "Saturday", Value: 6}
	Sunday      = models.WorkingDay{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Day: "Sunday", Value: 7}
	WorkingDays = []models.WorkingDay{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}
)

var (
	openingHours_1 = models.JsonTime{}

	openingHours_2 = models.JsonTime{}
)

var (
	FakeOffice_1 = models.Office{
		Base:         models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		OpeningHours: openingHours_1.FromString("00:00"),
		ClosingHours: openingHours_1.FromString("23:59"),
		WorkingDays: []models.WorkingDay{
			Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday,
		},
		Vendor:   vendors.FakeVendor_1,
		Location: location.FakeLocation_1,
	}
	FakeOffice_2 = models.Office{
		Base:         models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		OpeningHours: openingHours_2.FromString("00:00"),
		ClosingHours: openingHours_2.FromString("23:59"),
		WorkingDays: []models.WorkingDay{
			Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday,
		},
		Vendor:   vendors.FakeVendor_2,
		Location: location.FakeLocation_2,
	}

	FakeOfficeList = []models.Office{FakeOffice_1, FakeOffice_2}
)
