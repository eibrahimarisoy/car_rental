package car

import (
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"
)

var (
	locationName1, locationName2 = "locationName1", "locationName2"

	FakeLocation_1 = models.Location{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     &locationName1,
		IsActive: true,
	}

	FakeLocation_2 = models.Location{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     &locationName2,
		IsActive: false,
	}
)

// var FakeOfficeData = []models.Office{
// 	&Office{
// 		Base: &models.Base{
// 			ID:        uuid.New(),
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		OpeningHours: "00:00",
// 		ClosingHours: "23:59",}}}
