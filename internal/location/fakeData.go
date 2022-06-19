package location

import (
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/google/uuid"
)

var (
	LocationId1, LocationId2, LocationId3, LocationId4 = uuid.New(), uuid.New(), uuid.New(), uuid.New()

	locationName1, locationName2, locationName3, locationName4 = "locationName1", "locationName2", "locationName3", "locationName4"

	FakeLocation_1 = models.Location{
		Base:     models.Base{ID: LocationId1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     &locationName1,
		IsActive: true,
	}

	FakeLocation_2 = models.Location{
		Base:     models.Base{ID: LocationId2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     &locationName2,
		IsActive: false,
	}

	FakeLocation_3 = models.Location{
		Base:     models.Base{ID: LocationId3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     &locationName3,
		IsActive: true,
	}

	FakeLocation_4 = models.Location{
		Base:     models.Base{ID: LocationId4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     &locationName4,
		IsActive: false,
	}
)

var FakeLocationsData = []models.Location{FakeLocation_1, FakeLocation_2, FakeLocation_3}

var FakeDataWithPagination = pgHelper.Pagination{
	Limit: 2,
	Page:  1,
	Q:     "",
}
